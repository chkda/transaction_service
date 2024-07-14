package app

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/chkda/transaction_service/internal/db"
	"github.com/chkda/transaction_service/internal/interfaces/transactions"
	"github.com/chkda/transaction_service/pkg/datastores/cache"
)

const (
	inMemoryCacheTTL = time.Duration(time.Second * 5)
	kvCacheTTL       = time.Duration(time.Second * 10)
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

	log.Println("[INFO]:app: inmemory cache miss:key:", txnId)

	txnBytes, err = h.kvCache.Read(ctx, strconv.Itoa(int(txnId)))
	if err == nil {
		txn, err := unmarshalToTransaction(txnBytes)
		if err == nil {
			return txn, nil
		}
	}

	log.Println("[INFO]:app: kv cache miss:key:", txnId)

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
		TTL:   inMemoryCacheTTL,
	}
	go h.inMemoryCache.Write(context.Background(), inMemoryCachePayload)

	// Setting to distributed kv cache
	kvCachePayload := &cache.Payload{
		Key:   strconv.Itoa(int(txnId)),
		Value: txnBytes,
		TTL:   kvCacheTTL,
	}
	go h.kvCache.Write(context.Background(), kvCachePayload)

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
