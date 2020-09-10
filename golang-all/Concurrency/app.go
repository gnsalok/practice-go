package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://facebook.com",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	//option 1
	// for {
	// 	go checkLink(<-c, c)
	// }

	//better syntax
	for l := range c {
		time.Sleep(time.Second)
		go checkLink(l, c)
	}

	fmt.Println(<-c)

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might down")
		c <- link
		return
	} else {
		fmt.Println(link, "is up")
		c <- link
	}
}
