/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// nickName string // ×小写字母开头hi被GORM忽略！
	NickName string
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		NickNameTextInput := pterm.DefaultInteractiveTextInput
		NickNameTextInput.DefaultText = "Please enter your nickname"
		NickName, _ := NickNameTextInput.Show()
		fmt.Println(NickName)

		db, _ := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
		db.AutoMigrate(&User{})

		db.Create(&User{
			NickName: NickName,
		})

		pterm.Info.Println("Save info to database")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
