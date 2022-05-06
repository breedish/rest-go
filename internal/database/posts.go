package database

import (
	"github.com/google/uuid"
	"time"
)

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	_, err := c.GetUser(userEmail)
	if err != nil {
		return Post{}, err
	}

	post := Post{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UserEmail: userEmail,
		Text:      text,
	}

	db, err := c.readDB()
	db.Posts[post.ID] = post

	err = c.updateDB(db)
	return post, err
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	posts := []Post{}

	_, err := c.GetUser(userEmail)
	if err != nil {
		return posts, err
	}

	db, err := c.readDB()
	if err != nil {
		return posts, err
	}
	for _, p := range db.Posts {
		if p.UserEmail == userEmail {
			posts = append(posts, p)
		}
	}

	return posts, nil
}

func (c Client) DeletePost(id string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	delete(db.Posts, id)

	return c.updateDB(db)
}
