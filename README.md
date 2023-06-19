# Billing System

## Introduction
這是一個支援多種優惠方案模式且具良好方案可擴充性的小專案,因重點放在可擴充架構的實現這邊主要是將結果給印出來,並未實作DB實際扣款扣點及檢查vip資格是否符合的部分;另外為了保持彈性與復用性,優惠內容我都是透過調用者自定義的EventPatternCreatObj來設定,此外支援有活動時會彈性調整方案B及C的優惠內容,我各優惠方案的程式檔案都有加上ChangeSetting function.另外因為有points和coins不容易自動判斷最佳方案且尚未定義相關判斷準則,這邊是假設UI那會有選項讓User可以選擇自己要使用的付款方案,所以調用方應該是知道自己要call哪一種付款Pattern(strategy)而不是library這邊要自動幫User選一種的方式.這test file在這其實是使用方法參考的檔案,最後我也有開另一新專案並引入這個module試用過沒問題. ps:參考區塊鏈多是用wei這個非常小有18次方的單位作為最小單位,所以這邊我也預設採用這個概念,計算我就都只用整數計算.  
主要參考SOLID並搭配使用設計模式,一開始本來只打算採用策略模式,但因希望盡量讓調用的客戶端不用知道各個策略的實際類別,而是可以用一致的方式調用較方便使用,所以引入了簡單工廠,又因GO動態調用不如Java般方便,所以引入Registe並在init時呼叫執行,整個設計上遵循開放封閉的原則.

## 使用說明
首先go get此module,除了要新增優惠方案模式的類別會需要新增go程式檔,其他情況一般會用到得都在BillingComponent檔內。主要使用方式可參考test file內的consume function,下面是常用的重要function和struct,更詳細用例可參考test file

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

//代表customer帳戶資訊的struct
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