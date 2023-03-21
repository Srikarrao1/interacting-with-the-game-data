package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	houses := map[string][]string {
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory", "bobo"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "bobo", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	delete(houses, "bobo")

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("Please provide a Hogwards house name\n")
		return
	}

	house, students := args[0], houses[args[0]]

	if students == nil {
		fmt.Printf("Sorry, I don't know anything about %q.\n", house)
		return
	}
	clone := append([]string(nil), students...)
	sort.Strings(clone)


	fmt.Printf("~~~ %s students ~~~\n\n", house)

	for _, student := range clone {
		fmt.Printf("\t+%s\n", student)
	}
}