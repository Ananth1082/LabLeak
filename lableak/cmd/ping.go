/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: pingServer,
}

func pingServer(cmd *cobra.Command, args []string) {
	d, err := strconv.Atoi(cmd.Flag("interval").Value.String())
	if err != nil {
		fmt.Println("Enter a number for interval")
		return
	}
	var s sync.WaitGroup
	s.Add(1)
	go func() {
		fmt.Println("Press <Enter> to stop the action")
		for {
			http.Get("http://lableak.onrender.com/ping")
			log.Println("Ping")
			time.Sleep(time.Duration(d) * time.Second)
		}
	}()
	go func() {
		var a string
		fmt.Scanln(&a)
		s.Done()
	}()
	s.Wait()
}
func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pingCmd.Flags().IntP("interval", "i", 100, "Time interval between ping calls")
}
