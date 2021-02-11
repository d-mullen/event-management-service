package zenkit

import (
	"context"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

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
		if !StringInSlice(v, validAudiences) {
			return errInvalidAudience
		}
		return nil
	case []string:
		for _, a := range v {
			if StringInSlice(a, validAudiences) {
				return nil
			}
		}
		return errInvalidAudience
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
		return errInvalidAudience
	default:
		return errInvalidAudience
	}
}

func createValidationKeyGetter(ctx context.Context) (jwt.Keyfunc, error) {
	log := ContextLogger(ctx)
	if !viper.IsSet(JWTAudienceConfig) || !viper.IsSet(JWTIssuerConfig) || !viper.IsSet(JWKSURL) {
		return nil, errors.New("JWT validation config not set")
	}
	audience := viper.GetStringSlice(JWTAudienceConfig)
	issuer := viper.GetString(JWTIssuerConfig)
	jwksURL := viper.GetString(JWKSURL)
	cacheTTL := time.Duration(viper.GetInt64(JWKSCacheMinutes)) * time.Minute

	//TODO: set cache time as config
	jwksClient, err := NewJWKSClient(jwksURL, cacheTTL)
	if err != nil {
		log.WithError(err).WithFields(logrus.Fields{
			"jwksURL":  jwksURL,
			"cacheTTL": cacheTTL,
		}).
			Error("failed to create JWKS client")
		return nil, errors.Wrap(err, "Could not create JWKS client")
	}

	return func(token *jwt.Token) (interface{}, error) {
		// Verify 'aud' claim
		err := VerifyAudience(audience, token)
		if err != nil {
			return token, errors.New("Invalid audience.")
		}
		// Verify 'iss' claim
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
		}

		cert, err := getPemCert(jwksClient, token)
		if err != nil {
			log.WithError(err).Error("failed to get pem cert")
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	}, nil
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

func getMiddlewareLogger(ctx context.Context) *logrus.Entry {
	l := ContextLogger(ctx)
	if l.Logger.Out == ioutil.Discard {
		l = Logger(viper.GetString(ServiceLabel))
	}
	return l
}

func JWTHandler(ctx context.Context, handler http.Handler) (http.Handler, error) {
	keyGetter, err := createValidationKeyGetter(ctx)
	if err != nil {
		return nil, err
	}
	jwtMiddle := jwtmiddleware.New(jwtmiddleware.Options{
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			log := getMiddlewareLogger(r.Context())
			log.Errorf("JWTHandler error: %s", err)
			jwtmiddleware.OnError(w, r, err)
		},
		ValidationKeyGetter: keyGetter,
		SigningMethod:       jwt.SigningMethodRS256,
	})
	return jwtMiddle.Handler(handler), nil
}
