/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sortsearch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("sort called")
		sortTest()
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func sortTest() {
	// Unsorted list of fruits
	fruits := []string{"grape", "orange", "apple", "banana"}

	// Sort the fruits slice
	sort.Strings(fruits)

	// Item to search for
	item := "bananaA"

	// Perform binary search
	index := sort.SearchStrings(fruits, item)

	// Check if the item was found
	if index < len(fruits) && fruits[index] == item {
		fmt.Printf("Item '%s' found at index %d\n", item, index)
	} else {
		fmt.Printf("Item '%s' not found\n", item)
	}
}
