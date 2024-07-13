package app

import (
	"context"
	"encoding/json"
	"time"

	"github.com/chkda/transaction_service/pkg/datastores/cache"
)

const (
	typesKeyPrefix = "t_"
)

type CacheableTxnIds struct {
	Ids []int32 `json:"txn_ids"`
}

func (h *Handler) GetTransactionsWithSameType(ctx context.Context, txnType string) ([]int32, error) {
	txnBytes, err := h.kvCache.Read(ctx, typesKeyPrefix+txnType)
	if err == nil {
		txnIds, err := unmarshalToCacheableTxnIds(txnBytes)
		if err == nil {
			return txnIds.Ids, nil
		}
	}

	txnIds, err := h.dbHandler.FetchTransactionIds(ctx, txnType)
	if err != nil {
		return nil, err
	}

	txnBytes, err = json.Marshal(&CacheableTxnIds{
		Ids: txnIds,
	})
	if err != nil {
		return nil, err
	}
	kvCachePayload := &cache.Payload{
		Key:   typesKeyPrefix + txnType,
		Value: txnBytes,
		TTL:   time.Duration(time.Second * 900),
	}
	go h.kvCache.Write(ctx, kvCachePayload)
	return txnIds, nil
}

func unmarshalToCacheableTxnIds(data []byte) (*CacheableTxnIds, error) {
	txnIds := &CacheableTxnIds{}
	err := json.Unmarshal(data, txnIds)
	if err != nil {
		return nil, err
	}
	return txnIds, nil
}
