package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))
	//Output: 100000
}
