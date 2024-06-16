/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/google/go-github/v62/github"
	"github.com/spf13/cobra"
)

// revokeCmd represents the remove command
var revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke access of user from a repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the credentials file path
		usr, err := user.Current()
		if err != nil {
			log.Fatal("Filesystemt user not found")
		}
		homeDir := usr.HomeDir
		repoGuardDirPath := ".repoguard"
		credentialsFileName := "credentials"
		credentialsFilePath := homeDir + "/" + repoGuardDirPath + "/" + credentialsFileName

		// first, check the credentials file exist
		_, errExistFile := os.Stat(credentialsFilePath)
		if os.IsNotExist(errExistFile) {
			log.Fatal("Configure the repoguard: $ repoguard configure")
		}

		data, errRead := os.ReadFile(credentialsFilePath)
		if errRead != nil {
			log.Fatal("Repoguarrd not configured ", errRead)
		}
		token := string(data)

		if len(token) == 0 {
			log.Fatal("Configure the repoguard with correct credentials")

		}

		// get the username and repo name from flags
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
	rootCmd.AddCommand(revokeCmd)
	revokeCmd.Flags().String("o", "", "Owner of repo")
	revokeCmd.Flags().String("u", "", "username of GitHub account")
	revokeCmd.Flags().String("r", "", "reposiotory name")
}
