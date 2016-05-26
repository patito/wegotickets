package scrape

import (
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/patito/wegotickets/event"
	"github.com/patito/wegotickets/html"
)

type Scrape struct {
	Events []event.Event
	Doc    *goquery.Document
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
	}
}

func getInfo(doc *goquery.Document, tag string) string {
	info := ""
	doc.Find(tag).Each(func(i int, s *goquery.Selection) {
		info = s.Text()
	})

	return info
}

func (s *Scrape) parseAndSave(doc *goquery.Document) {

	artist := getInfo(doc, artistAttr)
	if strings.Contains(strings.ToUpper(artist), "COMEDY") {
		return
	}

	e := event.Event{}

	e.Artist = artist
	e.Date = getInfo(doc, dateAttr)
	venueAndCity := getInfo(doc, venueAttr)
	sliceVenueCity := strings.Split(venueAndCity, ":")
	if len(sliceVenueCity) > 1 {
		e.Venue = sliceVenueCity[1]
	} else {
		e.Venue = sliceVenueCity[0]
	}
	e.City = sliceVenueCity[0]
	e.Price = doc.Find(html.TagStrong).First().Text()

	log.WithFields(log.Fields{"Event": e}).Info("Adding new event")
	s.Events = append(s.Events, e)
}

func (s *Scrape) getConcertInfo(link string) {

	doc, err := goquery.NewDocument(link)
	if err != nil {
		log.Fatal(err)
	}

	s.parseAndSave(doc)
}

func (s *Scrape) getConcertLink(i int, qs *goquery.Selection) {

	if link, ok := qs.Find("h2 a").Attr(html.TagHREF); ok {
		s.getConcertInfo(link)
	}
}

func (s *Scrape) FindAllEvents() {
	s.Doc.Find(linksAttr).Each(s.getConcertLink)
}

func New() *Scrape {
	doc, err := goquery.NewDocument(os.Getenv("URL"))
	if err != nil {
		log.Fatal(err)
	}
	return &Scrape{[]event.Event{}, doc}
}
