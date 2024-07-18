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

// usersCmd represents the rg command which is core command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Get users list who have access",
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

		// get the repo name from flags
		repo, err := cmd.Flags().GetString("r")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(repo) > 0 {
			getListofCollaboratorUser(username, repo, token)
		} else {
			fmt.Println("Provide repository name: --r=<repository>")
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
			fmt.Println(*userList[i].Login, "   ROLE:", *userList[i].RoleName)
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
	usersCmd.Flags().String("r", "", "Repsitory name")
}
