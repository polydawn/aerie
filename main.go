package main

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	r, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}
	itr, err := r.Branches()
	if err != nil {
		panic(err)
	}
	itr.ForEach(func(ref *plumbing.Reference) error {
		fmt.Printf("%v\n", ref)
		return nil
	})
}

func printRemoteRefs(r git.Repository) {
	// TODO likely you need need auth & ssh config for this to work
	remotes, err := r.Remotes()
	if err != nil {
		panic(err)
	}
	for _, remote := range remotes {
		remoteRefs, err := remote.List(&git.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, ref := range remoteRefs {
			fmt.Printf("%v\n", ref)
		}
	}
}
