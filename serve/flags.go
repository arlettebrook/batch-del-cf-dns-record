package serve

import (
	"os"
	"sync"

	"github.com/spf13/pflag"

	"github.com/arlettebrook/batch-del-cf-dns-record/models"
)

var config *models.Config

var once sync.Once

func parseFlags(config *models.Config) {
	flagSet := pflag.NewFlagSet("Batch-del-cf-dns-record", pflag.ExitOnError)
	flagSet.StringVarP(&config.LogLevel,
		"log_level", "l", "info", "Log level")
	flagSet.StringVarP(&config.ApiToken,
		"api_token", "a", "", "Cloudflare API Token")
	flagSet.StringVarP(&config.ZoneID,
		"zone_id", "z", "", "Cloudflare Zone ID")
	flagSet.SortFlags = false
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
}

func GetConfig() *models.Config {
	once.Do(func() {
		config = &models.Config{}
		parseFlags(config)
	})
	return config
}
