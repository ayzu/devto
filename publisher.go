package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Publisher publishes article in dev.to.
type Publisher struct {
	key    string
	client *Client
}

// NewPublisher constructor.
func NewPublisher(key string) *Publisher {
	client, err := NewClient("https://dev.to/api")
	if err != nil {
		log.Fatal("init DevTo client: ", err)
	}
	return &Publisher{key: key, client: client}
}

// MyArticles lists all articles for the current user.
func (p *Publisher) MyArticles() ([]Article, error) {
	resp, err := p.client.GetUserAllArticles(
		context.Background(),
		&GetUserAllArticlesParams{},
		func(ctx context.Context, req *http.Request) error {
			req.Header.Add("api-key", p.key)
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("obtaining articles: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	var data []ArticleMe
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding server response: %v", err)
	}

	var out []Article
	for _, am := range data {
		out = append(out, articleMeToArticle(am))
	}

	return out, nil
}

func articleMeToArticle(am ArticleMe) Article {
	return Article{
		Title:     am.Title,
		Tags:      am.TagList,
		Published: am.Published,
		Text:      am.BodyMarkdown,
	}
}

// TODO: parse response.
type PublishedInfo struct {
	ID                 int
	URL                string
	ReadingTimeMinutes int
}

// Publish an article.
func (p *Publisher) Publish(a Article) error {
	body := CreateArticleJSONRequestBody{
		Article: &struct {
			BodyMarkdown   *string   `json:"body_markdown,omitempty"`
			CanonicalUrl   *string   `json:"canonical_url,omitempty"`
			Description    *string   `json:"description,omitempty"`
			MainImage      *string   `json:"main_image,omitempty"`
			OrganizationId *int32    `json:"organization_id,omitempty"`
			Published      *bool     `json:"published,omitempty"`
			Series         *string   `json:"series,omitempty"`
			Tags           *[]string `json:"tags,omitempty"`
			Title          *string   `json:"title,omitempty"`
		}{
			BodyMarkdown:   &a.Text,
			CanonicalUrl:   nil,
			Description:    nil,
			MainImage:      nil,
			OrganizationId: nil,
			Published:      &a.Published,
			Series:         nil,
			Tags:           nil,
			Title:          &a.Title,
		},
	}

	resp, err := p.client.CreateArticle(
		context.Background(),
		body,
		func(ctx context.Context, req *http.Request) error {
			req.Header.Add("api-key", p.key)
			return nil
		})
	if err != nil {
		return fmt.Errorf("posting article: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	return nil
}
