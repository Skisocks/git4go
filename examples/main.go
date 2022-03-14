package main

import (
	git "github.com/Skisocks/git4go"
)

func main() {
	gitAuth := git.NewAuth("Username", "email@company.com", "token")
	if err := gitAuth.SetGlobalCredentials(); err != nil {
		panic(err)
	}
	gitClient := git.NewClient("https://githube,com/Skisocks/git4go.git", "tmp/git/", gitAuth)

	if err := gitClient.Clone(); err != nil {
		panic(err)
	}

	if err := gitClient.Add("tmp/git/git4go/README.md"); err != nil {
		panic(err)
	}

	if err := gitClient.Commit("Update README"); err != nil {
		panic(err)
	}

	if err := gitClient.Push(); err != nil {
		panic(err)
	}
}
