/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("test called")
		test()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func test() {
	// Create a set using a map
	set := make(map[string]struct{})

	// Add items to the set
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}
	set["orange"] = struct{}{}
	set["grape"] = struct{}{}

	// Item to search for
	item := "banana"

	// Check if the item is in the set
	if _, ok := set[item]; ok {
		fmt.Printf("Item '%s' found in the set\n", item)
	} else {
		fmt.Printf("Item '%s' not found in the set\n", item)
	}
}
