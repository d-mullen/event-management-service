package zenkit

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Claims struct {
	jwt.RegisteredClaims
	Tenant             string           `json:"https://dev.zing.ninja/tenant,omitempty"`
	Email              string           `json:"https://dev.zing.ninja/email,omitempty"`
	Connection         string           `json:"https://dev.zing.ninja/connection,omitempty"`
	RestrictionFilters jwt.ClaimStrings `json:"https://dev.zing.ninja/restrictionfilters,omitempty"`
	Roles              []string         `json:"https://zenoss.com/roles,omitempty"`
	Groups             []string         `json:"https://zenoss.com/groups,omitempty"`
}

func JWTHandler(ctx context.Context, handler http.Handler) (http.Handler, error) {
	if !viper.IsSet(JWTAudienceConfig) || !viper.IsSet(JWTIssuerConfig) || !viper.IsSet(JWKSURL) {
		return nil, errors.New("JWT validation config not set")
	}

	var (
		audience = viper.GetStringSlice(JWTAudienceConfig)
		issuer   = viper.GetString(JWTIssuerConfig)
		jwksURL  = viper.GetString(JWKSURL)
		cacheTTL = time.Duration(viper.GetInt64(JWKSCacheMinutes)) * time.Minute
	)

	log := ContextLogger(ctx).WithFields(logrus.Fields{
		"jwksURL":  jwksURL,
		"cacheTTL": cacheTTL,
		"audience": audience,
		"issuer":   issuer,
	})

	jwksClient, err := NewJWKSClient(jwksURL, cacheTTL)
	if err != nil {
		log.WithError(err).Error("failed to create JWKS client")
		return nil, errors.Wrap(err, "Could not create JWKS client")
	}

	keyFn := createKeyFunc(ctx, jwksClient, audience, issuer)
	validator := func(ctx context.Context, tokenString string) (interface{}, error) {
		token, err := jwt.Parse(tokenString, keyFn, jwt.WithValidMethods([]string{"RS256", "HS256"}))
		if err != nil {
			return nil, fmt.Errorf("failed to parse a jwt token: %w", err)
		}

		if !token.Valid {
			log.Error("token is invalid")
			return nil, errors.New("token is invalid")
		}

		return token, nil
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log := ContextLogger(ctx)
		if log.Logger.Out == ioutil.Discard {
			log = Logger(viper.GetString(ServiceLabel))
		}

		log.Errorf("JWTHandler error: %s", err)

		jwtmiddleware.DefaultErrorHandler(w, r, err)
	}

	jwtMiddle := jwtmiddleware.New(validator, jwtmiddleware.WithErrorHandler(errorHandler))
	return jwtMiddle.CheckJWT(handler), nil
}

func createKeyFunc(ctx context.Context, client *JWKSClient, audience []string, issuer string) jwt.Keyfunc {
	log := ContextLogger(ctx).WithFields(logrus.Fields{
		"audience": audience,
		"issuer":   issuer,
	})

	return func(token *jwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		err := VerifyAudience(audience, token)
		if err != nil {
			log.WithError(err).Error("failed to validate audience")
			return token, errors.New("audience is invalid")
		}

		// Verify 'iss' claim
		validIssuer := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
		if !validIssuer {
			return token, errors.New("issuer is invalid")
		}

		cert, err := getPemCert(client, token)
		if err != nil {
			log.WithError(err).Error("failed to get pem cert")
			return token, fmt.Errorf("failed to get pem cert: %w", err)
		}

		result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		if err != nil {
			log.WithError(err).Error("failed to parse public RSA key from PEM")
			return nil, err
		}

		return result, nil
	}
}

func VerifyAudience(validAudiences []string, token *jwt.Token) error {
	var (
		errNilAudience     = errors.New("got nil audience from token")
		errInvalidAudience = errors.New("invalid audience")
	)

	aud := token.Claims.(jwt.MapClaims)["aud"]
	if aud == nil {
		return errNilAudience
	}

	switch v := aud.(type) {
	case string:
		if StringInSlice(v, validAudiences) {
			return nil
		}
	case []string:
		for _, a := range v {
			if StringInSlice(a, validAudiences) {
				return nil
			}
		}
	case []interface{}:
		for _, a := range v {
			aV, ok := a.(string)
			if !ok {
				return errors.New("invalid audience: slice element is not a string")
			}

			if StringInSlice(aV, validAudiences) {
				return nil
			}
		}
	}
	return errInvalidAudience
}

func getPemCert(client *JWKSClient, token *jwt.Token) (string, error) {
	cert := ""
	jwks, err := client.GetJWK(token.Header["kid"].(string))
	if err != nil {
		return cert, err
	}
	cert = "-----BEGIN CERTIFICATE-----\n" + jwks.X5c[0] + "\n-----END CERTIFICATE-----"
	return cert, nil
}

func GetTokenClaims(token string) (*Claims, error) {
	claims := &Claims{}
	parser := &jwt.Parser{
		ValidMethods:         nil,
		UseJSONNumber:        false,
		SkipClaimsValidation: true,
	}

	_, _, err := parser.ParseUnverified(token, claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
