package db

import (
	"context"
	"fmt"
	"strings"
)

type TransactionRow struct {
	Id       int32   `db:"id"`
	Type     string  `db:"type"`
	Amount   float64 `db:"amount"`
	ParentId *int32  `db:"parent_id"`
}

const (
	TRANSACTION_TABLE = "transactions"
)

var columns = []string{
	"Id",
	"Amount",
	"Type",
	"ParentId",
}

func (h *Handler) InsertTransaction(ctx context.Context, txn *TransactionRow) error {
	query, values := h.insertTransactionQueryBuilder(txn)
	err := h.readerWriter.Write(ctx, query, values...)
	return err
}

func (h *Handler) insertTransactionQueryBuilder(txn *TransactionRow) (string, []interface{}) {
	query := "INSERT INTO " + TRANSACTION_TABLE + "(" + strings.Join(columns, ",") + ") " + " VALUES (?,?,?,?) "
	values := []interface{}{
		txn.Id,
		txn.Amount,
		txn.Type,
		txn.ParentId,
	}
	return query, values
}

func (h *Handler) FetchTransaction(ctx context.Context, id int32) (*TransactionRow, error) {
	query := "SELECT * FROM " + TRANSACTION_TABLE + " WHERE Id in (?) "
	args := []interface{}{id}
	rows, err := h.readerWriter.Read(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	txn := &TransactionRow{}
	for rows.Next() {
		err := rows.StructScan(txn)
		if err != nil {
			return nil, err
		}
		break
	}
	return txn, nil
}

func (h *Handler) FetchTransactionIds(ctx context.Context, txnType string) ([]int32, error) {
	query := "SELECT * FROM " + TRANSACTION_TABLE + " WHERE Type IN (?) "
	args := []interface{}{txnType}
	rows, err := h.readerWriter.Read(ctx, query, args)
	if err != nil {
		return nil, err
	}
	txnIds := make([]int32, 0, 3)
	for rows.Next() {
		txn := &TransactionRow{}
		err := rows.StructScan(txn)
		if err != nil {
			// TODO : Add logger instead of print
			fmt.Println(err)
			continue
		}
		txnIds = append(txnIds, txn.Id)
	}
	return txnIds, nil
}
