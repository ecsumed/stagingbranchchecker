package main

import (
	"fmt"
)

func main() {
    args := parseFlags()

    conf := parseConfig(args)

	for _, repoPath := range conf.Repos {
		repo := Repo{path: repoPath}

		fmt.Println("Branch: ", repo.GetHeadInfo()["branch"])
		fmt.Println("author: ", repo.GetHeadInfo()["author"])
		fmt.Println("message: ", repo.GetHeadInfo()["message"])
	}
}
