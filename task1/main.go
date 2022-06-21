package main

func main() {
	//вхідний слайс
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	//результуючий слайс
	var result []int
	//тимчасова мапа для виявлення дублів
	translateMap := make(map[int]int)
	//цикл проходить по елементам масиву
	for i, v := range arr {
		// перевіряємо чи існує елемент мапи з ключем що дорівнює значенню елемента масиву, якщо ні додаємо до мапи елемент та записуємо це значення в результуючий слайс
		if _, ok := translateMap[v]; !ok {
			translateMap[v] = i
			result = append(result, v)
		}
	}
}
