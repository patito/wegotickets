package main

import "github.com/patito/wegotickets/scrape"

func main() {
	s := scrape.New()
	s.FindAllEvents()
	s.PPrint()
}
