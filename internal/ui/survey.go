package ui

import (
	"fmt"
	"mirorim-cli/internal/utils"

	"github.com/AlecAivazis/survey/v2"
)

// PromptProjectDetails prompts the user to select the project type and input a project name.
func PromptProjectDetails() (string, string, error) {
	// Define options for project types
	projectTypes := []string{"Expo-managed app", "Bare React Native app"}

	// Prompt for project type selection
	var projectType string
	prompt := &survey.Select{
		Message: "Select the type of React Native project:",
		Options: projectTypes,
	}
	err := survey.AskOne(prompt, &projectType)
	if err != nil {
		return "", "", fmt.Errorf("failed to get project type: %v", err)
	}

	// Prompt for project name
	var projectName string
	promptName := &survey.Input{
		Message: "Enter the name of your project:",
	}
	err = survey.AskOne(promptName, &projectName, survey.WithValidator(utils.ValidateProjectName))
	if err != nil {
		return "", "", fmt.Errorf("failed to get project name: %v", err)
	}

	// Map project type to internal representation
	if projectType == "Expo-managed app" {
		return "expo", projectName, nil
	} else if projectType == "Bare React Native app" {
		return "bare", projectName, nil
	}
	return "", "", fmt.Errorf("invalid project type selected")
}
