/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"repoguard/rg"

	"github.com/google/go-github/v62/github"
	"github.com/spf13/cobra"
)

// revokeCmd represents the remove command
var revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke access of user from a repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		username, err := rg.LoadUserName()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		token, err := rg.LoadToken()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// get the revokeUser and repo name from flags
		revokeUser, errrevokeUser := cmd.Flags().GetString("u")
		repo, errRepo := cmd.Flags().GetString("r")

		if errrevokeUser != nil {
			log.Println("Error getting revokeUser")
		}

		if errRepo != nil {
			log.Println("Error getting repository")
		}

		if len(repo) > 0 && len(revokeUser) > 0 {
			err := revokeUserAccess(username, revokeUser, repo, token)
			if err != nil {
				log.Println("Error revoking user access")
				os.Exit(1)
			} else {
				log.Println("Successfully removed user", revokeUser)
			}

		} else {
			fmt.Println("Provide repository and username of user to revoke access: --r=<repository> --u=<revokeUser>")
		}
	},
}

func revokeUserAccess(owner, revokeUser, repo, token string) error {
	client := github.NewClient(nil).WithAuthToken(token)
	ctx := context.Background()

	_, err := client.Repositories.RemoveCollaborator(ctx, owner, repo, revokeUser)

	return err
}

func init() {
	rootCmd.AddCommand(revokeCmd)
	revokeCmd.Flags().String("u", "", "revokeUser of GitHub account")
	revokeCmd.Flags().String("r", "", "reposiotory name")
}
