package billingPattern

import "fmt"

type VipCoinPattern struct {
	vipDiscount map[int]int //ex: VIP1: 95折 [1]95
}

type VipCoinPatternFactory struct {
	// thePattern VipCoinPattern
}

func (p *VipCoinPattern) ChangeSetting(param EventPatternCreatObj) {
	if param.vipDiscount != nil {
		vipDiscountMap := make(map[int]int)
		// vipDiscountMap[0] = 100
		for k, v := range param.vipDiscount {
			vipDiscountMap[k] = v
		}
		vipDiscountMap[0] = 100
		p.vipDiscount = vipDiscountMap
	}
}

func (p *VipCoinPattern) CalculatePrice(param *CalculatePriceParam) {
	coinTotal := param.coinTotal
	userId := param.userId
	vipLevel := param.vipLevel

	//check whether the vipLevel of the userId is correct

	molecular := p.vipDiscount[vipLevel]
	coinCost := coinTotal * molecular / 100
	fmt.Printf("user %v should cost %v coins\n", userId, coinCost)
}

func (f *VipCoinPatternFactory) Create(parameters EventPatternCreatObj) (EventPriceType, error) {
	vipDiscountMap := make(map[int]int)
	// vipDiscountMap[0] = 100
	for k, v := range parameters.vipDiscount {
		vipDiscountMap[k] = v
	}
	vipDiscountMap[0] = 100
	pattern := &VipCoinPattern{
		vipDiscount: vipDiscountMap,
	}
	return pattern, nil
}

func init() {

	Register("VipCoinPattern", &VipCoinPatternFactory{})
	// factory.Register(driverName, &filesystemDriverFactory{})
}
