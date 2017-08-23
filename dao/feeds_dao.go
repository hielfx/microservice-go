package dao

import (
	"database/sql"
	"fmt"
	"time"

	e "github.com/jhidalgo3/microservice-go/handleErr"
	"github.com/pkg/errors"
)

type Feed struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type FeedItem struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	URL         string         `json:"url"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	FeedID      int            `json:"feed_id" db:"feed_id"`
	Feed        `db:"feed"`    // << NOTE PREFIX
}

func CreateFeeds(f Feed) (int, error) {

	//result, err := DBAccess.db.Exec("INSERT into feeds (title,url,created_at) VALUES (?,?,?)", f.Title, f.URL, time.Now())
	result, err := DBAccess.db.Exec("INSERT into feeds (title,url) VALUES (?,?)", f.Title, f.URL)

	if err != nil {
		fmt.Printf("%v", err)
		return -1, errors.Wrap(err, e.DbQueryFail)
	}
	lastID, _ := result.LastInsertId()

	return int(lastID), nil
}

func DeleteFeeds(id int) error {
	_, err := DBAccess.db.Exec("DELETE FROM feeds WHERE feeds.id = ?", id)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.Wrap(err, e.DbQueryFail)
	}

	return nil
}

func GetFeeds() ([]Feed, error) {
	feeds := []Feed{}
	err := DBAccess.db.Select(&feeds, "SELECT * FROM feeds")
	return feeds, err
}

func CreateFeedItems(item FeedItem) (int, error) {
	result, err := DBAccess.db.Exec(
		`INSERT INTO bank_db.feed_items
			(title, url, description, feed_id)
			VALUES(?, ?, ?, ?)`, item.Title, item.URL, item.Description, item.FeedID)

	if err != nil {
		fmt.Printf("%v", err)
		return -1, errors.Wrap(err, e.DbQueryFail)
	}
	lastID, _ := result.LastInsertId()

	return int(lastID), nil
}

func DeleteFeedItems(id int) error {
	_, err := DBAccess.db.Exec("DELETE FROM feed_items WHERE feed_items.id = ?", id)
	if err != nil {
		fmt.Printf("%v", err)
		return errors.Wrap(err, e.DbQueryFail)
	}

	return nil
}

func GetFeedItems() ([]FeedItem, error) {
	items := []FeedItem{}
	sql := `SELECT
      feed_items.*,
      feeds.id "feed.id",
      feeds.title "feed.title",
      feeds.url "feed.url",
      feeds.created_at "feed.created_at"
    FROM
      feed_items JOIN feeds ON feed_items.feed_id = feeds.id;`
	err := DBAccess.db.Select(&items, sql)
	return items, err
}
