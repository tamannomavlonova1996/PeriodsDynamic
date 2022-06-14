package main

import (
	"fmt"
	"srav/sravneny"
	"srav/types"
)

func main() {
	first := map[types.Category]types.Money{
		"auto": 20,
		"food": 30,
	}
	second := map[types.Category]types.Money{
		"auto": 10,
		"food": 50,
	}
	result := sravneny.PeriodsDynamic(first, second)
	fmt.Println(result)
}
