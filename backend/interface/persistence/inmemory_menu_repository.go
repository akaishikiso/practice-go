package persistence

import (
	"sync"

	"practice-go/backend/domain/menu"
)

type InMemoryMenuRepo struct {
	data   []*menu.Menu
	nextID int
	mu     sync.Mutex
}

func NewInMemoryMenuRepo() *InMemoryMenuRepo {
	return &InMemoryMenuRepo{nextID: 1}
}

func (r *InMemoryMenuRepo) Save(m *menu.Menu) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	m.ID = r.nextID
	r.nextID++
	r.data = append(r.data, m)
	return nil
}

func (r *InMemoryMenuRepo) FindAll() ([]*menu.Menu, error) {
	return r.data, nil
}

func (r *InMemoryMenuRepo) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, m := range r.data {
		if m.ID == id {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

var ErrNotFound = error(nil)