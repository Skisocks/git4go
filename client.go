package git

import (
	"fmt"
	"os"
	"strings"
)

const (
	DefaultDirWritePermissions = 0766
)

type client struct {
	rootDir  string
	repoUrl  string
	repoName string
	repoDir  string
}

// NewClient returns a new git client
func NewClient(repoUrl string, rootDir string, gitAuth *Auth) *client {
	repoName := GetRepositoryName(repoUrl)
	repoUrlWithAuth := fmt.Sprintf("https://%s@%s", gitAuth.Token, strings.TrimPrefix(repoUrl, "https://"))
	repoDir := fmt.Sprintf("%s/%s", rootDir, repoName)

	return &client{
		rootDir:  rootDir,
		repoUrl:  repoUrlWithAuth,
		repoName: repoName,
		repoDir:  repoDir,
	}
}

// Clone clones the remote repository into a folder of the same name in the given root directory
func (c *client) Clone() error {
	if err := os.MkdirAll(c.rootDir, DefaultDirWritePermissions); err != nil {
		return fmt.Errorf("failed to create unique directory for '%s': %v", c.rootDir, err)
	}

	_, err := RunCommand(c.rootDir, "clone", c.repoUrl, c.repoName)
	if err != nil {
		return fmt.Errorf("failed to clone repo: %v", err)
	}
	return nil
}

// Pull updates the local repository with any changes made in the remote
func (c *client) Pull() error {
	_, err := RunCommand(c.repoDir, "pull")
	if err != nil {
		return fmt.Errorf("failed to pull in dir %s: %v", c.repoDir, err)
	}
	return nil
}

// Add adds a file to staging
func (c *client) Add(filePath string) error {
	_, err := RunCommand(c.repoDir, "add", filePath)
	if err != nil {
		return fmt.Errorf("failed to add changes for %s: %v", filePath, err)
	}
	return nil
}

// AddAll adds all the changes in the local repository to staging
func (c *client) AddAll() error {
	_, err := RunCommand(c.repoDir, "add", "--all")
	if err != nil {
		return fmt.Errorf("failed to add all changes: %v", err)
	}
	return nil
}

// IsLocalUpToDate returns a boolean based on whether the local repository
// is up-to-date with the remote
func (c *client) IsLocalUpToDate() bool {
	status, _ := RunCommand(c.repoDir, "status", "-s", c.repoDir)
	status = strings.TrimSpace(status)
	return len(status) == 0
}

// Commit commits any changes made in the local repository with the given message
func (c *client) Commit(message string) error {
	_, err := RunCommand(c.repoDir, "commit", "-m", message)
	if err != nil {
		return fmt.Errorf("failed to commit changes: %v", err)
	}
	return nil
}

// Reset resets the local repository by a given number of commits
func (c *client) Reset(previousHead int) error {
	_, err := RunCommand(c.repoDir, "reset", "--soft", fmt.Sprintf("HEAD~%d", previousHead))
	if err != nil {
		return fmt.Errorf("failed to reset: %v", err)
	}
	return nil
}

// Push pushes any committed changes in the local repository to the remote
func (c *client) Push() error {
	_, err := RunCommand(c.repoDir, "push")
	if err != nil {
		return fmt.Errorf("failed to push changes: %v", err)
	}
	return nil
}

// GetRepositoryName returns the name of the remote repository
func (c *client) GetRepositoryName() string { return c.repoName }

// GetRepositoryName takes the URL of a git repository and returns its name.
func GetRepositoryName(repoUrl string) string {
	url := strings.TrimPrefix(repoUrl, "https://")
	url = strings.TrimSuffix(url, ".git")
	splitUrl := strings.Split(url, "/")
	return splitUrl[len(splitUrl)-1]
}
