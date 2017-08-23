package dao

import (
	"fmt"
	"log"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	setupConf()
	db, err := NewBankAPI(viper.GetString("database.URL"))
	if err != nil {
		log.Fatal(fmt.Errorf("FATAL: %+v\n", err))
	}
	DBAccess = db
	DeleteAllBanks()
}

func TestGetFeeds(t *testing.T) {
	id, _ := CreateFeeds(Feed{Title: "FEED_TITLE", URL: "FEED_URL"})
	feeds, err := GetFeeds()

	logFatalOnTest(t, err)
	assert.Len(t, feeds, 1, "Expected size is 1")

	DeleteFeeds(id)
}

func TestGetFeedItems(t *testing.T) {

	id, _ := CreateFeeds(Feed{Title: "FEED_TITLE", URL: "FEED_URL"})
	idItem, _ := CreateFeedItems(FeedItem{Title: "TITLE", URL: "URL", FeedID: id})

	feeds_items, err := GetFeedItems()
	logFatalOnTest(t, err)
	assert.Len(t, feeds_items, 1, "Expected size is 1")

	DeleteFeedItems(idItem)
	DeleteFeeds(id)
}
