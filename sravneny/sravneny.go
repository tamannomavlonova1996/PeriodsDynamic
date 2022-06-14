package sravneny

import (
	"fmt"
	"srav/types"
)

func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money,
) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	//for k := range first {
	//	categories[k] = second[k] - first[k]
	//}

	for _, val := range first {
		fmt.Println(val)
		//fmt.Println(first[key])
	}

	return categories
}
