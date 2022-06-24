package zenkit

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrorInvalidToken = errors.New("invalid JWT")
)

func ParseUnverified(token string, claims interface{}) error {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errors.Wrap(ErrorInvalidToken, "wrong number of segments")
	}

	claimBytes, err := decodeSegment(parts[1])
	if err != nil {
		return errors.Wrap(err, "unable to decode claims segment of JWT")
	}

	if err := json.Unmarshal(claimBytes, claims); err != nil {
		return errors.Wrap(err, "unable to deserialize claims")
	}
	return nil
}

func decodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}
