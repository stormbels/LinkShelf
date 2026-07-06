package main

import (
	"context"
	"errors"

	"github.com/stormbels/linkshelf/internal/config"
	"github.com/stormbels/linkshelf/internal/storage"
	postgresstorage "github.com/stormbels/linkshelf/internal/storage/postgres"
)

func newLinkStorage(cfg *config.Config) (storage.LinkStorage, func(), error) {
	switch cfg.Storage.Type {
	case "memory":
		return storage.NewMemoryStorage(), func() {}, nil
	case "postgres":
		postgresStorage, err := postgresstorage.New(context.Background(), cfg.Storage.Postgres.DSN())
		if err != nil {
			return nil, nil, err
		}

		return postgresStorage, postgresStorage.Close, nil
	default:
		return nil, nil, errors.New("unknown storage type: " + cfg.Storage.Type)
	}
}
