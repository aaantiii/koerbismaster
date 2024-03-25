package koerbismaster

import (
	"os"
	"strconv"
	"sync"
	"time"
)

const statsFile = "./stats.csv"

var stats = &Stats{}

type Stats struct {
	Timestamp time.Time `json:"timestamp"`
	Previous  time.Time `json:"previous"`
	Messages  int       `json:"messages"`
	mu        sync.Mutex
}

// IncMessages increments the message counter.
func (s *Stats) IncMessages() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Messages++
}

// Reset prepares the stats for the next interval.
func (s *Stats) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Previous = s.Timestamp
	s.Messages = 0
}

func (s *Stats) Save() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Timestamp = time.Now()
	if s.Previous.IsZero() {
		return nil // Skip first save as there is no previous data.
	}

	file, err := openStatsFile()
	if err != nil {
		return err
	}
	defer file.Close()
	csv := s.Timestamp.Format(time.RFC3339) + "," + s.Previous.Format(time.RFC3339) + "," + strconv.Itoa(s.Messages) + "\n"
	_, err = file.WriteString(csv)
	return err
}

func openStatsFile() (*os.File, error) {
	var file *os.File
	var err error
	if _, err = os.Stat(statsFile); os.IsNotExist(err) {
		file, err = os.Create(statsFile)
		if err != nil {
			return nil, err
		}
		_, err = file.WriteString("timestamp,previous,messages\n")
	} else {
		file, err = os.OpenFile(statsFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	return file, err
}
