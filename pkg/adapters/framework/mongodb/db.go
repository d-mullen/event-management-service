package mongodb

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/pkg/errors"
	"github.com/zenoss/zenkit/v5"
)

type (
	Config struct {
		Address    string
		DBName     string
		Username   string
		Password   string
		CACert     string
		ClientCert string
		// options list https://www.mongodb.com/docs/manual/reference/connection-string/#connection-string-options
		Options    map[string]string
		DefaultTTL time.Duration
	}
)

func (cfg Config) URI() string {
	var (
		urlWithParams = url.URL{Scheme: "mongodb"}
		val           = url.Values{}
	)

	if len(cfg.Options) > 0 {
		for k, v := range cfg.Options {
			val.Set(k, v)
		}
	}

	switch {
	case len(cfg.Username) > 0 && len(cfg.Password) > 0:
		urlWithParams.User = url.UserPassword(cfg.Username, cfg.Password)
		val.Set("authMechanism", "SCRAM-SHA-256")
	case len(cfg.CACert) > 0 && len(cfg.ClientCert) > 0:
		val.Set("tlsCAFile", cfg.CACert)
		val.Set("tlsCertificateKeyFile", cfg.ClientCert)
		val.Set("authMechanism", "MONGODB-X509")
		val.Set("readPreference", "secondary")
	}
	urlWithParams.RawQuery = val.Encode()

	urlWithParams.Host = cfg.Address // host:port
	urlWithParams.Path = "/"

	return (&urlWithParams).String()
}

func NewMongoDatabaseConnection(ctx context.Context, cfg Config) (Database, error) {
	log := zenkit.ContextLogger(ctx)
	log.Info("Connecting to mongo db....")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI()))
	if err != nil {
		log.Errorf("failed to connect to MongoDB(%s): %q", cfg.URI(), err)
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get MongoDB client at %v", cfg.URI()))
	}
	log.Debugf("Connected to mongo db: %s", cfg.Address)
	// Ping the primary
	log.Debug("Pinging the primary...")
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Error("failed to connect to MongoDB")
		return nil, err
	}
	db := NewDatabaseHelper(client.Database(cfg.DBName))
	return db, nil
}
