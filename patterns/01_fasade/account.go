package main

import "fmt"

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("некорректное имя учётной записи")
	}
	fmt.Println("Учётная запись подтверждена")
	return nil
}
