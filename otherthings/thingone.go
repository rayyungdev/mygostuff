package otherthings

import (
	"fmt"
	"sync"
	"time"
)

type thingOne struct {
	LastUpdated time.Time
	Data        map[string]string
}

type MyThingCollection struct {
	NewCollectionName string
	TTL               time.Duration
	Collection        map[string]*thingOne
	mu                sync.RWMutex
}

func StartNewCollection(ttl time.Duration, collectionName string) *MyThingCollection {
	fmt.Println("Starting New Collection called: ", collectionName)
	collection := &MyThingCollection{NewCollectionName: collectionName, TTL: ttl, Collection: make(map[string]*thingOne)}

	return collection
}

func (m *MyThingCollection) Lock() {
	m.mu.Lock()
}

func (m *MyThingCollection) Unlock() {
	m.mu.Unlock()
}

func (m *MyThingCollection) RLock() {
	m.mu.RLock()
}

func (m *MyThingCollection) RUnlock() {
	m.mu.RUnlock()
}

func (m *MyThingCollection) AddToCollectionNoLock(Key string, Data map[string]string) error {
	_, exists := m.Collection[Key]
	if exists {
		return fmt.Errorf("item %v already exists", Key)
	}

	newThing := &thingOne{
		Data:        Data,
		LastUpdated: time.Now().Truncate(time.Second),
	}

	m.Collection[Key] = newThing
	return nil
}

func (m *MyThingCollection) AddToCollectionLocked(Key string, Data map[string]string) error {
	m.Lock()
	defer m.Unlock()
	return m.AddToCollectionNoLock(Key, Data)
}

func (m *MyThingCollection) RetrieveItemNoLock(Key string) (*thingOne, error) {
	item, exist := m.Collection[Key]
	if !exist {
		return nil, fmt.Errorf("item %v does not exist", Key)
	}
	return item, nil
}

func (m *MyThingCollection) RetrieveItemLocked(Key string) (*thingOne, error) {
	m.RLock()
	defer m.RUnlock()
	return m.RetrieveItemNoLock(Key)
}
