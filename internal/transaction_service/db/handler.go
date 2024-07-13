package db

import "github.com/chkda/transaction_service/pkg/datastores/database"

type Handler struct {
	readerWriter database.ReaderWriter
}

func New(readerWriter database.ReaderWriter) *Handler {
	return &Handler{
		readerWriter: readerWriter,
	}
}
