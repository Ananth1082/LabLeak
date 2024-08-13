/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Ananth1082/LabLeak/repository"
	"github.com/spf13/cobra"
)

// mendCmd represents the mend command
var mendCmd = &cobra.Command{
	Use:   "mend",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: mendManual,
}

func mendManual(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Enter arguments in the right format")
		cmd.Help()
		return
	}
	resource := strings.Split(args[0], "/")
	if len(resource) != 3 {
		fmt.Println("Invalid resource path")
		cmd.Help()
		return
	}
	password := cmd.Flag("password").Value.String()
	if pass, err := repository.CheckTokensForSection(resource[0], password); !pass || err != nil {
		fmt.Println("Unauthorized... This action will be logged to the admin")
		return
	}
	var content string
	if fileFlag := cmd.Flag("file").Value.String(); fileFlag != "" {
		fbytes, err := os.ReadFile(fileFlag)
		if err != nil {
			fmt.Println("Invalid file path")
			cmd.Help()
			return
		}
		content = string(fbytes)
	} else if textFlag := cmd.Flag("text").Value.String(); textFlag != "" {
		content = textFlag
	}

	err := repository.UpdateManual(resource[0], resource[1], resource[2], content)
	if err != nil {
		fmt.Println("Error mending manual", err.Error())
	}
}

func init() {
	rootCmd.AddCommand(mendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	mendCmd.Flags().StringP("password", "p", "", "password for admin access")
	mendCmd.Flags().StringP("file", "f", "", "file path to get new content")
	mendCmd.Flags().StringP("text", "t", "", "text content to replace the old one")
}
