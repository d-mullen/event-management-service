package redis

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/spaolacci/murmur3"
	"github.com/zenoss/event-management-service/pkg/models/event"
)

const (
	cursorKeyPrefix = "eqs-cursors"
)

type (
	RedisCommander interface {
		GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd
		Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd
	}
	cursorAdapter struct {
		client RedisCommander
	}
)

func CursorToString(c *event.Cursor) string {
	b, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	asMap := make(map[string]any)
	err = json.Unmarshal(b, &asMap)
	if err != nil {
		return ""
	}
	keys := make([]string, 0, len(asMap))
	for k := range asMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	buf := bytes.Buffer{}
	for _, k := range keys {
		buf.WriteString(fmt.Sprintf("%s=%s,", k, asMap[k]))
	}
	hash := murmur3.New128()
	hash.Write(buf.Bytes())
	v1, v2 := hash.Sum128()
	data := make([]byte, 8*2)
	binary.LittleEndian.PutUint64(data[:8], v1)
	binary.LittleEndian.PutUint64(data[8:], v2)
	return base64.URLEncoding.EncodeToString(data)
}

var _ event.CursorRepository = &cursorAdapter{}

func NewAdapter(cl RedisCommander) *cursorAdapter {
	return &cursorAdapter{
		client: cl,
	}
}

func (a *cursorAdapter) New(ctx context.Context, req *event.Cursor) (string, error) {
	if req == nil {
		return "", errors.New("invalid cursor: nil value")
	}
	if len(req.ID) == 0 {
		req.ID = uuid.New().String()
	}
	cur := CursorToString(req)
	if len(cur) == 0 {
		return "", errors.New("invalid cursor")
	}
	cursorBytes, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal cursor: %q", err)
	}
	cmd := a.client.Set(ctx, fmt.Sprintf("%s-%s", cursorKeyPrefix, cur), cursorBytes, time.Hour)
	if err := cmd.Err(); err != nil {
		return "", err
	}
	return cur, nil
}

func (a *cursorAdapter) Get(ctx context.Context, key string) (*event.Cursor, error) {
	cmd := a.client.GetEx(ctx, fmt.Sprintf("%s-%s", cursorKeyPrefix, key), 30*time.Minute)
	b, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	cursor := &event.Cursor{}
	err = json.Unmarshal(b, cursor)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cursor: %q", err)
	}
	return cursor, nil
}
