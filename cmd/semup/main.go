package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/semver/v3"
	"github.com/fuku710/semup"
	"github.com/go-git/go-git/v5"
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

	vs, err := semup.ListVersions(r)
	if err != nil {
		log.Fatal(err)
	}

	if semver.Collection(vs).Len() == 0 {
		fmt.Println("No versions")
		os.Exit(0)
	}

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
