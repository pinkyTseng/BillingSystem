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

func consume(param *EventPatternCreatObj, calcParam *CalculatePriceParam) {
	priceContext := EventPatternContext(*param)
	priceContext.CalculatePrice(calcParam)
}
