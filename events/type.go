package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
}

type Event struct {
}
