package structs

import (
	"fmt"
	"unsafe"
)

type User struct {
	ID         int
	Name       string
	Active     bool
	LoginCount int
	Tags       []string
}

func UpdateUserName(users map[int]User, id int, name string) bool {
	// TODO: изменить имя существующего пользователя, хранящегося в map как значение.
	// Если пользователя нет, вернуть false и не добавлять новый ключ.

	user, ok := users[id]
	if !ok {
		return false
	}
	user.Name = name
	users[id] = user
	return true
}

func ActivateUser(users map[int]User, id int) bool {
	// TODO: сделать существующего пользователя активным.
	// Если пользователя нет, вернуть false.

	user, ok := users[id]
	if !ok {
		return false
	}
	user.Active = true
	users[id] = user
	return true
}

func IncrementLoginCount(users map[int]User, id int) bool {
	// TODO: увеличить LoginCount существующего пользователя на единицу.
	// Если пользователя нет, вернуть false.
	user, ok := users[id]
	if !ok {
		return false
	}
	user.LoginCount += 1
	users[id] = user
	return true
}

func CopyUsers(users map[int]User) map[int]User {
	// TODO: вернуть независимую копию map и вложенных слайсов Tags.
	// Изменения любой части результата не должны менять исходные данные.

	copyUsers := make(map[int]User)

	if users == nil {
		return copyUsers
	}
	if len(users) > 0 {
		copyUsers = make(map[int]User, len(users))
	}

	for key, value := range users {
		var copyTags []string
		if value.Tags != nil {
			copyTags = make([]string, len(value.Tags))
			copy(copyTags, value.Tags)
		}
		newUser := value
		newUser.Tags = copyTags

		copyUsers[key] = newUser
	}
	return copyUsers
}

func BuildPointerIndex(users []User) map[int]*User {
	// TODO: построить индекс ID -> *User из независимых копий входных пользователей.
	// Указатели разных ID не должны вести к одному объекту или менять исходный слайс.

	newUsers := make(map[int]*User)

	if users == nil {
		return newUsers
	}
	if len(users) > 0 {
		newUsers = make(map[int]*User, len(users))
	}

	for _, value := range users {
		var copyTags []string
		if value.Tags != nil {
			copyTags = make([]string, len(value.Tags))
			copy(copyTags, value.Tags)
		}
		newUser := value
		newUser.Tags = copyTags

		newUsers[value.ID] = &newUser
	}
	return newUsers
}

func RenameThroughPointer(users map[int]*User, id int, name string) bool {
	// TODO: изменить имя объекта по существующему ненулевому указателю.
	// Для отсутствующего ключа или nil-значения вернуть false.

	if users == nil {
		return false
	}

	user, ok := users[id]
	if !ok || user == nil {
		return false
	}

	user.Name = name
	return true
}

func (u User) DisplayName() string {
	// TODO: вернуть строку "ID:Name".
	return fmt.Sprintf("%d:%s", u.ID, u.Name)
}

type Admin struct {
	User
	Permissions []string
}

func (a Admin) DisplayName() string {
	// TODO: вернуть строку "admin:ID:Name", перекрывая одноимённый метод встроенного User.
	return fmt.Sprintf("admin:%d:%s", a.ID, a.Name)
}

func (a Admin) HasPermission(permission string) bool {
	// TODO: проверить точное наличие permission в списке администратора.
	// Регистр, пробелы и повторения не нормализуются.

	for _, value := range a.Permissions {
		if value == permission {
			return true
		}
	}
	return false
}

func RenameEmbeddedUser(admin *Admin, name string) {
	// TODO: изменить Name у встроенного User.
	// Для nil admin функция должна безопасно завершиться.
	if admin == nil {
		return
	}
	admin.Name = name
}

type Account struct {
	ID      int
	Balance int
}

type PremiumAccount struct {
	Account
	Bonus int
}

func (p PremiumAccount) Total() int {
	// TODO: вернуть сумму основного баланса и бонуса.
	return p.Bonus + p.Balance
}

func PremiumLabel(p PremiumAccount) string {
	// TODO: вернуть строку "ID:Balance total=Total".
	return fmt.Sprintf("%d:%d total=%v", p.ID, p.Balance, p.Total())
}

type BadLayout struct {
	Active   bool
	Amount   int64
	Verified bool
	Retries  int32
}

type GoodLayout struct {
	Amount   int64
	Retries  int32
	Active   bool
	Verified bool
}

func UserValueSize(user User) uintptr {
	// TODO: вернуть фактический размер переданного значения User в байтах.
	return unsafe.Sizeof(user)
}

func UserPointerSize(user *User) uintptr {
	// TODO: вернуть фактический размер переменной-указателя на User в байтах.
	// Результат не зависит от того, nil указатель или нет.
	return unsafe.Sizeof(user)
}

func LayoutSizes() (bad, good, saved uintptr) {
	// TODO: вернуть размеры BadLayout и GoodLayout, затем число сэкономленных байт.
	bad = unsafe.Sizeof(BadLayout{})
	good = unsafe.Sizeof(GoodLayout{})
	saved = bad - good

	return bad, good, saved
}

var _ uintptr = unsafe.Sizeof(User{})
