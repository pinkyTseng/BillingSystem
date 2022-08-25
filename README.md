# Billing System

## Introduction
主要遵循開放封閉原則並參考搭配使用設計模式,因為有和Oliver確定一下題目重點的需求,所以這邊主要是將結果給印出來,並未實作DB實際扣款扣點及檢查vip資格是否符合的部分;另外為了保持彈性與復用性,優惠內容我都是透過調用者自定義的EventPatternCreatObj來設定,此外因為題目有要求活動時會彈性調整方案B及C的優惠內容,我各優惠方案的程式檔案都有加上ChangeSetting function.另外必須說明一下的是因為這邊結果我主要是用print呈現,這邊的test cases我並沒有做特別的錯誤判斷,這邊的test file我主要是印出結果方便自己確認運行結果並且可以當作使用方法參考的檔案,最後我也有開另一新專案並引入這個module試用過沒問題. ps:因為區塊鏈多是用wei這個非常小有18次方的單位作為最小單位,所以這邊我也預設採用這個概念,計算我就都只用整數計算.

## 使用說明
首先go get此module,除了要新增優惠方案模式的類別會需要新增go程式檔,其他情況一般會用到得都在BillingComponent檔內。主要使用方式可參考test file內如下的consume function,下面是常用的重要function和struct,更詳細用例可參考test file

```
func consume(param *EventPatternCreatObj, calcParam *CalculatePriceParam) {
	priceContext := EventPatternContext(*param)
	priceContext.CalculatePrice(calcParam)
}  

//除了FactoryType,其他不一定都要有,看FactoryType
type EventPatternCreatObj struct {
	FactoryType     string //"CommonPattern", "VipCoinPattern", "PointUsePattern", "ReachPointDiscountPattern"
	VipDiscount     map[int]int //ex: VIP1: 95折 [1]95, 分子整數部分
	PointPercentage string      //ex: 1:1
	Others          map[string]interface{} //保留給擃充的優惠方案始用,ReachPointDiscountPattern有用到
}

type CalculatePriceParam struct {
	CoinTotal  int
	PointTotal int
	VipLevel   int    //ex:0(not VIP), 1, 2, 3
	UserId     string //used to check vip level and execute SQL update
}

EventPatternCreatObj Json表示最全示意例子：
{
  "FactoryType" : "ReachPointDiscountPattern",
  "VipDiscount" : {1:95, 2:90, 3:85}, 
  "PointPercentage" : "1:1",
  "Others" :  {
    "reachPointDiscount" : {
      "pointThreashold" : 100,
      "pointReachDiscount" : 90
    }
  }
}

```