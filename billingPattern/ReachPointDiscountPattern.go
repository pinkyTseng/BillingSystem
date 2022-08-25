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
	if param.pointPercentage != "" {
		pointPercentage := param.pointPercentage
		ratioArr := strings.Split(pointPercentage, ":")
		point, _ := strconv.Atoi(ratioArr[0])
		coin, _ := strconv.Atoi(ratioArr[1])

		p.pointPercentage = pointPercentage //ex: 1:1 point:coin
		p.pointPartRatio = point
		p.coinPartRatio = coin
	}

	if param.others != nil && param.others["reachPointDiscount"] != nil {
		reachPointDiscountMap := param.others["reachPointDiscount"].(map[string]int)
		if reachPointDiscountMap["pointThreashold"] != 0 {
			p.pointThreashold = reachPointDiscountMap["pointThreashold"]
		}
		if reachPointDiscountMap["pointReachDiscount"] != 0 {
			p.pointReachDiscount = reachPointDiscountMap["pointReachDiscount"]
		}
	}
}

func (p *ReachPointDiscountPattern) CalculatePrice(param *CalculatePriceParam) {
	coinTotal := param.coinTotal
	userId := param.userId
	pointTotal := param.pointTotal

	coinCost := coinTotal - pointTotal*p.coinPartRatio/p.pointPartRatio
	if pointTotal >= p.pointThreashold {
		coinCost = coinCost * p.pointReachDiscount / 100
	}
	fmt.Printf("user %v should cost  %v points & %v coins\n", userId, pointTotal, coinCost)
}

func (f *ReachPointDiscountPatternFactory) Create(parameters EventPatternCreatObj) (EventPriceType, error) {
	pointPercentage := parameters.pointPercentage
	ratioArr := strings.Split(pointPercentage, ":")
	point, _ := strconv.Atoi(ratioArr[0])
	coin, _ := strconv.Atoi(ratioArr[1])

	reachPointDiscountMap := parameters.others["reachPointDiscount"].(map[string]int)
	pointThreashold := reachPointDiscountMap["pointThreashold"]
	pointReachDiscount := reachPointDiscountMap["pointReachDiscount"]

	pattern := &ReachPointDiscountPattern{
		pointPercentage:    pointPercentage, //ex: 1:1 point:coin
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
