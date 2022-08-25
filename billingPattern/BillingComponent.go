package billingPattern

import (
	"fmt"
	"log"
)

var EventPriceFactoryMap = make(map[string]EventPriceFactory)

type EventPriceFactory interface {
	Create(parameters EventPatternCreatObj) (EventPriceType, error)
}

func Register(name string, factory EventPriceFactory) {
	EventPriceFactoryMap[name] = factory
}

func EventPatternContext(param EventPatternCreatObj) EventPriceType {
	eventPriceType, err := EventPriceFactoryMap[param.FactoryType].Create(param)
	if err != nil {
		fmt.Printf("some unexpected error happen")
	}
	return eventPriceType
}

type EventPatternCreatObj struct {
	FactoryType     string
	VipDiscount     map[int]int //ex: VIP1: 95折 [1]95
	PointPercentage string      //ex: 1:1
	Others          map[string]interface{}
}
type CalculatePriceParam struct {
	CoinTotal  int
	PointTotal int
	VipLevel   int    //ex:0(not VIP), 1, 2, 3
	UserId     string //used to check vip level and execute SQL update
}

type EventPriceType interface {
	CalculatePrice(param *CalculatePriceParam)
	ChangeSetting(param EventPatternCreatObj)
}

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
