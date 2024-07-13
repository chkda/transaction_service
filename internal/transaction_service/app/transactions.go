package app

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/chkda/transaction_service/internal/transaction_service/db"
	"github.com/chkda/transaction_service/internal/transaction_service/interfaces/transactions"
	"github.com/chkda/transaction_service/pkg/datastores/cache"
)

func (h *Handler) AddTransaction(
	ctx context.Context,
	txnId int32,
	txnDetails *transactions.Transaction,
) error {
	txn := &db.TransactionRow{
		Id:       txnId,
		Amount:   txnDetails.Amount,
		Type:     txnDetails.Type,
		ParentId: txnDetails.ParentId,
	}
	err := h.dbHandler.InsertTransaction(ctx, txn)
	return err
}

func (h *Handler) GetTransaction(ctx context.Context, txnId int32) (*transactions.Transaction, error) {
	txnBytes, err := h.inMemoryCache.Read(ctx, strconv.Itoa(int(txnId)))
	if err == nil {
		txn, err := unmarshalToTransaction(txnBytes)
		if err == nil {
			return txn, nil
		}
	}

	// TODO: Add logger for cache miss

	txnBytes, err = h.kvCache.Read(ctx, strconv.Itoa(int(txnId)))
	if err == nil {
		txn, err := unmarshalToTransaction(txnBytes)
		if err == nil {
			return txn, nil
		}
	}

	// TODO: Add logger for cache miss
	txnRow, err := h.dbHandler.FetchTransaction(ctx, txnId)
	if err != nil {
		return nil, err
	}
	txn := &transactions.Transaction{
		Amount:   txnRow.Amount,
		Type:     txnRow.Type,
		ParentId: txnRow.ParentId,
	}
	txnBytes, err = json.Marshal(txn)
	if err != nil {
		return nil, err
	}

	// Setting in memory to handle sudden traffic burst
	inMemoryCachePayload := &cache.Payload{
		Key:   strconv.Itoa(int(txnId)),
		Value: txnBytes,
		TTL:   time.Duration(time.Second * 300),
	}
	go h.inMemoryCache.Write(ctx, inMemoryCachePayload)

	// Setting to distributed kv cache
	kvCachePayload := &cache.Payload{
		Key:   strconv.Itoa(int(txnId)),
		Value: txnBytes,
		TTL:   time.Duration(time.Second * 900),
	}
	go h.kvCache.Write(ctx, kvCachePayload)

	return txn, nil
}

func unmarshalToTransaction(data []byte) (*transactions.Transaction, error) {
	txn := &transactions.Transaction{}
	err := json.Unmarshal(data, txn)
	if err != nil {
		return nil, err
	}
	return txn, nil
}
