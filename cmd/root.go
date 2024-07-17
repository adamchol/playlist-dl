/*
Copyright © 2024 Adam Cholewiński ITSRICHARDSHANK@GMAIL.COM
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "playlist-dl",
	Short: "playlist-dl is a CLI tool for downloading videos from m3u8 playlists.",
	Run:   runRoot,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("version", "v", false, "Version of the tool")
}

func runRoot(cmd *cobra.Command, args []string) {
	currentVersion := "0.1.1"

	version, _ := cmd.Flags().GetBool("version")
	if version {
		fmt.Println("playlist-dl CLI " + currentVersion)
	} else {
		cmd.Help()
	}
}
