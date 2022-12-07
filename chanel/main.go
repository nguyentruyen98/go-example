package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.facebook.com",
		"https://www.google.com",
		"https://positivethinking.udemy.com",
	}

	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
