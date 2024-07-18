/*
Copyright © 2024 repoguard harisheoran@protonmail.com
*/
package rg

import (
	"errors"
	"os"
	"os/user"
)

var (
	repoGuardDirPath = ".repoguard"
	tokenFileName    = "credentials"
	usernameFileName = "configs"
)

// load the path of the files
func LoadConfigFilePath(credentialsFileName string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDir := usr.HomeDir
	credentialsFilePath := homeDir + "/" + repoGuardDirPath + "/" + credentialsFileName

	return credentialsFilePath, nil
}

// check if configuration file already exist or not
func IsConfigurationExist(FilePath string) bool {
	_, errExistFile := os.Stat(FilePath)
	if os.IsNotExist(errExistFile) {
		return false
	}
	return true
}

// load saved password from .repoguard/credentials
func LoadToken() (string, error) {
	filePath, err := LoadConfigFilePath(tokenFileName)
	if err != nil {
		return "", err
	}

	if !IsConfigurationExist(filePath) {
		return "", errors.New("configure the repoguard: $ repoguard configure")
	}

	data, errRead := os.ReadFile(filePath)
	if errRead != nil {
		return "", errRead
	}
	token := string(data)
	if len(token) == 0 {
		return "", errors.New("configure the repoguard with correct credentials")
	}

	return token, nil
}

// load username from file .repoguard/configs
func LoadUserName() (string, error) {

	filePath, err := LoadConfigFilePath(usernameFileName)
	if err != nil {
		return "", err
	}

	if !IsConfigurationExist(filePath) {
		return "", errors.New("configure the repoguard: $ repoguard configure")
	}

	data, errRead := os.ReadFile(filePath)
	if errRead != nil {
		return "", errRead
	}
	userName := string(data)
	if len(userName) == 0 {
		return "", errors.New("configure the repoguard with correct credentials")
	}

	return userName, nil

}
