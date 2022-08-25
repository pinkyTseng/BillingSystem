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

func (p *CommonPattern) CalculatePrice(param *CalculatePriceParam) {
	coinTotal := param.coinTotal
	userId := param.userId
	coinCost := coinTotal
	fmt.Printf("user %v should cost %v coins\n", userId, coinCost)
}

func (f *CommonPatternFactory) Create(parameters EventPatternCreatObj) (EventPriceType, error) {
	pattern := &CommonPattern{}
	return pattern, nil
}

func init() {

	Register("CommonPattern", &CommonPatternFactory{})
	// factory.Register(driverName, &filesystemDriverFactory{})
}
