package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://positivethinking.udemy.com",
		"https://www.facebook.com",
		"https://www.google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

	}
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)

	// }
	// for l := range c {
	// 	go func(l string) {
	// 		time.Sleep(time.Second * 5)
	// 		checkLink(l, c)
	// 	}(l)
	// }
	for {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
