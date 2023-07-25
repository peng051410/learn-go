package store

import (
	"bookstore/store"
	"bookstore/store/factory"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"sync"
)

func init() {
	factory.Register("redis", &RedisStore{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	})
}

var ctx = context.Background()

type RedisStore struct {
	sync.RWMutex
	client *redis.Client
}

func (r *RedisStore) Create(book *store.Book) error {

	r.Lock()
	defer r.Unlock()

	_, err := r.client.HGet(ctx, "book", book.Id).Result()
	if err == redis.Nil {
		data, err := json.Marshal(book)
		if err != nil {
			return err
		}

		sErr := r.client.HSet(ctx, "book", book.Id, string(data))
		if sErr != nil {
			return sErr.Err()
		}

		return nil
	} else if err != nil {
		return err
	}
	return store.ErrExist
}

func (r *RedisStore) Update(book *store.Book) error {
	r.Lock()
	defer r.Unlock()

	_, err := r.client.HGet(ctx, "book", book.Id).Result()
	if err == redis.Nil {
		return store.ErrorNotFound
	} else if err != nil {
		return err
	}

	data, err := json.Marshal(book)
	if err != nil {
		return err
	}

	//可变接口的坑，多传参数也不会报错，只有运行时会报错
	sErr := r.client.HSet(ctx, "book", book.Id, string(data))
	if sErr != nil {
		return sErr.Err()
	}
	return nil
}

func (r *RedisStore) Get(id string) (store.Book, error) {
	result, err := r.client.HGet(ctx, "book", id).Result()
	if err == redis.Nil {
		return store.Book{}, store.ErrorNotFound
	} else if err != nil {
		return store.Book{}, err
	}
	var book store.Book
	serr := json.Unmarshal([]byte(result), &book)
	if serr != nil {
		return store.Book{}, serr
	}
	return book, nil
}

func (r *RedisStore) GetAll() ([]store.Book, error) {
	resultMap, err := r.client.HGetAll(ctx, "book").Result()
	if err != nil {
		return []store.Book{}, err
	}
	if len(resultMap) == 0 {
		return []store.Book{}, store.ErrorNotFound
	}

	var books []store.Book
	for _, v := range resultMap {
		var book store.Book
		json.Unmarshal([]byte(v), &book)
		books = append(books, book)
	}
	return books, nil
}

func (r *RedisStore) Delete(id string) error {
	_, err := r.client.HDel(ctx, "book", id).Result()
	if err != nil {
		return err
	}
	return nil
}
