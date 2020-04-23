package main

import (
	"fmt"
	"os"
	week3 "week_3"
)

func main() {
	fb := new(week3.Facebook)
	err := export(fb, "fbdata.xml")
	twtr := new(week3.Twitter)
	err = export(twtr, "twtrdata.xml")

	if err != nil {
		panic(err)
	}
}

//ScrollFeeds prints all social media feeds
func ScrollFeeds(platforms ...week3.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("==============================")
	}
}

func export(u week3.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return err
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}
