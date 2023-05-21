/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// sortsearch2Cmd represents the sortsearch2 command
var sortsearch2Cmd = &cobra.Command{
	Use:   "sortsearch2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("sortsearch2 called")
		// testSortSearch2()
		test2()
	},
}

func init() {
	rootCmd.AddCommand(sortsearch2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sortsearch2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sortsearch2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func testSortSearch2() {
	// Unsorted list of fruits
	fruits := []string{"grape", "orange", "apple", "banana"}

	// Sort the fruits slice with case-insensitive comparison
	sort.SliceStable(fruits, func(i, j int) bool {
		return strings.ToLower(fruits[i]) < strings.ToLower(fruits[j])
	})

	// Item to search for (case-insensitive)
	item := "BanAna"

	// Perform case-insensitive binary search
	index := sort.Search(len(fruits), func(i int) bool {
		return strings.EqualFold(fruits[i], item)
	})

	// Check if the item was found
	if index < len(fruits) && strings.EqualFold(fruits[index], item) {
		fmt.Printf("Item '%s' found at index %d\n", item, index)
	} else {
		fmt.Printf("Item '%s' not found\n", item)
	}
}

func test2() {
	// Unsorted list of fruits
	fruits := []string{"grape", "orange", "apple", "banana"}

	// Item to search for (case-insensitive)
	item := "BanAna"

	// Sort the fruits slice with case-insensitive comparison
	sort.SliceStable(fruits, func(i, j int) bool {
		return strings.ToLower(fruits[i]) < strings.ToLower(fruits[j])
	})

	// Perform case-insensitive binary search
	index := binarySearch(fruits, item, strings.EqualFold)

	// Check if the item was found
	if index >= 0 && strings.EqualFold(fruits[index], item) {
		fmt.Printf("Item '%s' found at index %d\n", item, index)
	} else {
		fmt.Printf("Item '%s' not found\n", item)
	}

	// Loop over the items in the list and show the index and item
	for i, fruit := range fruits {
		fmt.Printf("Index: %d, Item: %s\n", i, fruit)
	}
}

// binarySearch performs a case-insensitive binary search on a sorted slice using a custom comparison function
func binarySearch(slice []string, item string, cmp func(string, string) bool) int {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := (low + high) / 2
		if cmp(slice[mid], item) {
			return mid
		}
		if strings.ToLower(slice[mid]) < strings.ToLower(item) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
