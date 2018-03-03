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

	hi := HeadInfo{Branch: "None", Author: "None", Msg: "None"}

	repo, err := git.PlainOpen(r.path)

	if err != nil {
        hi.Branch = r.path
        hi.Author = err.Error()
        hi.Msg = err.Error()
		return &hi
	}

	head, _ := repo.Head()
	commit, _ := repo.CommitObject(head.Hash())

    headReference := head.Name()


    if headReference.IsBranch() {
        hi.Branch = headReference.Short()
    } else {
        hi.Branch = fmt.Sprintf("%s (Head detached)", head.Hash())
    }

    hi.Author = commit.Author.Name
    hi.Msg = commit.Message

    return &hi
}
