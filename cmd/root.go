/*
Copyright © 2024 Adam Cholewiński ITSRICHARDSHANK@GMAIL.COM
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "playlist-dl",
	Short: "playlist-dl is a CLI tool for downloading videos from m3u8 playlists.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Version = "0.1.1"
}
