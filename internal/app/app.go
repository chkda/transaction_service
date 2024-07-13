package app

import (
	"github.com/chkda/transaction_service/internal/db"
	"github.com/chkda/transaction_service/pkg/datastores/cache"
	"github.com/chkda/transaction_service/pkg/datastores/database"
)

type Handler struct {
	inMemoryCache cache.ReaderWriter
	kvCache       cache.ReaderWriter
	dbHandler     *db.Handler
}

func New(
	inMemoryCache cache.ReaderWriter,
	kvCache cache.ReaderWriter,
	dbReaderWriter database.ReaderWriter,
) *Handler {
	dbHandler := db.New(dbReaderWriter)
	return &Handler{
		kvCache:       kvCache,
		dbHandler:     dbHandler,
		inMemoryCache: inMemoryCache,
	}
}
