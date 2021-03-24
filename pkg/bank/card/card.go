package card

import (
	"bank/pkg/bank/types"
)

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active || card.Balance <= 0 {
			continue
		}
		sum += card.Balance
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	r := make([]types.PaymentSource, 0)
	for _, card := range cards {
		if !card.Active || card.Balance <= 0 {
			continue
		}
		r = append(r, types.PaymentSource{
			Type:    "card",
			Number:  string(card.PAN),
			Balance: card.Balance,
		})
	}
	return r
}
