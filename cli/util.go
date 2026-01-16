package cli

import (
	"fmt"
)

func prompt(out *Output, message string) (string, error) {
	out.Printf("%s: ", message)
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		if err.Error() == "unexpected newline" {
			return "", nil
		}
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	return response, nil
}

func confirmPrompt(out *Output, message string, defaultYes bool) (bool, error) {
	var fullMessage string
	if defaultYes {
		fullMessage = fmt.Sprintf("%s (Y/n)", message)
	} else {
		fullMessage = fmt.Sprintf("%s (y/N)", message)
	}

	response, err := prompt(out, fullMessage)
	if err != nil {
		return false, err
	}

	if defaultYes {
		if response != "y" && response != "Y" && response != "" {
			return false, nil
		}
		return true, nil
	} else {
		if response != "n" && response != "N" && response != "" {
			return true, nil
		}
		return false, nil
	}
}
