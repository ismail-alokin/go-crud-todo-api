package db

import (
	"context"
	"os"

	"github.com/edgedb/edgedb-go"
	"github.com/ismail-alokin/go-crud-api-todo/config/logger"
)

var DbClient *edgedb.Client

func init() {
	var err error

	logger.Println("Connecting to EdgeDB")
	ctx := context.Background()

	options := edgedb.Options{
		Database:    os.Getenv("EDGEDB_SERVER_DATABASE"),
		User:        os.Getenv("EDGEDB_SERVER_USER"),
		Password:    edgedb.NewOptionalStr(os.Getenv("EDGEDB_SERVER_PASSWORD")),
		Host:        os.Getenv("EDGEDB_HOST"),
		TLSSecurity: os.Getenv("EDGEDB_CLIENT_TLS_SECURITY"),
	}

	DbClient, err = edgedb.CreateClient(ctx, options)

	if err != nil {
		logger.Fatalf("Edgedb connection creation failed", err)
	}

	if DbClient != nil {
		logger.Println("EdgeDB client connected successfully")
	} else {
		logger.Println("EdgeDB client is nil, connection failed")
	}
}
