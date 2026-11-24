package storage

import (
	"errors"
)

// Event описывает структуру нашего события.
type Event struct {
	ID    int
	Title string
	Date  string
}

// Storage для хранения событий в памяти.
type Storage struct {
	events []Event
	nextID int
}

func NewStorage() *Storage {
	return &Storage{
		events: make([]Event, 0),
		nextID: 1,
	}
}

func (s *Storage) AddEvent(title, date string) int {
	ev := Event{
		ID:    s.nextID,
		Title: title,
		Date:  date,
	}
	s.events = append(s.events, ev)
	s.nextID++
	return ev.ID
}

func (s *Storage) ListEvents() []Event {
	return s.events
}

func (s *Storage) DeleteEvent(id int) error {
	for i, ev := range s.events {
		if ev.ID == id {
			s.events = append(s.events[:i], s.events[i+1:]...)
			return nil
		}
	}
	return errors.New("событие не найдено")
}
