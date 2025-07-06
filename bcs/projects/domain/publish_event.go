package domain

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/oklog/ulid/v2"
)

type Event struct {
	EventID  string
	Metadata struct {
		PublishedAt time.Time
	}
	EventType string
	Data      string
}

func PublishEvent(searchString string, eventType string, eventData string) error {
	path, err := ProjectPath(searchString)
	if err != nil {
		return err
	}

	events, err := readEventStore(path)
	if err != nil {
		return err
	}

	events = append(events, Event{
		EventID: ulid.Make().String(),
		Metadata: struct{ PublishedAt time.Time }{
			PublishedAt: time.Now(),
		},
		EventType: eventType,
		Data:      eventData,
	})

	return writeEventStore(path, events)
}

func readEventStore(projectPath string) ([]Event, error) {
	events := make([]Event, 0)

	eventStorePath := filepath.Join(projectPath, "event_store.json")
	jsonFile, err := os.Open(eventStorePath)
	if err != nil {
		return events, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &events)

	return events, nil
}

func writeEventStore(projectPath string, events []Event) error {
	eventStorePath := filepath.Join(projectPath, "event_store.json")

	jsonString, err := json.Marshal(events)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(eventStorePath, jsonString, os.ModePerm)
}
