/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure repoguard for use",
	Long:  `Configure the repoguard by providing it the GitHub Token`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("GitHub Token: ")
		// fmt.Scanln(&GITHUB_TOKEN)
		GITHUB_TOKEN, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("Password", string(bytePassword))
		saveToken(string(GITHUB_TOKEN))
	},
}

func saveToken(GITHUB_TOKEN string) {
	// Get Home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Filesystemt user not found")
	}
	homeDir := usr.HomeDir
	repoGuardDirPath := ".repoguard"
	credentialsFilePath := homeDir + "/" + repoGuardDirPath

	// Create the directory to store credential file
	errDirCreate := os.Mkdir(credentialsFilePath, 0700)
	if errDirCreate != nil && !os.IsExist(errDirCreate) {
		log.Fatal(err)
	}

	// Create credential file inside directory
	var filename = credentialsFilePath + "/credentials"
	_, errCheck := os.Stat(filename)

	if os.IsNotExist(errCheck) {
		_, err := os.Create(filename)

		if err != nil {
			log.Fatal(err)
		}
	}

	// Write the credentials file with GitHub Token
	errWrite := os.WriteFile(filename, []byte(GITHUB_TOKEN), 0700)
	if errWrite != nil {
		log.Fatal("Error Saving Token")
	}
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
