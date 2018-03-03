package main

import (
	"fmt"
)

func main() {
	repoPath := "/tmp/test"

    repo := Repo{path: repoPath}

    fmt.Print("%s", repo.GetHeadInfo())
}
