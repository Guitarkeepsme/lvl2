package main

import "fmt"

type Ledger struct {
}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Создаю запись в базе для учётной записи под номером %s с транзакцией типа %s в количестве %d\n", accountID, txnType, amount)
}
