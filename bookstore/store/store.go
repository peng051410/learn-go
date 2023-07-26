package store

import "errors"

var (
	ErrorNotFound = errors.New("not found")
	ErrExist      = errors.New("exist")
)

// Book 要存储的实体
type Book struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Press   string   `json:"press"`
}

// Store 存储层的接口
type Store interface {
	Create(*Book) error
	Update(*Book) error
	Get(string) (Book, error)
	GetAll() ([]Book, error)
	Delete(string) error
}
