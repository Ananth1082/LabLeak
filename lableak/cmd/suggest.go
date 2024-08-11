/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Ananth1082/LabLeak/repository"
	"github.com/spf13/cobra"
)

// suggestCmd represents the suggest command
var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: suggest,
}

func suggest(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Enter comment")
		cmd.Help()
		return
	}
	if name := cmd.Flag("username").Value.String(); name != "" {
		suggestion := args[0]
		err := repository.AddSuggestion(name, suggestion)
		if err != nil {
			fmt.Println("Username already used")
			return
		}
	} else {
		fmt.Println("Enter username")
		return
	}
}
func init() {
	rootCmd.AddCommand(suggestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// suggestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	suggestCmd.Flags().StringP("username", "u", "", "username for the suggestion")
}
