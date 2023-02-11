package main

// https://github.com/go-git/go-git/tree/master/_examples

import (
	"context"
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	// . "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	url := "https://github.com/src-d/go-siva"
	// url = "https://github.com/go-git/go-git"
	url = "https://github.com/go-git/go-billy"
	url = "https://github.com/BrotherSam66/goreleaser"
	fmt.Println("################## cloneContext(url)")
	cloneContext(url)
	fmt.Println("################## clone(url)")
	// clone(url)

}

// Example of how to:
// - Clone a repository into memory
// - Get the HEAD reference
// - Using the HEAD reference, obtain the commit this reference is pointing to
// - Using the commit, obtain its history and print it
func clone(url string) {
	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	// Info("git clone https://github.com/src-d/go-siva")
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})
	// CheckIfError(err)
	if err != nil {
		fmt.Println(" error: ", err)
	}

	// Gets the HEAD history from HEAD, just like this command:
	// Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	// CheckIfError(err)
	if err != nil {
		fmt.Println(" error: ", err)
	}

	// ... retrieves the commit history
	since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2019, 7, 30, 0, 0, 0, 0, time.UTC)
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	// CheckIfError(err)
	if err != nil {
		fmt.Println(" error: ", err)
	}

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	// CheckIfError(err)
	if err != nil {
		fmt.Println(" error: ", err)
	}
}

func cloneContext(url string) {
	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	// Info("git clone https://github.com/src-d/go-siva")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*11)
	defer cancel()
	r, err := git.CloneContext(ctx, memory.NewStorage(), nil, &git.CloneOptions{
		URL: url,
	})
	// CheckIfError(err)
	if err != nil {
		fmt.Println(" error: ", err)
	}

	// Gets the HEAD history from HEAD, just like this command:
	// Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	// CheckIfError(err)
	if err != nil {
		fmt.Println(" error: ", err)
	}

	// ... retrieves the commit history
	since := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Date(2019, 7, 30, 0, 0, 0, 0, time.UTC)
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	// CheckIfError(err)
	if err != nil {
		fmt.Println(" error: ", err)
	}

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	// CheckIfError(err)
	if err != nil {
		fmt.Println(" error: ", err)
	}
}
