package db

import (
	"context"
	"log"
)

func (h *Handler) FetchChildrenIds(ctx context.Context, parentId int32) ([]int32, error) {
	query := " SELECT * FROM " + TRANSACTION_TABLE + " WHERE ParentId in (?) "
	values := []interface{}{parentId}
	rows, err := h.readerWriter.Read(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	txnIds := make([]int32, 0, 3)
	for rows.Next() {
		txn := &TransactionRow{}
		err := rows.StructScan(txn)
		if err != nil {
			log.Println("[ERROR]:db:struct unmarshalling error:", parentId)
			continue
		}
		txnIds = append(txnIds, txn.Id)
	}
	return txnIds, nil
}
