package billingPattern

import (
	"fmt"
)

type CommonPatternFactory struct {
}

type CommonPattern struct {
}

func (p *CommonPattern) ChangeSetting(param EventPatternCreatObj) {

}

func (p *CommonPattern) CalculatePrice(param *CalculatePriceParam) int {
	coinTotal := param.CoinTotal
	userId := param.UserId
	coinCost := coinTotal
	fmt.Printf("user %v should cost %v coins\n", userId, coinCost)
	return coinCost
}

func (f *CommonPatternFactory) Create(parameters EventPatternCreatObj) (EventPriceType, error) {
	pattern := &CommonPattern{}
	return pattern, nil
}

func init() {
	Register("CommonPattern", &CommonPatternFactory{})
}
