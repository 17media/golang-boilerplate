package settings

import (
	"github.com/spf13/viper"
)

// DevSettings overrides settings to development mode
func DevSettings() {
	// DEBUG mode
	viper.SetDefault("DEBUG", true)
}
