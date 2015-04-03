// prints out user information
package main

import (
	"flag"
	"fmt"
	"github.com/DrItanium/fakku"
)

var userName = flag.String("username", "", "The username to lookup")

func main() {
	//TODO: file bug report for the fact that user has uid attached to the end of it :(
	flag.Parse()
	if *userName == "" {
		fmt.Println("ERROR: no username specified")
	} else {
		user, err := fakku.GetUser(*userName)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Something bad happened! Perhaps fakku is down?")
		} else {
			fmt.Printf("Username: %s\nUrl: %s\nRank: %s\n", user.Username, user.Url, user.Rank)
			//TODO: file bug report that avatar link isn't working :(
			fmt.Printf("Avatar\n\tURL: %s\n\tWidth: %d\n\tHeight: %d\n", user.Avatar, user.AvatarWidth, user.AvatarHeight)
			fmt.Printf("Registration date: %s\nLast visit: %s\n", user.RegistrationDate(), user.LastVisit())
			fmt.Printf("Subscription count: %d\nNumber of posts: %d\nNumber of topics: %d\n", user.Subscribed, user.Posts, user.Topics)
			//TODO: file bug report about user_timezone not existing
			fmt.Printf("Number of comments: %d\nSignature: \"%s\"\n", user.Comments, user.Signature)
			fmt.Printf("Reputation\n\tForum: %d\n\tComment: %d\n", user.ForumReputation, user.CommentReputation)
			//var hasGold, isOnline string
			fmt.Printf("Has Fakku Gold? %s\nIs online? %s\n", YesNo(user.Gold()), YesNo(user.Online()))
			favs, err0 := user.Favorites()
			if err0 != nil {
				fmt.Println(err0)
				return
			}
			fmt.Println("Favorites")
			for _, element := range favs.Favorites {
				url, err1 := element.Url()
				if err1 != nil {
					fmt.Println(err1)
					return
				}
				fmt.Printf("\t%s - %s\n", element.Name, url)
			}
		}
	}
}
func YesNo(result bool) string {
	if result {
		return "yes"
	} else {
		return "no"
	}
}
