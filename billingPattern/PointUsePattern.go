package billingPattern

import (
	"fmt"
	"strconv"
	"strings"
)

type PointUsePattern struct {
	pointPercentage string //ex: 1:1 point:coin
	pointPartRatio  int
	coinPartRatio   int
}

type PointUsePatternFactory struct {
	// thePattern VipCoinPattern
}

func (p *PointUsePattern) CalculatePrice(param *CalculatePriceParam) {
	coinTotal := param.coinTotal
	userId := param.userId
	pointTotal := param.pointTotal

	coinCost := coinTotal - pointTotal*p.coinPartRatio/p.pointPartRatio
	fmt.Printf("user %v should cost  %v points & %v coins\n", userId, pointTotal, coinCost)
}

func (f *PointUsePatternFactory) Create(parameters EventPatternCreatObj) (EventPriceType, error) {
	pointPercentage := parameters.pointPercentage
	ratioArr := strings.Split(pointPercentage, ":")
	point, _ := strconv.Atoi(ratioArr[0])
	coin, _ := strconv.Atoi(ratioArr[1])
	pattern := &PointUsePattern{
		pointPercentage: pointPercentage, //ex: 1:1 point:coin
		pointPartRatio:  point,
		coinPartRatio:   coin,
	}
	return pattern, nil
}

func init() {
	Register("PointUsePattern", &PointUsePatternFactory{})
}
