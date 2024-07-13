package db

import "context"

type SumResult struct {
	Id  int32   `db:"Id"`
	Sum float64 `db:"Sum"`
}

func (h *Handler) FetchSumForTransactionId(ctx context.Context, id int32) (*SumResult, error) {
	query := "SELECT Id, SUM(Amount) AS SUM FROM " + TRANSACTION_TABLE + " WHERE Id IN (?) "
	args := []interface{}{id}
	rows, err := h.readerWriter.Read(ctx, query, args)
	if err != nil {
		return nil, err
	}
	res := &SumResult{}
	for rows.Next() {
		err := rows.StructScan(res)
		if err != nil {
			return nil, err
		}
		break
	}
	return res, nil
}
