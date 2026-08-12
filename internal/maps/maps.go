package maps

import "slices"

type User struct {
	ID     int
	Name   string
	City   string
	Active bool
	Tags   []string
}

type Product struct {
	SKU      string
	Quantity int
	Price    int
}

func CountWords(words []string) map[string]int {
	// TODO: посчитать, сколько раз встречается каждое слово.
	// Вернуть новую готовую к записи map. Регистр и пустая строка имеют значение.

	newMap := make(map[string]int)
	for _, word := range words {
		newMap[word]++
	}
	return newMap
}

func GetOrDefault(values map[string]int, key string, fallback int) int {
	// TODO: вернуть значение существующего ключа или fallback, если ключ отсутствует.
	// Нулевое значение по существующему ключу должно возвращаться как обычное значение.
	value, ok := values[key]
	if !ok {
		return fallback
	}
	return value
}

func DeleteAndReport(values map[string]int, key string) bool {
	// TODO: удалить ключ и сообщить, существовал ли он до вызова.
	// Для отсутствующего ключа, пустой и nil map вернуть false без panic.
	_, ok := values[key]
	delete(values, key) //безопасная встроенная функция, не вызывает панику
	return ok
}

func MergeCounters(left, right map[string]int) map[string]int {
	// TODO: объединить два счётчика в новой map, складывая значения общих ключей.
	// Исходные map после вызова должны остаться без изменений.
	finalMap := make(map[string]int, len(left)+len(right))
	for k, v := range left {
		finalMap[k] = v
	}
	for k, v := range right {
		finalMap[k] = finalMap[k] + v
	}
	return finalMap
}

func KeysByValueAtLeast(values map[string]int, min int) []string {
	// TODO: вернуть ключи, значения которых не меньше min.
	// Результат должен иметь стабильный алфавитный порядок.
	keys := make([]string, 0, len(values))
	for k := range values {
		if values[k] >= min {
			keys = append(keys, k)
		}
	}
	slices.Sort(keys)
	return keys
}

func MaxKeyByValue(values map[string]int) (string, bool) {
	// TODO: найти ключ с максимальным значением.
	// Для пустого входа вернуть "", false; при равенстве выбрать меньший ключ по алфавиту.
	if len(values) == 0 {
		return "", false
	}
	maxValue := 0
	maxKey := ""
	hasValue := false

	for k, v := range values {
		if !hasValue || v > maxValue || (v == maxValue && k < maxKey) {
			maxValue = v
			maxKey = k
			hasValue = true
		}
	}
	return maxKey, true
}

func BuildUserIndex(users []User) map[int]User {
	// TODO: построить индекс пользователей по ID.
	// Если один ID встречается несколько раз, в результате должен остаться последний пользователь.
	if len(users) == 0 {
		return nil
	}
	idMap := make(map[int]User, len(users))
	for _, user := range users {
		idMap[user.ID] = user
	}
	return idMap
}

func GroupActiveUsersByCity(users []User) map[string][]User {
	// TODO: сгруппировать только активных пользователей по городу.
	// Порядок пользователей внутри каждого города должен совпадать с исходным слайсом.
	cityMap := make(map[string][]User, len(users))
	for _, user := range users {
		if user.Active {
			cityMap[user.City] = append(cityMap[user.City], user)
		}
	}
	if len(cityMap) == 0 {
		return nil
	}
	return cityMap
}

func CountTags(users []User) map[string]int {
	// TODO: посчитать все появления тегов у всех пользователей.
	// Повторяющийся тег у одного пользователя также считается отдельным появлением.
	tagsMap := make(map[string]int)
	for _, user := range users {
		for _, tag := range user.Tags {
			tagsMap[tag]++
		}
	}
	return tagsMap
}

func BuildInventory(products []Product) map[string]Product {
	// TODO: построить склад по SKU.
	// При повторном SKU оставить последнюю запись из входного слайса.
	store := make(map[string]Product, len(products))

	for _, product := range products {
		store[product.SKU] = product
	}
	return store
}

func ReserveStock(inventory map[string]Product, sku string, count int) bool {
	// TODO: уменьшить остаток существующего товара на положительное count, если товара хватает.
	// При неуспехе вернуть false и не менять склад.
	if count < 0 {
		return false
	}

	product, ok := inventory[sku]
	if !ok {
		return false
	}

	if product.Quantity < count {
		return false
	}

	product.Quantity -= count
	inventory[sku] = product
	return true
}

func Restock(inventory map[string]*Product, sku string, count int) bool {
	// TODO: увеличить остаток товара, хранящегося в map как указатель.
	// Ключ должен существовать, указатель не должен быть nil, count должен быть положительным.
	if count < 0 {
		return false
	}

	product, ok := inventory[sku]
	if !ok || product == nil {
		return false
	}

	product.Quantity += count //автоматическое (неявное) разыменование указателя, так как структура
	return true
}

func LowStockSKUs(inventory map[string]Product, limit int) []string {
	// TODO: вернуть SKU товаров с остатком не больше limit.
	// Результат должен иметь стабильный алфавитный порядок.
	sku := make([]string, 0, len(inventory))
	for _, v := range inventory {
		if v.Quantity <= limit {
			sku = append(sku, v.SKU)
		}
	}
	slices.Sort(sku)
	return sku
}

func InventoryValue(inventory map[string]Product) int {
	// TODO: посчитать общую стоимость склада как сумму Quantity * Price.
	// Пустой и nil склад имеют стоимость 0.
	if len(inventory) == 0 {
		return 0
	}
	cost := 0

	for _, v := range inventory {
		cost += v.Quantity * v.Price
	}
	return cost
}

func ApplyPriceUpdates(inventory map[string]Product, updates map[string]int) int {
	// TODO: применить положительные новые цены только к существующим товарам.
	// Вернуть число реально обновлённых товаров; неизвестные SKU и неположительные цены пропустить.
	count := 0

	for k, value := range updates {
		if value > 0 {
			product, ok := inventory[k]
			if ok {
				product.Price = value
				inventory[k] = product
				count++
			}
		}
	}
	return count
}
