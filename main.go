package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: git-tag-cli <tag_name> <tag_message> <remote_url>")
		os.Exit(1)
	}

	tagName := os.Args[1]
	tagMessage := os.Args[2]
	remoteURL := os.Args[3]

	// Create the Git tag
	createTag(tagName, tagMessage)

	// Push the tag to the remote
	pushTag(tagName, remoteURL)

	// Compare with the previous tag
	compareTags(tagName)
}

func createTag(tagName, tagMessage string) {
	cmd := exec.Command("git", "tag", "-a", tagName, "-m", tagMessage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error creating tag: %v", err)
	}
	fmt.Printf("Tag %s created.\n", tagName)
}

func pushTag(tagName, remoteURL string) {
	cmd := exec.Command("git", "push", "origin", tagName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error pushing tag to remote: %v", err)
	}
	fmt.Printf("Tag %s pushed to remote %s.\n", tagName, remoteURL)
}

func compareTags(tagName string) {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags", tagName+"^")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error comparing tags: %v", err)
	}

	previousTag := strings.TrimSpace(string(out))
	cmd = exec.Command("git", "diff", "--name-status", previousTag, tagName)
	out, err = cmd.Output()
	if err != nil {
		log.Fatalf("Error comparing tags: %v", err)
	}

	fmt.Printf("Comparison between %s and %s:\n", previousTag, tagName)
	fmt.Println(string(out))
}
