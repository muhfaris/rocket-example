package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var App Config

type AppConfig struct {
	App Config `mapstructure:"app"`
}

type Config struct {
	Name  string      `mapstructure:"name"`
	Port  int         `mapstructure:"port"`
	Fiber FiberConfig `mapstructure:"fiber"`
	Debug Debugging   `mapstructure:"debug"`

	Cache Cache `mapstructure:"cache"`

	Datastore Datastore `mapstructure:"datastore"`
}

type Debugging struct {
	Config bool `mapstructure:"config"`
}

type FiberConfig struct {
	EnablePrintRoutes        bool `mapstructure:"enable_print_routes"`
	EnableSplittingOnParsers bool `mapstructure:"enable_splitting_on_parsers"`
}

type Cache struct {
	Redis RedisConfig `mapstructure:"redis"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Datastore struct {
	MySQL MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("$HOME/.config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file, %s", err)
	}

	var _app AppConfig
	err := viper.Unmarshal(&_app)
	if err != nil {
		return fmt.Errorf("unable to decode config into struct, %v", err)
	}

	if _app.App.Debug.Config {
		viper.Debug()
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})

	// env
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	// assigned to app Global
	App = _app.App

	return nil
}
