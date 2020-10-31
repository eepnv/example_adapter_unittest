package layers

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// a simple implementation of get/set using redis
var ctx = context.Background()

func getRedisDatabase() DatabaseService {
	return RedisDatabaseImpl{
		rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "CvyZzBU8wa",
			DB:       0,
		}),
	}
}

type RedisDatabaseImpl struct {
	rdb *redis.Client
}

func (r RedisDatabaseImpl) ReadNotepad(id string) (*Notepad, error) {
	val, err := r.rdb.Get(ctx, id).Result()
	if err != nil {
		return nil, err
	}
	notepad := Notepad{
		ID:      id,
		Content: val,
	}
	return &notepad, nil
}
func (r RedisDatabaseImpl) UpdateNotepad(notepad Notepad) error {
	err := r.rdb.Set(ctx, notepad.ID, notepad.Content, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// note: you can use any database as long as what your provide to the function is consistent
// e.g. not only redis, you can also use memcache / mongodb / even text file or mySQL!

// an even more simple implementation of get/set using in-memory map
type SimpleDatabaseImpl struct {
	db map[string]string
}

func (f SimpleDatabaseImpl) ReadNotepad(id string) (*Notepad, error) {
	notepad := Notepad{
		ID:      id,
		Content: f.db[id],
	}
	return &notepad, nil
}
func (f SimpleDatabaseImpl) UpdateNotepad(notepad Notepad) error {
	f.db[notepad.ID] = notepad.Content
	return nil
}
