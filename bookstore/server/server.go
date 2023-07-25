package server

import (
	"bookstore/server/middleware"
	"bookstore/store"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type BookStoreServer struct {
	s   store.Store
	srv *http.Server
}

func (bs *BookStoreServer) createBookHandler(writer http.ResponseWriter, request *http.Request) {
	dec := json.NewDecoder(request.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bs.s.Create(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (bs *BookStoreServer) getBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "missing id", http.StatusBadRequest)
		return
	}

	book, err := bs.s.Get(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response(writer, book)
}

func (bs *BookStoreServer) updateBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "missing id", http.StatusBadRequest)
		return
	}

	doc := json.NewDecoder(request.Body)
	var book store.Book
	if err := doc.Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	book.Id = id
	if err := bs.s.Update(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (bs *BookStoreServer) getAllBooksHandler(writer http.ResponseWriter, request *http.Request) {
	books, err := bs.s.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	response(writer, books)
}

func (bs *BookStoreServer) delBookHandler(writer http.ResponseWriter, request *http.Request) {

	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "missing id", http.StatusBadRequest)
		return
	}

	if err := bs.s.Delete(id); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func response(writer http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)
	if err != nil {
		log.Fatal("write response error: ", err)
	}

}

// NewBookStoreServer 接收一个接口类型，返回一个具体类型
func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	srv := &BookStoreServer{
		s: s,
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.updateBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.getBookHandler).Methods("GET")
	router.HandleFunc("/books", srv.getAllBooksHandler).Methods("GET")
	router.HandleFunc("/book/{id}", srv.delBookHandler).Methods("DELETE")

	srv.srv.Handler = middleware.Logging(middleware.Validating(router))
	return srv
}

func (bs *BookStoreServer) ListenAndServe() (<-chan error, error) {
	var err error
	errChan := make(chan error)
	go func() {
		err = bs.srv.ListenAndServe()
		errChan <- err
	}()

	select {
	case err := <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (bs *BookStoreServer) Shutdown(ctx context.Context) error {
	return bs.srv.Shutdown(ctx)
}
