package billingPattern

import (
	// "fmt"
	"fmt"
	"log"
)

var EventPriceFactoryMap = make(map[string]EventPriceFactory)

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
	eventPriceType, err := EventPriceFactoryMap[param.factoryType].Create(param)
	if err != nil {
		fmt.Printf("some unexpected error happen")
	}
	return eventPriceType
}

// func Create(name string, parameters map[string]interface{}) (storagedriver.StorageDriver, error) {

// 	return driverFactory.Create(parameters)

// }

type EventPatternCreatObj struct {
	factoryType     string
	eventPriceType  EventPriceType
	vipDiscount     map[int]int //ex: VIP1: 95折 [1]95
	pointPercentage string      //ex: 1:1
	//others map[string]interface{}
}
type CalculatePriceParam struct {
	coinTotal  int
	pointTotal int
	vipLevel   int    //ex:0(not VIP), 1, 2, 3
	userId     string //used to check vip level and execute SQL update
}

type EventPriceType interface {
	CalculatePrice(param *CalculatePriceParam)
}

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
