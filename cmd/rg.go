/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v62/github"
	"github.com/spf13/cobra"
)

// rgCmd represents the rg command which is core command
var rgCmd = &cobra.Command{
	Use:   "rg",
	Short: "Run the Repoguard",
	Long: `Command Available:
	Check collaborator list:
	1. rg --u <username of GitHub> --r <repo name>`,
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("GITHUB_TOKEN_RG")
		if len(token) == 0 {
			fmt.Println("Error: GITHUB_TOKEN_RG environment variable is not set.")
			fmt.Println("Please set the GITHUB_TOKEN_RG: export GITHUB_TOKEN_RG=<your token here>")
			os.Exit(1)
		}
		username, errUsername := cmd.Flags().GetString("u")
		repo, errRepo := cmd.Flags().GetString("r")

		if errUsername != nil {
			log.Println("Error getting username")
		}
		if errRepo != nil {
			log.Println("Error getting repository")
		}

		if len(repo) > 0 && len(username) > 0 {
			getListofCollaboratorUser(username, repo, token)
		} else {
			fmt.Println("Provide username and repository name: --u=<username> --r=<repository>")
		}
	},
}

// Display Collaborator users list in nice format
func getListofCollaboratorUser(username, repo, token string) {

	userList := getCollaboratorListFromGithub(username, repo, token)

	if len(userList) > 0 {
		fmt.Println("The following users have access to the", repo, "repository:")
		fmt.Println()
		for i := 0; i < len(userList); i++ {
			fmt.Println(*userList[i].Login)
		}
	} else {
		fmt.Print("No Collaborator on repository", repo)
	}
}

// Return the list of collobarators of a repo from GitHub
func getCollaboratorListFromGithub(username, repo, token string) []*github.User {
	client := github.NewClient(nil).WithAuthToken(token)
	ctx := context.Background()

	user, _, err := client.Repositories.ListCollaborators(ctx, username, repo, &github.ListCollaboratorsOptions{})
	if err != nil {
		log.Print("Not found")
	}
	return user
}

func init() {
	rootCmd.AddCommand(rgCmd)

	// flags
	rgCmd.Flags().String("u", "", "Username of GitHub")
	rgCmd.Flags().String("r", "", "Repsitory name")
}
