package services

import (
	"BankApp/models"
	"errors"
	"fmt"
)

const MaxDepositAmount = 10000

func UpdateAccount(updated models.Account) error {
	accounts, err := LoadAccounts()
	if err != nil {
		return err
	}

	for i := range accounts {
		if accounts[i].Username == updated.Username {
			accounts[i] = updated
			return SaveAccounts(accounts)
		}
	}

	return errors.New("Account does not exist.")
}

func Deposit(acc *models.Account, amount float64) error {
	if acc == nil {
		return errors.New("Account is nil.")
	}
	if amount <= 0 {
		return errors.New(
			"Deposit amount must be greater than zero",
		)
	}
	if amount > MaxDepositAmount {
		return fmt.Errorf(
			"deposit amount must not exceed %.2f per transaction",
			float64(MaxDepositAmount),
		)
	}

	updated := *acc
	updated.Balance += amount

	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Balance = updated.Balance
	return nil
}

func Withdraw(acc *models.Account, amount float64) error {
	if acc == nil {
		return errors.New("account is nil")
	}
	if amount <= 0 {
		return errors.New(
			"withdraw amount must be greater than zero",
		)
	}
	if amount > acc.Balance {
		return errors.New("account balance is not enough")
	}

	updated := *acc
	updated.Balance -= amount

	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Balance = updated.Balance
	return nil
}

func Transfer(from *models.Account, toUser string, amount float64) error {
	if from == nil {
		return errors.New("Sender account is NULL")
	}

	if amount <= 0 {
		return errors.New("Transfer amount must be higher than zero.")
	}

	if from.Username == toUser {
		return errors.New("Sender and reicever must be diference accounts")
	}

	accounts, err := LoadAccounts()
	if err != nil {
		return err
	}

	var sender *models.Account
	var receiver *models.Account

	for i := range accounts {
		if accounts[i].Username == from.Username {
			sender = &accounts[i]
		}

		if accounts[i].Username == toUser {
			receiver = &accounts[i]
		}
	}

	if sender == nil || receiver == nil {
		return errors.New("Sender or Receiver account is not existed!")
	}

	if sender.Balance < amount {
		return errors.New("Not enough cash, please deposit more!")
	}

	sender.Balance -= amount
	receiver.Balance += amount

	if err := SaveAccounts(accounts); err != nil {
		return err
	}

	from.Balance = sender.Balance

	return nil
}
