package billingPattern

import "fmt"

type VipCoinPattern struct {
	vipDiscount map[int]int //ex: VIP1: 95折 [1]95
}
type VipCoinPatternFactory struct {
}

func (p *VipCoinPattern) ChangeSetting(param EventPatternCreatObj) {
	if param.VipDiscount != nil {
		vipDiscountMap := make(map[int]int)
		for k, v := range param.VipDiscount {
			vipDiscountMap[k] = v
		}
		vipDiscountMap[0] = 100
		p.vipDiscount = vipDiscountMap
	}
}

func (p *VipCoinPattern) CalculatePrice(param *CalculatePriceParam) {
	coinTotal := param.CoinTotal
	userId := param.UserId
	vipLevel := param.VipLevel

	//check whether the vipLevel of the userId is correct in real app
	molecular := p.vipDiscount[vipLevel]
	coinCost := coinTotal * molecular / 100
	fmt.Printf("user %v should cost %v coins\n", userId, coinCost)
}

func (f *VipCoinPatternFactory) Create(parameters EventPatternCreatObj) (EventPriceType, error) {
	vipDiscountMap := make(map[int]int)
	for k, v := range parameters.VipDiscount {
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
}
