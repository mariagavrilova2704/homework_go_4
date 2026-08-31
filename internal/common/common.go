package common

import (
	"strconv"

	"github.com/rinatkh/homework_go_4/internal/methods"
)

type Operation struct {
	AccountID int
	Kind      string
	Amount    int
}

type LedgerResult struct {
	Balances map[int]int
	Errors   []string
}

func ApplyOperations(accounts map[int]*methods.Account, operations []Operation) LedgerResult {
	// TODO: последовательно применить deposit/withdraw к счетам и продолжать после ошибок.
	// В Balances вернуть итоговые балансы ненулевых счетов, в Errors — сообщения в порядке операций.

	result := LedgerResult{
		Balances: make(map[int]int),
		Errors:   []string{},
	}

	if accounts == nil {
		for _, operation := range operations {
			result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": not found")
		}
		return result
	}

	for _, operation := range operations {
		account, ok := accounts[operation.AccountID]

		if !ok || account == nil {
			result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": not found")
			continue
		}

		var err error
		switch operation.Kind {

		case "deposit":
			err = account.Deposit(operation.Amount)
			if err != nil {
				switch err {
				case methods.ErrNilAccount:
					result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": not found")
				case methods.ErrInactive:
					result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": inactive account")
				case methods.ErrInvalidAmount:
					result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": invalid amount")
				default:
					result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": "+err.Error())
				}
			}

		case "withdraw":
			err = account.Withdraw(operation.Amount)
			if err != nil {
				switch err {
				case methods.ErrNilAccount:
					result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": not found")
				case methods.ErrInactive:
					result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": inactive account")
				case methods.ErrInvalidAmount:
					result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": invalid amount")
				case methods.ErrInsufficientFunds:
					result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": insufficient funds")
				default:
					result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": "+err.Error())
				}
			}

		default:
			result.Errors = append(result.Errors, "account "+strconv.Itoa(operation.AccountID)+": unknown operation \""+operation.Kind+"\"")
		}

	}
	for id, account := range accounts {
		if account != nil && account.Balance != 0 {
			result.Balances[id] = account.Balance
		}
	}

	return result
}

func BuildAccountIndex(accounts []methods.Account) map[int]*methods.Account {
	// TODO: построить индекс по ID из независимых копий счетов.
	// Разные ключи не должны вести к одной переменной, изменения индекса не должны менять входной слайс.
	finalResult := make(map[int]*methods.Account)
	if accounts == nil {
		return finalResult
	}
	clone := make([]methods.Account, len(accounts))
	copy(clone, accounts)

	for i := range clone {
		finalResult[clone[i].ID] = &clone[i]

	}
	return finalResult
}

type Event struct {
	UserID int
	Name   string
	Active *bool
}

func ApplyUserEvents(
	users map[int]string,
	active map[int]bool,
	events []Event,
) (map[int]string, map[int]bool) {
	// TODO: применить события к состоянию пользователей и вернуть обе map.
	// Непустое Name обновляет имя, ненулевой Active обновляет статус; nil map должны поддерживаться.

	if events == nil || len(events) == 0 {
		// Возможно, нужно инициализировать map даже если событий нет?
		if users == nil {
			users = make(map[int]string)
		}
		if active == nil {
			active = make(map[int]bool)
		}
		return users, active
	}

	if users == nil {
		users = make(map[int]string)
	}
	if active == nil {
		active = make(map[int]bool)
	}

	for _, value := range events {
		if value.Name != "" {
			users[value.UserID] = value.Name
		}
		if value.Active != nil {
			active[value.UserID] = *value.Active
		}
	}

	return users, active
}

//коммент для проверки
