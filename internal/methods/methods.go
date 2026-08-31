package methods

import (
	"errors"
	"fmt"
)

var (
	ErrNilAccount        = errors.New("nil account")
	ErrNilDestination    = errors.New("nil destination")
	ErrSameAccount       = errors.New("source and destination are the same")
	ErrInactive          = errors.New("inactive account")
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNilCart           = errors.New("nil cart")
	ErrInvalidSKU        = errors.New("invalid sku")
	ErrCartLimit         = errors.New("cart limit exceeded")
)

type Account struct {
	ID       int
	Owner    string
	Balance  int
	Active   bool
	Currency string
}

func (a Account) Label() string {
	// TODO: вернуть строку "#ID Owner: Balance Currency" с одним пробелом после двоеточия.
	return fmt.Sprintf("#%d %s: %d %s", a.ID, a.Owner, a.Balance, a.Currency)
}

func (a Account) CanWithdraw(amount int) bool {
	// TODO: сообщить, можно ли списать amount без нарушения состояния счёта.
	// Счёт должен быть активным, сумма — положительной, денег должно хватать.
	if a.Active == true && amount > 0 && a.Balance >= amount {
		return true
	}
	return false
}

func (a Account) Renamed(owner string) Account {
	// TODO: вернуть копию счёта с новым владельцем.
	// Исходное значение Account не должно изменяться.
	a.Owner = owner
	return a
}

func (a *Account) Rename(owner string) error {
	// TODO: изменить владельца исходного счёта.
	// Для nil receiver вернуть ErrNilAccount и не паниковать.
	if a == nil {
		return ErrNilAccount
	}

	a.Owner = owner
	return nil
}

func (a *Account) Deposit(amount int) error {
	// TODO: пополнить активный счёт на положительную сумму.
	// При ошибке вернуть подходящую sentinel-ошибку и сохранить прежний баланс.

	if a == nil {
		return ErrNilAccount
	}
	if a.Active == false {
		return ErrInactive
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	a.Balance = a.Balance + amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	// TODO: списать положительную сумму с активного счёта, если денег достаточно.
	// При ошибке вернуть подходящую sentinel-ошибку и сохранить прежний баланс.

	if a == nil {
		return ErrNilAccount
	}
	if a.Active == false {
		return ErrInactive
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	a.Balance = a.Balance - amount
	return nil
}

func (a *Account) TransferTo(destination *Account, amount int) error {
	// TODO: атомарно перевести сумму между двумя разными активными счетами.
	// При любой ошибке оба баланса должны остаться прежними.

	if a == nil {
		return ErrNilAccount
	}
	if destination == nil {
		return ErrNilDestination
	}
	if a == destination {
		return ErrSameAccount
	}
	if a.Active == false || destination.Active == false {
		return ErrInactive
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if a.Balance < amount {
		return ErrInsufficientFunds
	}
	a.Balance = a.Balance - amount
	destination.Balance = destination.Balance + amount
	return nil
}

func (a *Account) Activate() error {
	// TODO: сделать счёт активным.
	// Для nil receiver вернуть ErrNilAccount.

	if a == nil {
		return ErrNilAccount
	}

	if a.Active == false {
		a.Active = true
	}
	return nil
}

func (a *Account) Deactivate() error {
	// TODO: сделать счёт неактивным.
	// Для nil receiver вернуть ErrNilAccount.
	if a == nil {
		return ErrNilAccount
	}

	if a.Active {
		a.Active = false
	}
	return nil
}

func (a Account) Snapshot() Account {
	// TODO: вернуть независимую копию текущего значения счёта.
	copyAccount := a
	return copyAccount
}

func (a *Account) Reset() {
	// TODO: сбросить все поля существующего счёта в zero value.
	// Вызов для nil receiver должен быть безопасным.

	if a == nil {
		return
	}
	*a = Account{} //автоматическое разыменование происходит только когда мы "стучимся" к определенным полям структуры, если же мы хотим поменять все поля (то есть всю структуру), то используем стандартное разыменование
}

func (a *Account) ApplyBonus(percent int) error {
	// TODO: увеличить баланс активного счёта на указанный неотрицательный процент.
	// Дробная часть от целочисленного расчёта отбрасывается; при ошибке баланс не менять.

	if a == nil {
		return ErrNilAccount
	}
	if a.Active == false {
		return ErrInactive
	}
	if percent < 0 {
		return ErrInvalidAmount
	}
	a.Balance = a.Balance + a.Balance*percent/100
	return nil
}

func (a Account) SameOwner(other Account) bool {
	// TODO: сравнить владельцев двух счетов точным сравнением строк.
	if a.Owner == other.Owner {
		return true
	}
	return false
}

type Cart struct {
	Items    map[string]int
	MaxItems int
}

func (c *Cart) Add(sku string, count int) error {
	// TODO: добавить положительное количество товара с непустым SKU.
	// Общее число единиц после добавления не должно превышать MaxItems.
	if c == nil {
		return ErrNilCart
	}
	if sku == "" {
		return ErrInvalidSKU
	}
	if count <= 0 {
		return ErrInvalidAmount
	}

	sum := 0
	for _, value := range c.Items {
		sum = sum + value
	}
	if sum+count > c.MaxItems {
		return ErrCartLimit
	}
	// ТОЛЬКО ПОСЛЕ проверок всех ошибок мы проводим инициализацию мапы
	if c.Items == nil {
		c.Items = make(map[string]int)
	}

	c.Items[sku] = c.Items[sku] + count
	return nil
}

func (c Cart) TotalItems() int {
	// TODO: вернуть сумму количеств всех товаров корзины.

	sum := 0
	for _, value := range c.Items {
		sum = sum + value
	}
	return sum
}
