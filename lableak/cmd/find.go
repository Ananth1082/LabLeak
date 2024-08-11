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

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

hello find`,
	Run: findManual,
}

func findManual(cmd *cobra.Command, args []string) {
	resorces := strings.Split(args[0], "/")
	if len(resorces) != 3 {
		fmt.Println("Incorrect resource path <section>/<subject>/<manual name>")
		return
	}
	content, err := repository.GetManual(resorces[0], resorces[1], resorces[2])
	if err != nil {
		fmt.Println("Invalid resouce path- Resource path doesnt exist")
		return
	}
	fileFlag := cmd.Flag("file").Value.String()
	if fileFlag != "" {
		tmp, err := os.Create(fileFlag)
		if err != nil {
			fmt.Println("Error creating file", err)
			return
		}
		tmp.WriteString(content)
		return
	}
	fmt.Print(content)
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	findCmd.Flags().StringP("file", "f", "", "Give the file path to download the content")
}
