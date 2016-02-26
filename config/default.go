package config

import (
        "github.com/spf13/viper"
)

func main() {
        viper.setDefault("PORT", 8080)
        
        // AWS Section
        viper.setDefault("AWS_ACCESS_KEY_ID", "the truth lies beneath")
        viper.setDefault("AWS_SECRET_KEY", "shh.......")
        
        // DB Section
        viper.setDefault("MONGO_URL", "localhost/test")
        
        // Cache Section
        viper.setDefault("REDIS_URL", "redis://localhost:6379/test")
        
        
        viper.AutomaticEnv()
        
}
