package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Article represents an article.
type Article struct {
	Title     string   `json:"title"`
	Text      string   `json:"text"`
	Tags      []string `json:"tags"`
	Published bool     `json:"published"`
	Series    string   `json:"series"`
	User      string   `json:"user"`
}

func parseText(text string) (Article, error) {
	pattern := regexp.MustCompile(`#\s+(.+)\s+##\sMeta\s+tags:\s+(.+)\s+(?s)(.+)`)
	matches := pattern.FindStringSubmatch(text)
	if len(matches) != 4 {
		return Article{}, errors.New("unexpected file format: " + text)
	}

	return Article{
		Title:  matches[1],
		Text:   matches[3],
		Tags:   strings.Split(matches[2], ", "),
		Series: "",
	}, nil
}

func parseFile(filepath string) (Article, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return Article{}, fmt.Errorf("reading article file: %v", err)
	}
	text := string(bytes)
	return parseText(text)
}
