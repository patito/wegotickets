package scrape

import (
	"encoding/json"
	"fmt"
	"log"
)

// Print the JSON data
func (s *Scrape) PPrint() {

	b, err := json.MarshalIndent(s.Events, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
}
