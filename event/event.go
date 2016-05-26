package event

// Event type stores concert information
type Event struct {
	Artist string `json:"artist"`
	City   string `json:"city"`
	Venue  string `json:"venue"`
	Date   string `json:"date"`
	Price  string `json:"price"`
}

// New creates new Event object
func New() *Event {
	return &Event{}
}
