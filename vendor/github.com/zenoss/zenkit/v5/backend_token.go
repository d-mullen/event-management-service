package zenkit

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// TokenSourceFactory helps define the oauth2 source Factory
type TokenSourceFactory func(ctx context.Context, clientID, clientSecret, audience, domain, tenant string) oauth2.TokenSource

var DefaultTokenSourceFactory TokenSourceFactory = sourceFactory

func sourceFactory(ctx context.Context, clientID, clientSecret, audience, domain, tenant string) oauth2.TokenSource {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     fmt.Sprintf("https://%s/oauth/token", domain),
		EndpointParams: url.Values{
			"audience":        []string{audience},
			"zing_tenant":     []string{tenant},
			"zing_connection": []string{"Username-Password-Authentication"},
		},
	}
	return config.TokenSource(ctx)
}

// BackendTokenSource is an oauth2 Token which allows backend services to comminicate with one another
type BackendTokenSource interface {
	Token(ctx context.Context, tenant string) (*oauth2.Token, error)
}

type backendTokenSource struct {
	clientID      string
	clientSecret  string
	audience      string
	domain        string
	tokens        map[string]oauth2.TokenSource
	mut           sync.RWMutex
	sourceFactory TokenSourceFactory
}

func (b *backendTokenSource) String() string {
	return fmt.Sprintf("client_id: %s client_secret: ***** audience: %s domain: %s",
		b.clientID, b.audience, b.domain)
}

func (b *backendTokenSource) Token(ctx context.Context, tenant string) (*oauth2.Token, error) {
	if tenant == "" {
		return nil, errors.New("you must provide non-empty tenant")
	}
	b.mut.Lock()
	ts, ok := b.tokens[tenant]
	if !ok {
		if b.sourceFactory != nil {
			ts = b.sourceFactory(context.Background(), b.clientID, b.clientSecret, b.audience, b.domain, tenant)
		} else {
			ts = sourceFactory(context.Background(), b.clientID, b.clientSecret, b.audience, b.domain, tenant)
		}
		b.tokens[tenant] = ts
	}
	b.mut.Unlock()
	return ts.Token()
}

// NewBackendTokenSourceFromConfig creates a backend token via viper configs
//   to allow services to communicate with one another.
func NewBackendTokenSourceFromConfig(v *viper.Viper) (BackendTokenSource, error) {
	clientId := v.GetString(BackendClientIDConfig)
	secret := v.GetString(BackendClientSecretConfig)
	audience := v.GetString(BackendClientAudience)
	domain := v.GetString(Auth0DomainConfig)
	message := "missing config %s"

	if len(clientId) == 0 {
		return nil, errors.New(fmt.Sprintf(message, BackendClientIDConfig))
	}

	if len(secret) == 0 {
		return nil, errors.New(fmt.Sprintf(message, BackendClientSecretConfig))
	}

	if len(audience) == 0 {
		return nil, errors.New(fmt.Sprintf(message, BackendClientAudience))
	}

	if len(domain) == 0 {
		return nil, errors.New(fmt.Sprintf(message, Auth0DomainConfig))
	}

	return NewBackendTokenSource(
		clientId,
		secret,
		audience,
		domain,
		DefaultTokenSourceFactory), nil
}

// NewBackendTokenSource creates a backend token to allow services to communicate with one another.
func NewBackendTokenSource(clientID, clientSecret, audience, domain string, factory TokenSourceFactory) BackendTokenSource {

	ret := &backendTokenSource{
		clientID:      clientID,
		clientSecret:  clientSecret,
		audience:      audience,
		domain:        domain,
		sourceFactory: factory,
	}
	ret.tokens = make(map[string]oauth2.TokenSource)
	return ret
}
