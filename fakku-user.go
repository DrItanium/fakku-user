// prints out user information
package main

import (
	"flag"
	"fmt"
	"github.com/DrItanium/fakku"
	"time"
)

var userName = flag.String("username", "", "The username to lookup")

func main() {
	flag.Parse()
	//TODO: add support for grabbing beyond the first page
	//TODO: add support for spitting out CLIPS style facts (template assertions
	//or instances)
	if *userName == "" {
		fmt.Printf("ERROR: no username specified")
	} else {
		user, err := fakku.GetUserProfile(*userName)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Something bad happened! Perhaps fakku is down?")
		} else {
			fmt.Printf("Username: %s\nUrl: %s\nRank: %s\n", user.Username, user.Url, user.Rank)
			fmt.Printf("Avatar\n\tURL: %s\n\tWidth: %d\n\tHeight: %d\n", user.Avatar, user.AvatarWidth, user.AvatarHeight)
			regDate, err0 := time.Parse(time.UnixDate, string(user.RegistrationDate))
			if err0 != nil {
				fmt.Println(err0)
				return
			}
			lastVisit, err1 := time.Parse(time.UnixDate, string(user.LastVisit))
			if err1 != nil {
				fmt.Println(err1)
				return
			}
			fmt.Printf("Registration date: %t\nLastVisit: %t\n", regDate, lastVisit)
		}

	}
	/*
		posts, err := fakku.GetUserProfile(*userName)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Something bad happened! Perhaps Fakku is down?")
		}
		for i := 0; i < int(posts.Total); i++ {
			tmp := posts.Index[i]
			switch tmp.(type) {
			case fakku.Content:
				content := tmp.(fakku.Content)
				tags := content.Tags
				fmt.Printf("%s - ", content.Name)
				// print out the tags one after another in a form that can be easily
				// grepped through
				if len(tags) == 0 {
					fmt.Printf("No tags!")
				} else {
					fmt.Printf("{ %s", tags[0].Attribute)
					for j := 1; j < len(tags); j++ {
						fmt.Printf(", %s", tags[j].Attribute)
					}
					fmt.Printf(" }")
				}
				fmt.Printf(" - %s", content.Url)
				fmt.Println()
			}
		}
	*/
}
