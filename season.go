package main

import (
	"fmt"
	"strings"
)

func main() {
	var seasonSelect string
	fmt.Scan(&seasonSelect)
	seasonSelect = strings.ToLower(seasonSelect)
	// Привести все вводимые варианты к нижнему регистру, для корректного сравнения в функции

	switch seasonSelect {
	case "декабрь", "январь", "февраль":
		fmt.Printf("%s%s является месяцем Зимы\n", strings.ToUpper(seasonSelect[:2]), seasonSelect[2:])
	case "март", "апрель", "май":
		fmt.Printf("%s ялвяется месяцем Весны\n", strings.ToUpper(seasonSelect))
	case "июнь", "июль", "август":
		fmt.Printf("%s ялвяется месяцем Лета\n", seasonSelect)
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Printf("%s является месяцем Осени\n", seasonSelect)
	}
	// 1-й 123412341234 20340230402340
	// 2-й 1234555151515 92349234923492394 123
}
