/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// test3Cmd represents the test3 command
var test3Cmd = &cobra.Command{
	Use:   "binarysearch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("test3 called")
		binarysearchTest()
	},
}

func init() {
	rootCmd.AddCommand(test3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// BinaryTreeNode represents a node in a binary tree
type BinaryTreeNode struct {
	Value       string
	Left, Right *BinaryTreeNode
}

// BuildBinaryTree builds a binary tree from a list of fruits
func BuildBinaryTree(fruits []string) *BinaryTreeNode {
	var root *BinaryTreeNode

	for _, fruit := range fruits {
		root = insertNode(root, fruit)
	}

	return root
}

// insertNode inserts a new node into the binary tree
func insertNode(root *BinaryTreeNode, value string) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{
			Value: value,
		}
	}

	if strings.ToLower(value) < strings.ToLower(root.Value) {
		root.Left = insertNode(root.Left, value)
	} else {
		root.Right = insertNode(root.Right, value)
	}

	return root
}

// SearchBinaryTree performs a case-insensitive search operation in a binary tree
func SearchBinaryTree(node *BinaryTreeNode, item string) bool {
	if node == nil {
		return false
	}

	if strings.EqualFold(node.Value, item) {
		return true
	}

	if strings.ToLower(item) < strings.ToLower(node.Value) {
		return SearchBinaryTree(node.Left, item)
	}

	return SearchBinaryTree(node.Right, item)
}

func binarysearchTest() {
	// List of fruits
	fruits := []string{"apple", "banana", "orange", "grape"}

	// Build the binary tree from the list of fruits
	root := BuildBinaryTree(fruits)

	// Item to search for (case-insensitive)
	item := "bananA"

	// Perform case-insensitive search operation in the binary tree
	found := SearchBinaryTree(root, item)

	if found {
		fmt.Printf("Item '%s' found in the binary tree\n", item)
	} else {
		fmt.Printf("Item '%s' not found in the binary tree\n", item)
	}
}
