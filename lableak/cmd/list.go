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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `hello list`,
	Run:   listResources,
}

func listResources(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		list, err := repository.ListSections()
		if err != nil {
			fmt.Println("Invalid resource path: Path doesn't exist")
			return
		}
		fmt.Println("List of all Subjects is:")
		for i, doc := range list {
			fmt.Printf("%d. %s\n", i+1, doc.ID)
		}
		return
	}
	resource := strings.Split(args[0], "/")

	if len(resource) == 1 {
		list, err := repository.ListSubjects(resource[0])
		if err != nil {
			fmt.Println("Invalid resource path: Path doesn't exist")
			return
		}
		fmt.Println("List of all Subjects is:")
		for i, doc := range list {
			fmt.Printf("%d. %s\n", i+1, doc.ID)
		}
	}
	if len(resource) == 2 {
		list, err := repository.ListManuals(resource[0], resource[1])
		if err != nil {
			fmt.Println("Invalid resource path: Path doesn't exist")
			return
		}
		fmt.Println("List of all Manuals is:")
		for i, doc := range list {
			fmt.Printf("%d. %s\n", i+1, doc.ID)
		}
	}

}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
