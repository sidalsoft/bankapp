package stats

import "github.com/sidalsoft/bankapp/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	for _, p := range payments {
		sum += p.Amount
	}
	return sum / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, p := range payments {
		if p.Category != category {
			continue
		}
		sum += p.Amount
	}
	return sum
}
