package database

import (
	"encoding/json"
	"os"
)

type Client struct {
	path string
}

const mode = os.FileMode(0644)

func NewClient(path string) Client {
	return Client{path}
}

func (c Client) createDB() error {
	data, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})

	if err != nil {
		return err
	}

	return os.WriteFile(c.path, data, mode)
}

func (c Client) EnsureDB() error {
	_, err := os.Stat(c.path)
	if os.IsNotExist(err) {
		return c.createDB()
	}
	return err
}

func (c Client) updateDB(db databaseSchema) error {
	bytes, err := json.Marshal(db)
	if err != nil {
		return err
	}
	return os.WriteFile(c.path, bytes, mode)
}

func (c Client) readDB() (databaseSchema, error) {
	bytes, err := os.ReadFile(c.path)
	db := databaseSchema{}

	if err != nil {
		return db, nil
	}
	err = json.Unmarshal(bytes, &db)
	return db, err
}
