package app

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"sync"

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

	log.Println("[INFO]:app: kv cache miss:key:", txnId)

	sumResult, err := h.dbHandler.FetchSumForTransactionId(ctx, txnId)
	if err != nil {
		return -1, err
	}
	childrenIds, err := h.dbHandler.FetchChildrenIds(ctx, txnId)
	if err != nil {
		return -1, err
	}

	sum := sumResult.Sum
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for _, childId := range childrenIds {
		wg.Add(1)
		go func() {
			defer wg.Done()
			childSum, err := h.GetSumForTxnId(context.Background(), childId)
			if err != nil {
				log.Println("[ERROR]:app:" + err.Error())
				return
			}
			mu.Lock()
			defer mu.Unlock()
			sum += childSum
		}()
	}
	wg.Wait()
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
		TTL:   kvCacheTTL,
	}
	go h.kvCache.Write(context.Background(), kvCachePayload)
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
