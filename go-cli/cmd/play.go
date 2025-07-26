/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		targetURL := "https://www.bilibili.com/video/BV13Z421M77N?vd_source=127961da3fc6c308c415223bd57e3f44"
		
		_, err := url.ParseRequestURI(targetURL)
		if err != nil {
			fmt.Printf("Invalid URL: %s\n", err)
			return
		}
		
		fmt.Println("春日影 MyGo!!!")
		if err := browser.OpenURL(targetURL); err != nil {
			fmt.Printf("Failed to open browser: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
