package store

import (
	mystore "bookstore/store"
	"bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (m *MemStore) Create(book *mystore.Book) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[book.Id]; ok {
		return mystore.ErrExist
	}

	nBook := *book
	m.books[book.Id] = &nBook

	return nil
}

func (m *MemStore) Update(book *mystore.Book) error {
	m.Lock()
	defer m.Unlock()

	oldBook, ok := m.books[book.Id]
	if !ok {
		return mystore.ErrorNotFound
	}

	nBook := *oldBook
	if book.Name != "" {
		nBook.Name = book.Name
	}
	if book.Authors != nil {
		nBook.Authors = book.Authors
	}
	if book.Press != "" {
		nBook.Press = book.Press
	}

	m.books[book.Id] = &nBook

	return nil
}

func (m *MemStore) Get(s string) (mystore.Book, error) {
	m.Lock()
	defer m.Unlock()

	book, ok := m.books[s]
	if ok {
		return *book, nil
	}

	return mystore.Book{}, mystore.ErrorNotFound
}

func (m *MemStore) GetAll() ([]mystore.Book, error) {
	m.Lock()
	defer m.Unlock()

	allBooks := make([]mystore.Book, 0, len(m.books))
	for _, book := range m.books {
		allBooks = append(allBooks, *book)
	}
	return allBooks, nil
}

func (m *MemStore) Delete(id string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[id]; !ok {
		return mystore.ErrorNotFound
	}

	delete(m.books, id)
	return nil
}
