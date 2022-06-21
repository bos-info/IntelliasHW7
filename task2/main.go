package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	//слайс створений з input
	stringSlice := strings.Fields(input)
	//ініціалізуємо максимальне значення МІНімальним значенням для int32 і мінімальне занчення МАКСимальним значенням для інт32
	var max, min int32 = math.MinInt32, math.MaxInt32
	for _, v := range stringSlice {
		// парсимо елемент слайсу в інт
		parsedInt64, err := strconv.ParseInt(v, 10, 32)
		// якщо сталась помилка конвертації , виводимо повідомлення і переходимо до наступного елементу слайсу
		if err != nil {
			log.Println("Помилка вхідних даних, перевірте правильність даних!")
			continue
		}
		// конвертуємо значення що розпарсили з int64 до int32
		parsedInt32 := int32(parsedInt64)
		//визначаємо мінімальне значення
		if parsedInt32 < min {
			min = parsedInt32
		}
		//визначаємо максимальне значення
		if parsedInt32 > max {
			max = parsedInt32
		}
	}
	// якщо лише 1 число, то виводимо тілько його
	if min == max {
		result = strconv.FormatInt(int64(max), 10)
	} else {
		result = strconv.FormatInt(int64(max), 10) + " " + strconv.FormatInt(int64(min), 10)
	}
}
