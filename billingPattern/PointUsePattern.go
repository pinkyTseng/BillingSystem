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
}

func (p *PointUsePattern) ChangeSetting(param EventPatternCreatObj) {
	if param.PointPercentage != "" {
		pointPercentage := param.PointPercentage
		ratioArr := strings.Split(pointPercentage, ":")
		point, _ := strconv.Atoi(ratioArr[0])
		coin, _ := strconv.Atoi(ratioArr[1])

		p.pointPercentage = pointPercentage
		p.pointPartRatio = point
		p.coinPartRatio = coin
	}
}

func (p *PointUsePattern) CalculatePrice(param *CalculatePriceParam) int {
	coinTotal := param.CoinTotal
	userId := param.UserId
	pointTotal := param.PointTotal

	coinCost := coinTotal - pointTotal*p.coinPartRatio/p.pointPartRatio
	fmt.Printf("user %v should cost  %v points & %v coins\n", userId, pointTotal, coinCost)
	return coinCost
}

func (f *PointUsePatternFactory) Create(parameters EventPatternCreatObj) (EventPriceType, error) {
	pointPercentage := parameters.PointPercentage
	ratioArr := strings.Split(pointPercentage, ":")
	point, _ := strconv.Atoi(ratioArr[0])
	coin, _ := strconv.Atoi(ratioArr[1])
	pattern := &PointUsePattern{
		pointPercentage: pointPercentage,
		pointPartRatio:  point,
		coinPartRatio:   coin,
	}
	return pattern, nil
}

func init() {
	Register("PointUsePattern", &PointUsePatternFactory{})
}
