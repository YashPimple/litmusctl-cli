/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package config

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Account struct {
	Username string
	Password string
}

var accounts = make(map[string]Account)

func showCommandOptions(commands []string) string {
	prompt := promptui.Select{
		Label: "Litmus config",
		Items: commands,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "It manages multiple ChaosCenter accounts within a system",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		availableCommands := []string{
			"set-account",
			"use-account",
			"get-account",
			"file",
			"view",
		}

		selectedCommand := showCommandOptions(availableCommands)

		fmt.Printf("You selected: %s\n", selectedCommand)

		switch selectedCommand {
		case "set-account":

			promptUsername := promptui.Prompt{
				Label: "Enter username:",
			}
			username, err := promptUsername.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			promptPassword := promptui.Prompt{
				Label: "Enter password:",
				Mask:  '*',
			}
			password, err := promptPassword.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			accounts["default"] = Account{

				Username: username,
				Password: password,
			}
			fmt.Println("Account settings saved!")

		case "get-account":
			if len(accounts) == 0 {
				fmt.Println("No account settings found.")
				return
			}

			fmt.Println("Available accounts:")
			for name := range accounts {
				fmt.Println(name)
			}

			prompt := promptui.Prompt{
				Label: "Enter account name:",
			}
			accountName, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			account, found := accounts[accountName]
			if !found {
				fmt.Println("Account not found.")
				return
			}

			fmt.Printf("Account Settings:\nUsername: %s\nPassword: %s\n", account.Username, account.Password)
		}

	},
}

func init() {

}
