package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ponzaa555/rssagg/internal/database"
)

// concurrency mean how many goroutine will be run at the same time
func startScaping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	// make request
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			fmt.Println("Error fetching feeds", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			// everytime  spawn goroutine
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait() // wait until all goroutine finish
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	//Mark fetchd feed
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched %v", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		// if item.Description is null will set Null value to database
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Cloudn't parse date %v with err %v", item.PubDate, err)
			continue
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreateAt:    time.Now().UTC(),
			UpdateAt:    time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubAt,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		// log.Println("Found Post", item.Title, "on feed", feed.Name)
		if err != nil {
			// fix problem run app again and it detect blog that already fetch
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			// will log in case not duplicate key
			log.Println("failed to create post:", err)
		}
	}
	log.Printf("Feed %s collected , %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
