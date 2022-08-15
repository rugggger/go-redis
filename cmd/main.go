package main

import (
	"fmt"

	transactionlog "github.com/rugggger/go-redis/src/transaction-log"
)

var logger transactionlog.TransactionLogger

func main() {

	fmt.Println("test")
	Put("test", "yaron")
	fmt.Println(Get("test"))
	Delete("test")
	fmt.Println(Get("test"))
	initializeTransactionLog()
	serve()

}

func initializeTransactionLog() error {
	var err error
	logger, err := transactionlog.NewFileTransactionLogger("transaction.log")
	if err != nil {
		return fmt.Errorf("failed to create event logger")
	}
	events, errors := logger.ReadEvents()
	e, ok := transactionlog.Event{}, true
	for ok && err == nil {
		select {
		case err, ok = <-errors:

		case e = <-events:
			switch e.EventType {
			case transactionlog.EventDelete:
				err = Delete(e.Key)
			case transactionlog.EventPut:
				err = Put(e.Key, e.Value)
			}

		}
	}

	logger.Run()
	return err
}
