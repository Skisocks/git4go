package main

import (
	git "personal/git4go"
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
}
