package main

import (
	"log"

	"github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
)

// Command-line args.
type args struct {
	Key      string `arg:"env"`
	Filepath string `arg:"positional"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var a args
	arg.MustParse(&a)

	p := NewPublisher(a.Key)
	article, err := parseFile(a.Filepath)
	if err != nil {
		log.Fatal(err)
	}

	err = p.Publish(article)
	if err != nil {
		log.Fatal(err)
	}
}
