package main

import (
    "fmt"
    "gopkg.in/src-d/go-git.v4"
)

type Repo struct {
    path string
}

type HeadInfo struct {
    Branch string
    Author string
    Msg string
}
    
// Return current head info: branch, author,
// message. 
func (r Repo) GetHeadInfo() *HeadInfo {

	repo, _ := git.PlainOpen(r.path)
	head, _ := repo.Head()
	commit, _ := repo.CommitObject(head.Hash())

    headReference := head.Name()

	hi := HeadInfo{Branch: "None", Author: "None", Msg: "None"}

    if headReference.IsBranch() {
        hi.Branch = headReference.Short()
    } else {
        hi.Branch = fmt.Sprintf("%s (Head detached)", head.Hash())
    }

    hi.Author = commit.Author.Name
    hi.Msg = commit.Message

    return &hi
}
