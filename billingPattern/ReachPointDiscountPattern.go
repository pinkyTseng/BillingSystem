package billingPattern

import (
	"fmt"
	"strconv"
	"strings"
)

type ReachPointDiscountPattern struct {
	pointPercentage    string //ex: 1:1 point:coin
	pointPartRatio     int
	coinPartRatio      int
	pointThreashold    int
	pointReachDiscount int //ex 90
}

type ReachPointDiscountPatternFactory struct {
}

func (p *ReachPointDiscountPattern) ChangeSetting(param EventPatternCreatObj) {
	if param.PointPercentage != "" {
		pointPercentage := param.PointPercentage
		ratioArr := strings.Split(pointPercentage, ":")
		point, _ := strconv.Atoi(ratioArr[0])
		coin, _ := strconv.Atoi(ratioArr[1])

		p.pointPercentage = pointPercentage //ex: 1:1 point:coin
		p.pointPartRatio = point
		p.coinPartRatio = coin
	}

	if param.Others != nil && param.Others["reachPointDiscount"] != nil {
		reachPointDiscountMap := param.Others["reachPointDiscount"].(map[string]int)
		if reachPointDiscountMap["pointThreashold"] != 0 {
			p.pointThreashold = reachPointDiscountMap["pointThreashold"]
		}
		if reachPointDiscountMap["pointReachDiscount"] != 0 {
			p.pointReachDiscount = reachPointDiscountMap["pointReachDiscount"]
		}
	}
}

func (p *ReachPointDiscountPattern) CalculatePrice(param *CalculatePriceParam) {
	coinTotal := param.CoinTotal
	userId := param.UserId
	pointTotal := param.PointTotal

	coinCost := coinTotal - pointTotal*p.coinPartRatio/p.pointPartRatio
	//in real case, need to pass and check the userId to judge whether the user is a VIP
	if pointTotal >= p.pointThreashold {
		coinCost = coinCost * p.pointReachDiscount / 100
	}
	fmt.Printf("user %v should cost  %v points & %v coins\n", userId, pointTotal, coinCost)
}

func (f *ReachPointDiscountPatternFactory) Create(parameters EventPatternCreatObj) (EventPriceType, error) {
	pointPercentage := parameters.PointPercentage
	ratioArr := strings.Split(pointPercentage, ":")
	point, _ := strconv.Atoi(ratioArr[0])
	coin, _ := strconv.Atoi(ratioArr[1])

	reachPointDiscountMap := parameters.Others["reachPointDiscount"].(map[string]int)
	pointThreashold := reachPointDiscountMap["pointThreashold"]
	pointReachDiscount := reachPointDiscountMap["pointReachDiscount"]

	pattern := &ReachPointDiscountPattern{
		pointPercentage:    pointPercentage,
		pointPartRatio:     point,
		coinPartRatio:      coin,
		pointThreashold:    pointThreashold,
		pointReachDiscount: pointReachDiscount,
	}
	return pattern, nil
}

func init() {
	Register("ReachPointDiscountPattern", &ReachPointDiscountPatternFactory{})
}
