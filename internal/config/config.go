package config

import (
	"github.com/spf13/pflag"
)

type AlbumType string

type AppConfig struct {
	UserID          int64  `config:"user,required,short=u,description=Numeric VK profile ID"`
	AccessToken     string `config:"token,required,short=t,description=User access_token with photos scope"`
	DestinationPath string `config:"dir,required,short=d,description=The path where you want to save the photos"`
	Album           string
}

func NewAppConfigWithDefaultValue() *AppConfig {
	return &AppConfig{
		Album: AlbumTypeWall,
	}
}

func (cfg *AppConfig) InstallFlags(flags *pflag.FlagSet) *AppConfig {
	flags.StringVarP(&cfg.Album, "album", "a", cfg.Album, `Album id or one of values (profile|wall|saved)`)
	flags.StringVarP(&cfg.AccessToken, "token", "t", "", `User access_token with photos scope`)
	flags.StringVarP(&cfg.DestinationPath, "dir", "d", "", `Directory where photos will be saved`)
	flags.Int64VarP(&cfg.UserID, "user", "u", 0, `Numeric VK profile ID`)

	return cfg
}

const (
	AlbumTypeProfile AlbumType = "profile"
	AlbumTypeWall              = "wall"
	AlbumTypeSaved             = "saved"
)
