package liberr

import (
	"errors"
)

//goland:noinspection ALL
var (
	// ErrFailedAction indicates failed write databases
	ErrFailedAction = errors.New("Failed write to databases. No data affects")

	// ErrUnexpected indicates this data is not found
	ErrUnexpected = errors.New("Data not found")

	// ErrDataNotFound indicates this data is not found
	ErrDataNotFound = errors.New("Data not found")

	// ErrInvalidAccountNumber indicates invalid id
	ErrInvalidAccountNumber = errors.New("Invalid account number")

	// ErrInvalidAmountTransfer indicates invalid id
	ErrInvalidAmountTransfer = errors.New("Invalid amount transfer")

	// InvalidAccountTransfer indicates invalid id
	InvalidAccountTransfer = "Invalid %s_account_number: %s"

	// ErrDuplicateAccountNumber indicates invalid id
	DuplicateAccountNumber = "Duplicate account number: %s"

	// CannotBeBlank indicates this fields is required
	CannotBeBlank = "%s: cannot be blank."
)
