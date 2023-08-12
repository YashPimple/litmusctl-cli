/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type Project struct {
	Name string
}

var projects = make(map[string]Project)

func showCommandOptions(commands []string) string {
	prompt := promptui.Select{
		Label: "Project Manager",
		Items: commands,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources for LitmusChaos agent plane",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		availableCommands := []string{
			"create-project",
		}

		selectedCommand := showCommandOptions(availableCommands)

		switch selectedCommand {
		case "create-project":
			validate := func(input string) error {
				if len(input) < 3 {
					return fmt.Errorf("Project name should have at least 3 characters")
				}
				return nil
			}

			templates := &promptui.PromptTemplates{
				Prompt:  "{{ . }} ",
				Valid:   "{{ . | green }} ",
				Invalid: "{{ . | red }} ",
				Success: "{{ . | bold }} ",
			}

			promptName := promptui.Prompt{
				Label:     "Enter project name",
				Validate:  validate,
				Templates: templates,
			}

			name, err := promptName.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				os.Exit(1)
			}

			projects[name] = Project{
				Name: name,
			}
			fmt.Println("Project created:", name)
		}
	},
}

func init() {
}
