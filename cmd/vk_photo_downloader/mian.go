package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/spf13/cobra"
)

func main() {
	rootCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var rootCmd = &cobra.Command{
		Use:   "vkphotodownloader",
		Short: "Tool for downloading photos from VK",
		Long: `The tool allows you to download photos from the desired VK album.
For using the tool you need to have an access token with the next scope: 
	- photos`,
	}

	cfg := NewConfig(rootCtx, rootCmd.Flags())

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		fmt.Println(cfg.Album)
		fmt.Println(cfg.AccessToken)
		fmt.Println(cfg.UserID)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	token := "<TOKEN>" // use os.Getenv("TOKEN")
	_ = api.NewVK(token)
}
