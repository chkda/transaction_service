package app

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/chkda/transaction_service/pkg/datastores/cache"
)

const (
	sumKeyPrefix = "s_"
)

type CacheableSumForTxnId struct {
	TxnId int32   `json:"txn_id"`
	Sum   float64 `json:"sum"`
}

func (h *Handler) GetSumForTxnId(ctx context.Context, txnId int32) (float64, error) {
	cacheBytes, err := h.kvCache.Read(ctx, sumKeyPrefix+strconv.Itoa(int(txnId)))
	if err == nil {
		cacheableSumForTxnId, err := unmarshalToCacheableSumForTxnId(cacheBytes)
		if err == nil {
			return cacheableSumForTxnId.Sum, nil
		}
	}
	sumResult, err := h.dbHandler.FetchSumForTransactionId(ctx, txnId)
	if err != nil {
		return -1, err
	}
	childrenIds, err := h.dbHandler.FetchChildrenIds(ctx, txnId)
	if err != nil {
		return -1, err
	}
	if len(childrenIds) == 0 {
		return sumResult.Sum, nil
	}
	sum := sumResult.Sum
	for _, childId := range childrenIds {
		// TODO: Add go routines line 37-43
		childSum, err := h.GetSumForTxnId(ctx, childId)
		if err != nil {
			// TODO: Add logger
			fmt.Println(err)
			continue
		}
		sum += childSum
	}
	cacheableSumForTxnId := &CacheableSumForTxnId{
		TxnId: txnId,
		Sum:   sum,
	}
	cacheBytes, err = json.Marshal(cacheableSumForTxnId)
	if err != nil {
		return 0, err
	}

	// Setting to distributed kv cache
	kvCachePayload := &cache.Payload{
		Key:   sumKeyPrefix + strconv.Itoa(int(txnId)),
		Value: cacheBytes,
		TTL:   time.Duration(time.Second * 900),
	}
	go h.kvCache.Write(ctx, kvCachePayload)
	return sum, nil
}

func unmarshalToCacheableSumForTxnId(data []byte) (*CacheableSumForTxnId, error) {
	cacheableSumForTxnId := &CacheableSumForTxnId{}
	err := json.Unmarshal(data, cacheableSumForTxnId)
	if err != nil {
		return nil, err
	}
	return cacheableSumForTxnId, nil
}
