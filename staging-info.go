package main

import (
	"fmt"
	git "gopkg.in/src-d/go-git.v4"
)

// type Clock struct {
//     hour int
//     minute int
// }
//
// func New(hour, minute int) Clock {
//     var minutes int
//     var hours int
//
//     minutes = 60 + minute
//     minutes = minutes % 60
//
//     hours = 24 + hour
//     hours = hours + (minute / 60)
//     hours = hours % 24
//
//     return Clock{hour: hours, minute: minutes}
// }
//
// func (c Clock) String() string {
//     return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
// }
//
// func (c Clock) Add(minutes int) Clock {
//     c.minute = c.minute + minutes
//
//     return Clock{hour: c.hour, minute: c.minute}
// }

func main() {
	repo_path := "/home/fahad/Workspace/go/src/staging-info"
	repo, _ := git.PlainOpen(repo_path)

	head, _ := repo.Head()

	commit, _ := repo.CommitObject(head.Hash())

    fmt.Printf("Current Head: %s\n", head.Name().Short())
    fmt.Printf("Commit Author: %s\n", commit.Author.Name)
    fmt.Printf("Commit Message: %s\n", commit.Message)

	// worktree, _ := repo.Worktree()
	// wt_status, _ := worktree.Status()
	// fmt.Printf("%s", wt_status)
}
