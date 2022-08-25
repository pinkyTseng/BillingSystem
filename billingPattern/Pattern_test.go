package billingPattern

import (
	"testing"
)

func TestSomeCase(t *testing.T) {

	commonObj := &EventPatternCreatObj{
		factoryType: "CommonPattern",
	}

	commonCalcObj := &CalculatePriceParam{
		coinTotal: 1000,
		// pointTotal int
		// vipLevel   int    //ex:0(not VIP), 1, 2, 3
		userId: "u1", //used to check vip level and execute SQL update
	}

	vipDiscount1 := make(map[int]int)
	vipDiscount1[1] = 95
	vipDiscount1[2] = 90
	vipDiscount1[3] = 85
	vipCoin1Obj := &EventPatternCreatObj{
		factoryType: "VipCoinPattern",
		vipDiscount: vipDiscount1,
	}

	vipCoin1CalcObj := &CalculatePriceParam{
		coinTotal: 1000,
		// pointTotal int
		vipLevel: 2,
		userId:   "u2",
	}

	vipDiscount2 := make(map[int]int)
	vipDiscount2[1] = 80
	vipDiscount2[2] = 75
	vipDiscount2[3] = 70
	vipCoin2Obj := &EventPatternCreatObj{
		factoryType: "VipCoinPattern",
		vipDiscount: vipDiscount2,
	}

	vipCoin2CalcObj := &CalculatePriceParam{
		coinTotal: 1000,
		// pointTotal int
		vipLevel: 1,
		userId:   "u3",
	}

	vipCoin1CalcObj_2 := &CalculatePriceParam{
		coinTotal: 1000,
		// pointTotal int
		vipLevel: 1,
		userId:   "u4",
	}

	vipCoin2CalcObj_2 := &CalculatePriceParam{
		coinTotal: 1000,
		// pointTotal int
		vipLevel: 3,
		userId:   "u5",
	}

	pointUse1Obj := &EventPatternCreatObj{
		factoryType:     "PointUsePattern",
		pointPercentage: "1:1",
	}

	pointUse1CalcObj := &CalculatePriceParam{
		coinTotal:  1000,
		pointTotal: 100,
		// vipLevel: 2,
		userId: "u6",
	}

	pointUse2Obj := &EventPatternCreatObj{
		factoryType:     "PointUsePattern",
		pointPercentage: "1:2",
	}

	pointUse2CalcObj := &CalculatePriceParam{
		coinTotal:  1000,
		pointTotal: 100,
		// vipLevel: 2,
		userId: "u7",
	}

	pointUse1CalcObj_2 := &CalculatePriceParam{
		coinTotal:  1000,
		pointTotal: 200,
		// vipLevel: 2,
		userId: "u8",
	}

	pointUse2CalcObj_2 := &CalculatePriceParam{
		coinTotal:  1000,
		pointTotal: 200,
		// vipLevel: 2,
		userId: "u9",
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
		factoryType:     "ReachPointDiscountPattern",
		pointPercentage: "1:1",
		others:          others,
	}

	pointUse1CalcObj := &CalculatePriceParam{
		coinTotal:  1000,
		pointTotal: 200,
		// vipLevel: 2,
		userId: "uu1",
	}

	// pointUse2Obj := &EventPatternCreatObj{
	// 	factoryType:     "PointUsePattern",
	// 	pointPercentage: "1:2",
	// }

	// pointUse2CalcObj := &CalculatePriceParam{
	// 	coinTotal:  1000,
	// 	pointTotal: 100,
	// 	// vipLevel: 2,
	// 	userId: "u7",
	// }

	// pointUse1CalcObj_2 := &CalculatePriceParam{
	// 	coinTotal:  1000,
	// 	pointTotal: 200,
	// 	// vipLevel: 2,
	// 	userId: "u8",
	// }

	// pointUse2CalcObj_2 := &CalculatePriceParam{
	// 	coinTotal:  1000,
	// 	pointTotal: 200,
	// 	// vipLevel: 2,
	// 	userId: "u9",
	// }

	consume(pointUse1Obj, pointUse1CalcObj)
	// consume(pointUse2Obj, pointUse2CalcObj)
	// consume(pointUse1Obj, pointUse1CalcObj_2)
	// consume(pointUse2Obj, pointUse2CalcObj_2)

}

func TestChangeSetting(t *testing.T) {

	vipDiscount1 := make(map[int]int)
	vipDiscount1[1] = 95
	vipDiscount1[2] = 90
	vipDiscount1[3] = 85
	vipCoin1Obj := &EventPatternCreatObj{
		factoryType: "VipCoinPattern",
		vipDiscount: vipDiscount1,
	}

	vipCoin1CalcObj := &CalculatePriceParam{
		coinTotal: 1000,
		// pointTotal int
		vipLevel: 2,
		userId:   "u2",
	}

	priceContext := EventPatternContext(*vipCoin1Obj)
	priceContext.CalculatePrice(vipCoin1CalcObj)

	vipDiscount2 := make(map[int]int)
	vipDiscount2[1] = 80
	vipDiscount2[2] = 75
	vipDiscount2[3] = 70
	vipCoin1Obj.vipDiscount = vipDiscount2

	priceContext.ChangeSetting(*vipCoin1Obj)
	priceContext.CalculatePrice(vipCoin1CalcObj)

}

func TestChangeSettingPointUse(t *testing.T) {
	pointUse1Obj := &EventPatternCreatObj{
		factoryType:     "PointUsePattern",
		pointPercentage: "1:1",
	}

	pointUse1CalcObj := &CalculatePriceParam{
		coinTotal:  1000,
		pointTotal: 100,
		// vipLevel: 2,
		userId: "u6",
	}

	priceContext := EventPatternContext(*pointUse1Obj)
	priceContext.CalculatePrice(pointUse1CalcObj)

	pointUse1Obj.pointPercentage = "1:2"
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
		factoryType:     "ReachPointDiscountPattern",
		pointPercentage: "1:1",
		others:          others,
	}

	pointUse1CalcObj := &CalculatePriceParam{
		coinTotal:  1000,
		pointTotal: 200,
		// vipLevel: 2,
		userId: "uu1",
	}

	priceContext := EventPatternContext(*pointUse1Obj)
	priceContext.CalculatePrice(pointUse1CalcObj)

	reachPointDiscount2 := make(map[string]int)
	reachPointDiscount2["pointThreashold"] = 50
	reachPointDiscount2["pointReachDiscount"] = 80

	others2 := make(map[string]interface{})
	others2["reachPointDiscount"] = reachPointDiscount2
	pointUse1Obj.others = others2

	pointUse1Obj.pointPercentage = "1:2"
	priceContext.ChangeSetting(*pointUse1Obj)
	priceContext.CalculatePrice(pointUse1CalcObj)

}

func consume(param *EventPatternCreatObj, calcParam *CalculatePriceParam) {
	priceContext := EventPatternContext(*param)
	priceContext.CalculatePrice(calcParam)
}
