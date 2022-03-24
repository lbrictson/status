package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/lbrictson/status/ent"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	client *ent.Client
}

type NewStoreConfig struct {
	FileLocation string
}

func MustNewStore(config NewStoreConfig) *Store {
	s := Store{}
	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%v?_fk=1", config.FileLocation))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	s.client = client
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &s
}
