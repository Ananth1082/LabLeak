/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/Ananth1082/LabLeak/repository"
	"github.com/spf13/cobra"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: killResources,
}

func killResources(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Enter arguments")
		return
	}
	resource := strings.Split(args[0], "/")
	pass := cmd.Flag("password").Value.String()
	if p, err := repository.CheckTokensForSection(resource[0], pass); err != nil && !p {
		fmt.Println("Unauthorized... This action will be logged to the admin")
		return
	}
	if len(resource) == 1 {
		err := repository.DeleteSection(resource[0])
		if err != nil {
			fmt.Println("Invalid resource path: Path doesn't exist")
			return
		}
		fmt.Println("Section deleted successfuly")
	} else if len(resource) == 2 {
		err := repository.DeleteSubject(resource[0], resource[1])
		if err != nil {
			fmt.Println("Invalid resource path: Path doesn't exist")
			return
		}
		fmt.Println("Subject deleted successfuly")
	} else if len(resource) == 3 {
		err := repository.DeleteManual(resource[0], resource[1], resource[2])
		if err != nil {
			fmt.Println("Invalid resource path: Path doesn't exist")
			return
		}
		fmt.Println("Manual deleted successfuly")
	} else {
		fmt.Println("Invalid resource path ")
	}
}
func init() {
	rootCmd.AddCommand(killCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// killCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	killCmd.Flags().StringP("password", "p", "", "password to get admin access")
}
