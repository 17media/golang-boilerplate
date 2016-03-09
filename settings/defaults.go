package settings

import (
	"github.com/spf13/viper"
)

// InitSettings initialize the settings for the app
func InitSettings() {
	// DEBUG mode
	viper.SetDefault("DEBUG", false)

	// Port
	viper.SetDefault("PORT", "8000")

	// AWS Section
	viper.SetDefault("AWS_ACCESS_KEY_ID", "the truth lies beneath")
	viper.SetDefault("AWS_SECRET_KEY", "shh.......")

	// DB Section
	viper.SetDefault("MONGODB_URL", "localhost/test")

	// Cache Section
	viper.SetDefault("REDIS_URL", "localhost:6379/test")

	viper.AutomaticEnv()

	if viper.GetBool("DEVELOPMENT") {
		DevSettings()
	}
}
