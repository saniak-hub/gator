package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	xmlData := RSSFeed{}
	if err := xml.Unmarshal(data, &xmlData); err != nil {
		return &RSSFeed{}, err
	}
	fmt.Println(xmlData)
	return &xmlData, nil
}

func (r *RSSFeed) parser() {
	r.Channel.Title = html.UnescapeString(r.Channel.Title)
	r.Channel.Link = html.UnescapeString(r.Channel.Link)
	r.Channel.Description = html.UnescapeString(r.Channel.Description)
	for i := range r.Channel.Item {
		r.Channel.Item[i].Title = html.UnescapeString(r.Channel.Item[i].Title)
		r.Channel.Item[i].Link = html.UnescapeString(r.Channel.Item[i].Link)
		r.Channel.Item[i].Description = html.UnescapeString(r.Channel.Item[i].Description)
		r.Channel.Item[i].PubDate = html.UnescapeString(r.Channel.Item[i].PubDate)
	}
}
