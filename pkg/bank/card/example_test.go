package card

import (
	"fmt"
	"github.com/sidalsoft/bankapp/pkg/bank/types"
)

func ExampleTotal() {

	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))

	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))

	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))

	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,
		},
	}))
	//Output:
	//100000
	//300000
	//0
	//0
}
func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
			PAN:     "salom",
		},
		{
			Balance: 2_000_00,
			Active:  true,
			PAN:     "depozit",
		},
		{
			Balance: 1_000_00,
			Active:  false,
			PAN:     "alek",
		},
		{
			Balance: -1_000_00,
			Active:  true,
			PAN:     "qwerty",
		},
	}
	r := PaymentSources(cards)
	for _, source := range r {
		fmt.Println(source.Number)
	}
	//Output:
	//salom
	//depozit
	//
	//
}
