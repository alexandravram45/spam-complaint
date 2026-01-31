package ingestor

type Ingestor struct {
	messages []string
}

func NewIngestor() *Ingestor {
	return &Ingestor{
		messages: []string{
			"Win money now!!!",
			"Hello, how are you?",
			"FREE prize waiting for you",
		},
	}
}

func (i *Ingestor) Fetch() []string {
	return i.messages
}
