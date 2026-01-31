package detector

import "strings"

type SpamDetector struct {
	keywords []string
}

func NewSpamDetector(keywords []string) *SpamDetector {
	return &SpamDetector{
		keywords: keywords,
	}
}

func (d *SpamDetector) IsSpam(message string) bool {
	msg := strings.ToLower(message)

	for _, k := range d.keywords {
		if strings.Contains(msg, k) {
			return true
		}
	}
	return false
}
