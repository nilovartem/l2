package cache

import (
	"fmt"
	"l2/develop/dev11/event"
	"sync"
	"time"
)

type Cache struct {
	Events map[int][]event.Event
	Mutex  sync.RWMutex
}

func (cache *Cache) CreateEvent(e *event.Event) error {
	cache.Mutex.Lock()
	defer cache.Mutex.Unlock()
	if events, ok := cache.Events[e.UserID]; ok {
		for _, event := range events {
			if event.EventID == e.EventID {
				return fmt.Errorf("event already exist")
			}
		}
	}
	cache.Events[e.UserID] = append(cache.Events[e.UserID], *e)
	return nil
}

func (cache *Cache) UpdateEvent(event *event.Event) error {
	cache.Mutex.Lock()
	defer cache.Mutex.Unlock()
	ind := -1
	events, ok := cache.Events[event.UserID]
	if !ok {
		return fmt.Errorf("user does not find")
	}

	for i, event := range events {
		if event.EventID == event.EventID {
			ind = i
			break
		}
	}
	if ind == -1 {
		return fmt.Errorf("event does not exist")
	}
	cache.Events[event.UserID][ind] = *event

	return nil
}

func (cache *Cache) DeleteEvent(userID, eventID int) (*event.Event, error) {
	cache.Mutex.Lock()
	defer cache.Mutex.Unlock()
	ind := -1
	events, ok := cache.Events[userID]
	if !ok {
		return nil, fmt.Errorf("user does not exist")
	}
	for i, event := range events {
		if event.EventID == eventID {
			ind = i
			break
		}
	}
	if ind == -1 {
		return nil, fmt.Errorf("event does not exist")
	}
	eventsLength := len(cache.Events[userID])
	deletedEvent := cache.Events[userID][ind]
	cache.Events[userID][ind] = cache.Events[userID][eventsLength-1]
	cache.Events[userID] = cache.Events[userID][:eventsLength-1]

	return &deletedEvent, nil
}

func (cache *Cache) GetEventsForDay(userID int, date time.Time) ([]event.Event, error) {
	cache.Mutex.RLock()
	defer cache.Mutex.RUnlock()

	var result []event.Event

	events, ok := cache.Events[userID]

	if !ok {
		return nil, fmt.Errorf("user does not exist")
	}
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			result = append(result, event)
		}
	}
	return result, nil
}

func (cache *Cache) GetEventsForWeek(userID int, date time.Time) ([]event.Event, error) {
	cache.Mutex.RLock()
	defer cache.Mutex.RUnlock()
	result := make([]event.Event, 0)
	events, ok := cache.Events[userID]
	if !ok {
		return nil, fmt.Errorf("user does not exist")
	}
	for _, event := range events {
		eventYear, eventWeek := event.Date.ISOWeek()
		currentYear, currentWeek := date.ISOWeek()
		if eventYear == currentYear && eventWeek == currentWeek {
			result = append(result, event)
		}
	}

	return result, nil
}

func (cache *Cache) GetEventsForMonth(userID int, date time.Time) ([]event.Event, error) {
	cache.Mutex.RLock()
	defer cache.Mutex.RUnlock()
	result := make([]event.Event, 0)
	events, ok := cache.Events[userID]
	if !ok {
		return nil, fmt.Errorf("user does not exist")
	}
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			result = append(result, event)
		}
	}
	return result, nil
}
