package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/saniak-hub/gator/internal/database"
)

func scrapeFeeds(s *state) {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.db.UpdateFeed(ctx, feed.ID); err != nil {
		log.Fatal(err)
	}

	xml, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range xml.Channel.Item {
		pubdate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Fatal(err)
		}

		post := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: pubdate, Valid: true},
			FeedID:      uuid.NullUUID{UUID: feed.ID, Valid: true},
		}

		_, err = s.db.CreatePost(ctx, post)
		if err != nil {
			fmt.Println(err)
		}
	}
}
