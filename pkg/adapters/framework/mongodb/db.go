package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"

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
		DefaultTTL time.Duration
	}
)

func (cfg Config) URI() string {
	switch {
	case len(cfg.Username) > 0 && len(cfg.Password) > 0:
		return fmt.Sprintf("mongodb://%s:%s@%s", cfg.Username, cfg.Password, cfg.Address)
	case len(cfg.CACert) > 0 && len(cfg.ClientCert) > 0:
		return fmt.Sprintf(
			"mongodb://%s/?authMechanism=MONGODB-X509&tlsCAFile=%s&tlsCertificateKeyFile=%s",
			cfg.Address, cfg.CACert, cfg.ClientCert)
	default:
		return fmt.Sprintf("mongodb://%s", cfg.Address)
	}
}

func NewMongoDatabaseConnection(ctx context.Context, cfg Config) (Database, error) {
	log := zenkit.ContextLogger(ctx)
	log.Info("Connecting to mongo db....")
	clientOpts := options.Client().ApplyURI(cfg.URI())
	if len(cfg.CACert) > 0 && len(cfg.ClientCert) > 0 {
		credential := options.Credential{
			AuthMechanism: "MONGODB-X509",
		}
		clientOpts = clientOpts.SetAuth(credential)
	}
	client, err := mongo.Connect(ctx, clientOpts)
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
