package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	var r, err = git.PlainOpen("./")
	if err != nil {
		log.Fatal(err)
	}
	tIter, err := r.Tags()
	if err != nil {
		log.Fatal(err)
	}
	tags := []string{}
	err = tIter.ForEach(func(t *plumbing.Reference) error {
		tags = append(tags, t.Name().Short())
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tags)
}
