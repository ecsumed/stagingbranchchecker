package main

import (
	"fmt"
)

func main() {
    args := parseFlags()

    conf := parseConfig(args)

	for _, repoPath := range conf.Repos {
		repo := Repo{path: repoPath}

		fmt.Println("", repo.GetHeadInfo())
	}
}
