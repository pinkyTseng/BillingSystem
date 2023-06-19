package billingPattern

import (
	"testing"
)

func TestCommonPattern(t *testing.T) {
	commonObj := &EventPatternCreatObj{
		FactoryType: "CommonPattern",
	}

	commonCalcObj := &CalculatePriceParam{
		CoinTotal: 1000,
		UserId:    "u1",
	}
	expected := 1000
	actual := consume(commonObj, commonCalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}
}

func TestVipCoinPattern(t *testing.T) {
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

	expected := 900
	actual := consume(vipCoin1Obj, vipCoin1CalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}
}

func TestPointUsePattern(t *testing.T) {
	pointUse1Obj := &EventPatternCreatObj{
		FactoryType:     "PointUsePattern",
		PointPercentage: "1:1",
	}

	pointUse1CalcObj := &CalculatePriceParam{
		CoinTotal:  1000,
		PointTotal: 100,
		UserId:     "u6",
	}

	expected := 900
	actual := consume(pointUse1Obj, pointUse1CalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}
}
func TestReachPointDiscountPattern(t *testing.T) {

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

	expected := 720
	actual := consume(pointUse1Obj, pointUse1CalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}
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

	expected := 900
	actual := priceContext.CalculatePrice(vipCoin1CalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}

	vipDiscount2 := make(map[int]int)
	vipDiscount2[1] = 80
	vipDiscount2[2] = 75
	vipDiscount2[3] = 70
	vipCoin1Obj.VipDiscount = vipDiscount2

	priceContext.ChangeSetting(*vipCoin1Obj)

	expected = 750
	actual = priceContext.CalculatePrice(vipCoin1CalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}

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
	expected := 900
	actual := priceContext.CalculatePrice(pointUse1CalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}

	pointUse1Obj.PointPercentage = "1:2"
	priceContext.ChangeSetting(*pointUse1Obj)

	expected = 800
	actual = priceContext.CalculatePrice(pointUse1CalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}
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

	expected := 720
	actual := priceContext.CalculatePrice(pointUse1CalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}

	reachPointDiscount2 := make(map[string]int)
	reachPointDiscount2["pointThreashold"] = 50
	reachPointDiscount2["pointReachDiscount"] = 80

	others2 := make(map[string]interface{})
	others2["reachPointDiscount"] = reachPointDiscount2
	pointUse1Obj.Others = others2

	pointUse1Obj.PointPercentage = "1:2"
	priceContext.ChangeSetting(*pointUse1Obj)

	expected = 480
	actual = priceContext.CalculatePrice(pointUse1CalcObj)
	if expected != actual {
		t.Errorf("CalculatePrice() = %v, want %v", actual, expected)
	}
}

func consume(param *EventPatternCreatObj, calcParam *CalculatePriceParam) int {
	priceContext := EventPatternContext(*param)
	return priceContext.CalculatePrice(calcParam)
}
