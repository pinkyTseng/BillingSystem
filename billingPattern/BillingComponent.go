package billingPattern

import (
	// "fmt"
	"fmt"
	"log"
)

var EventPriceFactoryMap = make(map[string]EventPriceFactory)

// type PatternTypeEnum int
// const (
//     commonPattern PatternTypeEnum = iota
//     vipCoinPattern
//     pointUsePattern
// 	reachPointDiscountPattern
// )

// public static UserPayService getByUserType(String type){
// 	return services.get(type);
// }

// public static void register(String userType,UserPayService userPayService){
// 	Assert.notNull(userType,"userType can't be null");
// 	services.put(userType,userPayService);
// }

type EventPriceFactory interface {
	Create(parameters EventPatternCreatObj) (EventPriceType, error)
	// Register(name string, factory EventPriceFactory)
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

// func Create(name string, parameters map[string]interface{}) (storagedriver.StorageDriver, error) {

// 	return driverFactory.Create(parameters)

// }

type EventPatternCreatObj struct {
	FactoryType string
	// eventPriceType  EventPriceType
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
