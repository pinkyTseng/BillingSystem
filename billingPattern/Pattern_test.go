package billingPattern

import (
	"testing"
)

func TestSomeCase(t *testing.T) {

	commonObj := &EventPatternCreatObj{
		FactoryType: "CommonPattern",
	}

	commonCalcObj := &CalculatePriceParam{
		CoinTotal: 1000,
		UserId:    "u1",
	}

	vipDiscount1 := make(map[int]int)
	vipDiscount1[1] = 95
	vipDiscount1[2] = 90
	vipDiscount1[3] = 85
	vipCoin1Obj := &EventPatternCreatObj{
		FactoryType: "VipCoinPattern",
		VipDiscount: vipDiscount1,
	}

	vipCoin1CalcObj := &CalculatePriceParam{
		CoinTotal: 1000,
		VipLevel:  2,
		UserId:    "u2",
	}

	vipDiscount2 := make(map[int]int)
	vipDiscount2[1] = 80
	vipDiscount2[2] = 75
	vipDiscount2[3] = 70
	vipCoin2Obj := &EventPatternCreatObj{
		FactoryType: "VipCoinPattern",
		VipDiscount: vipDiscount2,
	}

	vipCoin2CalcObj := &CalculatePriceParam{
		CoinTotal: 1000,
		VipLevel:  1,
		UserId:    "u3",
	}

	vipCoin1CalcObj_2 := &CalculatePriceParam{
		CoinTotal: 1000,
		VipLevel:  1,
		UserId:    "u4",
	}

	vipCoin2CalcObj_2 := &CalculatePriceParam{
		CoinTotal: 1000,
		VipLevel:  3,
		UserId:    "u5",
	}

	pointUse1Obj := &EventPatternCreatObj{
		FactoryType:     "PointUsePattern",
		PointPercentage: "1:1",
	}

	pointUse1CalcObj := &CalculatePriceParam{
		CoinTotal:  1000,
		PointTotal: 100,
		UserId:     "u6",
	}

	pointUse2Obj := &EventPatternCreatObj{
		FactoryType:     "PointUsePattern",
		PointPercentage: "1:2",
	}

	pointUse2CalcObj := &CalculatePriceParam{
		CoinTotal:  1000,
		PointTotal: 100,
		UserId:     "u7",
	}

	pointUse1CalcObj_2 := &CalculatePriceParam{
		CoinTotal:  1000,
		PointTotal: 200,
		UserId:     "u8",
	}

	pointUse2CalcObj_2 := &CalculatePriceParam{
		CoinTotal:  1000,
		PointTotal: 200,
		UserId:     "u9",
	}

	consume(commonObj, commonCalcObj)
	consume(vipCoin1Obj, vipCoin1CalcObj)
	consume(vipCoin2Obj, vipCoin2CalcObj)
	consume(vipCoin1Obj, vipCoin1CalcObj_2)
	consume(vipCoin2Obj, vipCoin2CalcObj_2)

	consume(pointUse1Obj, pointUse1CalcObj)
	consume(pointUse2Obj, pointUse2CalcObj)
	consume(pointUse1Obj, pointUse1CalcObj_2)
	consume(pointUse2Obj, pointUse2CalcObj_2)
}

func TestAddedPattern(t *testing.T) {

	reachPointDiscount := make(map[string]int)
	reachPointDiscount["pointThreashold"] = 100
	reachPointDiscount["pointReachDiscount"] = 90

	others := make(map[string]interface{})
	others["reachPointDiscount"] = reachPointDiscount

	pointUse1Obj := &EventPatternCreatObj{
		FactoryType:     "ReachPointDiscountPattern",
		PointPercentage: "1:1",
		Others:          others,
	}

	pointUse1CalcObj := &CalculatePriceParam{
		CoinTotal:  1000,
		PointTotal: 200,
		UserId:     "uu1",
	}

	consume(pointUse1Obj, pointUse1CalcObj)
}

func TestChangeSetting(t *testing.T) {

	vipDiscount1 := make(map[int]int)
	vipDiscount1[1] = 95
	vipDiscount1[2] = 90
	vipDiscount1[3] = 85
	vipCoin1Obj := &EventPatternCreatObj{
		FactoryType: "VipCoinPattern",
		VipDiscount: vipDiscount1,
	}

	vipCoin1CalcObj := &CalculatePriceParam{
		CoinTotal: 1000,
		VipLevel:  2,
		UserId:    "u2",
	}

	priceContext := EventPatternContext(*vipCoin1Obj)
	priceContext.CalculatePrice(vipCoin1CalcObj)

	vipDiscount2 := make(map[int]int)
	vipDiscount2[1] = 80
	vipDiscount2[2] = 75
	vipDiscount2[3] = 70
	vipCoin1Obj.VipDiscount = vipDiscount2

	priceContext.ChangeSetting(*vipCoin1Obj)
	priceContext.CalculatePrice(vipCoin1CalcObj)

}

func TestChangeSettingPointUse(t *testing.T) {
	pointUse1Obj := &EventPatternCreatObj{
		FactoryType:     "PointUsePattern",
		PointPercentage: "1:1",
	}

	pointUse1CalcObj := &CalculatePriceParam{
		CoinTotal:  1000,
		PointTotal: 100,
		UserId:     "u6",
	}

	priceContext := EventPatternContext(*pointUse1Obj)
	priceContext.CalculatePrice(pointUse1CalcObj)

	pointUse1Obj.PointPercentage = "1:2"
	priceContext.ChangeSetting(*pointUse1Obj)
	priceContext.CalculatePrice(pointUse1CalcObj)
}

func TestChangeSettingReachPointDiscount(t *testing.T) {
	reachPointDiscount := make(map[string]int)
	reachPointDiscount["pointThreashold"] = 100
	reachPointDiscount["pointReachDiscount"] = 90

	others := make(map[string]interface{})
	others["reachPointDiscount"] = reachPointDiscount

	pointUse1Obj := &EventPatternCreatObj{
		FactoryType:     "ReachPointDiscountPattern",
		PointPercentage: "1:1",
		Others:          others,
	}

	pointUse1CalcObj := &CalculatePriceParam{
		CoinTotal:  1000,
		PointTotal: 200,
		UserId:     "uu1",
	}

	priceContext := EventPatternContext(*pointUse1Obj)
	priceContext.CalculatePrice(pointUse1CalcObj)

	reachPointDiscount2 := make(map[string]int)
	reachPointDiscount2["pointThreashold"] = 50
	reachPointDiscount2["pointReachDiscount"] = 80

	others2 := make(map[string]interface{})
	others2["reachPointDiscount"] = reachPointDiscount2
	pointUse1Obj.Others = others2

	pointUse1Obj.PointPercentage = "1:2"
	priceContext.ChangeSetting(*pointUse1Obj)
	priceContext.CalculatePrice(pointUse1CalcObj)
}

func consume(param *EventPatternCreatObj, calcParam *CalculatePriceParam) {
	priceContext := EventPatternContext(*param)
	priceContext.CalculatePrice(calcParam)
}
