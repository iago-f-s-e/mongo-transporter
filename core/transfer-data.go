package core

import (
	"context"
	"fmt"
	"mongo_transporter/domain"
	"os"
	"sync"
)

func transferData(ctx context.Context, batchSize int64, receiver *domain.Receiver, sender *domain.Sender, wg *sync.WaitGroup) {
	var lastPosition int64 = 0
	count := 1
	receiver.DeleteAllCollection(ctx)

	for {
		documents, newLastPosition, err := sender.GetCollectionWithPagination(ctx, batchSize, lastPosition)

		if err != nil {
			os.Exit(1)
		}

		if len(documents) == 0 {
			break
		}

		wg.Add(1)
		go func(documents []interface{}, batchNumber int) {
			defer wg.Done()
			fmt.Printf("Inserting batch number %d with size %d\n", batchNumber, len(documents))
			receiver.InsertOnCollection(ctx, documents)
			fmt.Printf("Successful insert batch number  %d with size %d\n", batchNumber, len(documents))
		}(documents, count)

		lastPosition = newLastPosition
		count++
	}
}
