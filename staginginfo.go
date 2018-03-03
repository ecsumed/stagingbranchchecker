package main

import (
	"fmt"
)

func main() {
    args := parseFlags()


    fmt.Print("Conf File: %s", args["confFile"])

	repoPath := "/tmp/test"

    repo := Repo{path: repoPath}

    fmt.Print("%s", repo.GetHeadInfo())
}
