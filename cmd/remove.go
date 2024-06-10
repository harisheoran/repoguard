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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a user from repository access",
	Long:  `Revoke access of user from a repository`,
	Run: func(cmd *cobra.Command, args []string) {
		token := os.Getenv("GITHUB_TOKEN_RG")
		if len(token) == 0 {
			fmt.Println("Error: GITHUB_TOKEN_RG environment variable is not set.")
			fmt.Println("Please set the GITHUB_TOKEN_RG: export GITHUB_TOKEN_RG=<your token here>")
			os.Exit(1)
		}

		username, errUsername := cmd.Flags().GetString("u")
		repo, errRepo := cmd.Flags().GetString("r")
		owner, errOwner := cmd.Flags().GetString("o")

		if errUsername != nil {
			log.Println("Error getting username")
		}

		if errRepo != nil {
			log.Println("Error getting repository")
		}

		if errOwner != nil {
			log.Println("Error getting ownername")
		}

		if len(repo) > 0 && len(username) > 0 && len(owner) > 0 {
			err := revokeUserAccess(owner, username, repo, token)
			if err != nil {
				log.Println("Error revoking user access")
				os.Exit(1)
			} else {
				log.Println("Successfully removed user", username)
			}

		} else {
			fmt.Println("Provide Owner, repository and username to reovke access: --o=<owner> --r=<repository> --u=<username>")
		}
	},
}

func revokeUserAccess(owner, username, repo, token string) error {
	client := github.NewClient(nil).WithAuthToken(token)
	ctx := context.Background()

	_, err := client.Repositories.RemoveCollaborator(ctx, owner, repo, username)

	return err
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().String("o", "", "Owner of repo")
	removeCmd.Flags().String("u", "", "username of GitHub account")
	removeCmd.Flags().String("r", "", "reposiotory name")
}
