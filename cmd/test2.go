/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// test2Cmd represents the test2 command
var test2Cmd = &cobra.Command{
	Use:   "test2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("test2 called")
		test2()
	},
}

func init() {
	rootCmd.AddCommand(test2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func test2() {
	// Sample array of fruits
	fruits := []string{"apple", "banana", "orange", "grape"}

	// Create a set using a map
	set := make(map[string]struct{})

	// Populate the set with fruits from the array
	for _, fruit := range fruits {
		set[fruit] = struct{}{}
	}

	// Item to search for
	item := "banana"

	// Check if the item is in the set
	if _, ok := set[item]; ok {
		fmt.Printf("Item '%s' found in the array\n", item)
	} else {
		fmt.Printf("Item '%s' not found in the array\n", item)
	}
}
