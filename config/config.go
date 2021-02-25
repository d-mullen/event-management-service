package config

import (
	"github.com/spf13/viper"
)

const (
	// EventManagementEnabledConfig determines if the API endpoint for event management is enabled for this server
	EventManagementEnabledConfig = "eventmanagement.enabled"
)

// InitDefaults sets defaults values for this server's configuration
func init() {
	viper.SetDefault(EventManagementEnabledConfig, false)
}
