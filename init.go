package main

import (
	"github.com/spf13/viper"
	"github.com/zenoss/event-management-service/config"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
)

func MongoConfigFromEnv(v *viper.Viper) mongodb.Config {
	viperCfg := viper.GetViper()
	if v != nil {
		viperCfg = v
	}
	return mongodb.Config{
		Address:    viperCfg.GetString(config.MongoDBAddr),
		DBName:     viperCfg.GetString(config.MongoDBName),
		CACert:     viperCfg.GetString(config.MongoDBCACertificate),
		ClientCert: viperCfg.GetString(config.MongoDBClientCertificate),
		Options:    viper.GetStringMapString(config.MongoClientOptions),
	}
}
