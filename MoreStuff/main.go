package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	links := []string{
		"http://github.com/lordtt13",
		"http://fitgirl-repacks.site",
		"https://www.facebook.com",
		"https://www.youtube.com",
	}

	c := make(chan string)

	for _,i := range links{
		go checkLink(i, c)
	}

	for i := 0; i < len(links) ; i++{
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string)  {
	_, err := http.Get(link)
	if err!= nil{
		fmt.Println("Encountered Error", err)
		os.Exit(1)

		c <- "Link might be down"
	}

	c <- "The link's up"
}