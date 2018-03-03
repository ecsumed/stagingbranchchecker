package main

import (
    "fmt"
    "gopkg.in/src-d/go-git.v4"
)

type Repo struct {
    path string
}

    
// Return current head info: branch, author,
// message. 
func (r Repo) GetHeadInfo() map[string]string {
    info := make(map[string]string)

	repo, _ := git.PlainOpen(r.path)
	head, _ := repo.Head()
	commit, _ := repo.CommitObject(head.Hash())

    headReference := head.Name()

    if headReference.IsBranch() {
        info["branch"] = headReference.Short()
    } else {
        info["branch"] = fmt.Sprintf("%s (Head detached)", head.Hash())
    }

    info["author"] = commit.Author.Name
    info["message"] = commit.Message

    return info
}
