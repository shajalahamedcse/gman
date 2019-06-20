package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/shajalahamedcse/gman/credentials"
)

const (
	activeIndicator = " -> "
)

func main() {

	// Parse AWS Credentials
	profiles, err := credentials.Parse()
	if err != nil {
		fmt.Println("Parsing credentials: ", err)
		os.Exit(1)
	}
	fmt.Println(profiles)

	// Check the command
	if len(os.Args) == 1 {
		list(profiles)
	}
}

func list(profiles []credentials.Profile) {

	sort.Slice(profiles, func(i, j int) bool {
		return profiles[i].Name < profiles[j].Name
	})

	for _, profile := range profiles {
		fmt.Printf("%v\n", profile.String())
	}
}
