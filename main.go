package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/manifoldco/promptui"
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
	vs := []*semver.Version{}
	err = tIter.ForEach(func(t *plumbing.Reference) error {
		v, err := semver.NewVersion(t.Name().Short())
		if err == nil {
			vs = append(vs, v)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	if semver.Collection(vs).Len() == 0 {
		fmt.Println("No versions")
		os.Exit(0)
	}
	sort.Sort(sort.Reverse(semver.Collection(vs)))
	fmt.Println("Latest version:", vs[0])
	prompt := promptui.Select{
		Label: "Select next version",
		Items: []string{"Patch", "Minor", "Major"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	switch result {
	case "Patch":
		v := vs[0].IncPatch()
		fmt.Println(v)
		createTag(*r, v.String())
	case "Minor":
		v := vs[0].IncMinor()
		fmt.Println(v)
		createTag(*r, v.String())
	case "Major":
		v := vs[0].IncMajor()
		fmt.Println(v)
		createTag(*r, v.String())
	}
	os.Exit(0)
}

func createTag(r git.Repository, tag string) {
	hRef, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	ref, err := r.CreateTag("v"+tag, hRef.Hash(), &git.CreateTagOptions{
		Message: tag,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ref)
}
