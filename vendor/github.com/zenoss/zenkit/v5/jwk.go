package zenkit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

type JSONWebKeySet struct {
	Keys []JSONWebKey `json:"keys"`
}

type JSONWebKey struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type JWKSClient struct {
	cachedJWKs map[string]JSONWebKey

	fetchTime time.Time     //last time we fetched
	cacheTTL  time.Duration //how often to refresh
	single    singleflight.Group
	url       string
}

func NewJWKSClient(url string, cacheTTL time.Duration) (*JWKSClient, error) {
	client := &JWKSClient{
		url:        url,
		cachedJWKs: make(map[string]JSONWebKey),
		cacheTTL:   cacheTTL,
	}
	//eagerly get the jwks. Should we do this?
	_, err := client.populateKeys()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *JWKSClient) GetJWK(key string) (*JSONWebKey, error) {
	now := time.Now()
	if now.After(c.fetchTime.Add(c.cacheTTL)) {
		if _, err, _ := c.single.Do("populate", c.populateKeys); err != nil {
			return nil, err
		}
	}
	k, ok := c.cachedJWKs[key]
	if !ok {
		return nil, fmt.Errorf("could find jwks for %v", key)
	}
	return &k, nil
}

//populates the keys. First return argument is always nil,
//it is only here to satisfy the contract for singleflight
func (c *JWKSClient) populateKeys() (interface{}, error) {
	if time.Now().After(c.fetchTime.Add(c.cacheTTL)) {
		jwks, err := downloadJWKS(c.url)
		if err != nil {
			return nil, err
		}
		c.cachedJWKs = make(map[string]JSONWebKey)
		for _, j := range jwks.Keys {
			c.cachedJWKs[j.Kid] = j
		}
		c.fetchTime = time.Now()
	}
	return nil, nil
}

func downloadJWKS(url string) (*JSONWebKeySet, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jwks = JSONWebKeySet{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return nil, err
	}
	return &jwks, nil
}
