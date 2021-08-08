package main

import (
	"context"

	"github.com/spf13/pflag"

	"github.com/IASamoylov/vk-photo-downloader-go/internal/config"
)

func NewConfig(ctx context.Context, flags *pflag.FlagSet) *config.AppConfig {
	cfg := config.NewAppConfigWithDefaultValue()

	cfg.InstallFlags(flags)

	return cfg
}

func InitDownloadCommand(rootCmd.AddCommand(wordCmd))  {

}
