package services

import (
	"BankApp/models"
	"errors"
)

func Login(username, password string) (*models.Account, error) {
	accounts, err := LoadAccounts()
	if err != nil {
		return nil, err
	}

	for i := range accounts {
		if accounts[i].Username == username &&
			accounts[i].Password == password {
			return &accounts[i], nil
		}
	}

	return nil, errors.New("Invalid username or password.")
}

func ChangePassword(acc *models.Account, newPassword string) error {
	if acc == nil {
		return errors.New("Account dont exist.")
	}
	if newPassword == "" {
		return errors.New("New password must not be empty.")
	}

	updated := *acc
	updated.Password = newPassword

	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Password = updated.Password
	return nil
}
