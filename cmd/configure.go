/*
Copyright © 2024 repoguard harisheoran@protonmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"repoguard/rg"
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
		// Take user input for Github token
		fmt.Printf("GitHub Token: ")
		GITHUB_TOKEN, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		// save the token
		errSaveToken := saveToken(string(GITHUB_TOKEN))
		if errSaveToken != nil {
			fmt.Println(errSaveToken)
			os.Exit(1)
		}
		fmt.Println()

		// Take user input for GitHub account username
		fmt.Printf("Github Username:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		errScan := scanner.Err()
		if errScan != nil {
			log.Fatal(errScan)
		}

		// save the username
		errSaveUserName := saveUsername(scanner.Text())
		if errSaveUserName != nil {
			fmt.Println(errSaveUserName)
			os.Exit(1)
		}
	},
}

func saveToken(GITHUB_TOKEN string) error {
	// Get Home directory
	usr, err := user.Current()
	if err != nil {
		return err
	}

	homeDir := usr.HomeDir
	repoGuardDirPath := ".repoguard"
	credentialsFilePath := homeDir + "/" + repoGuardDirPath

	// Create the directory to store credential file
	errDirCreate := os.Mkdir(credentialsFilePath, 0700)
	if errDirCreate != nil && !os.IsExist(errDirCreate) {
		return err
	}

	// Create credential file inside directory
	var filename = credentialsFilePath + "/credentials"
	_, errCheck := os.Stat(filename)

	if os.IsNotExist(errCheck) {
		_, err := os.Create(filename)

		if err != nil {
			return err
		}
	}

	// Write the credentials file with GitHub Token
	errWrite := os.WriteFile(filename, []byte(GITHUB_TOKEN), 0700)
	if errWrite != nil {
		return errWrite
	}

	return nil
}

func saveUsername(username string) error {
	fileName, err := rg.LoadConfigFilePath("/configs")
	if err != nil {
		return err
	}

	// check if file exist, if not then create file
	if !rg.IsConfigurationExist(fileName) {
		_, err := os.Create(fileName)
		if err != nil {
			return err
		}
	}

	// write the configs file
	errWriteFile := os.WriteFile(fileName, []byte(username), 0644)
	if errWriteFile != nil {
		return errWriteFile
	}
	return nil
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
