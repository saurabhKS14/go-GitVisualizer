/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	Constants "gitviz/Utils"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitviz",
	Short: "A brief description of your application",
	Long: `git-viz is a command-line interface (CLI) tool built in Go to visualize Git repository data.

It provides several interactive features, such as commit heatmaps, contributor statistics, and more,
to help developers understand their Git history and collaboration trends.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.PlainOpen(".")
		if err != nil {
			log.Fatalf("This is not a git repository : %v", err)
		}
		if args[0] == "." {
			branches, err := repo.Branches()
			if err != nil {
				log.Fatalf("Failed to get the HEAD : %v", err)
			}
			err = branches.ForEach(func(ref *plumbing.Reference) error {
				currentBranch, _ := repo.Head()
				if currentBranch.Name().Short() == ref.Name().Short() {
					fmt.Println("* " + Constants.Green + ref.Name().Short() + Constants.Reset)
				} else {
					fmt.Println("  " + ref.Name().Short())
				}
				return nil
			})

			if err != nil {
				log.Fatalf("Error Iterating over the branches : %v", err)
			}
		} else {
			fmt.Println("This is a normal command")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitviz.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
