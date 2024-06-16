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

// usersCmd represents the rg command which is core command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Get users list who have access",
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
		fmt.Print("Noone have access to ", repo)
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
	rootCmd.AddCommand(usersCmd)

	// flags
	usersCmd.Flags().String("u", "", "Username of GitHub")
	usersCmd.Flags().String("r", "", "Repsitory name")
}
