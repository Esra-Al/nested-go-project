package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Esra-Al/nested-go-project/submodule"
)

func main() {
	modFileRootPath := "." // Replace with the actual path if needed
	gitTopLevelDir, err := getGitTopLevelDir(modFileRootPath)
	if err != nil {
		fmt.Printf("Error finding Git top-level directory: %v\n", err)
		return
	}

	fmt.Printf("Git top-level directory: %s\n", gitTopLevelDir)

	// Call the submodule function
	submodule.SayHello() // Call the submodule function
}

func getGitTopLevelDir(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return strings.TrimSpace(string(output)), nil
}
