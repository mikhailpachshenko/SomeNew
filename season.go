package main

import (
	"fmt"
	"strings"
)

func main() {
	var seasonSelect string
	fmt.Scan(&seasonSelect)
	seasonSelect = strings.ToLower(seasonSelect)
	// 1
	// 2
	// 3

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

}
