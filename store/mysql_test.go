package store

import (
	"httpLogInGo/parsers"
	"log"
	"testing"
)

func TestCheckTableIsOk(t *testing.T) {
	res, err := CheckTableIsOk()
	if err != nil {
		t.Error(err)
	}
	if res != true {
		t.Error("结果不符合！")
	}
}

func TestInsertOneLog(t *testing.T) {
	oneLog := `{"level":"info","time":"2019-03-14T16:06:30.401+0800","linenum":"/home/www/go/src/httpLogInGo/server.go:63","msg":"{\"type\":\"INFO\",\"ip\":\"134.159.204.85\",\"time\":\"2019-03-14 16:06:28\",\"event\":\"真皮支付测试版\",\"key\":\"ZhenPiPay_request_param\",\"request\":\"{\\\"zp_mer_id\\\":\\\"11155925\\\",\\\"zp_order_id\\\":\\\"20190314160628524956\\\",\\\"zp_order_amount\\\":1,\\\"zp_notify_url\\\":\\\"http://dd.jaapanapi.com/Pay_ZhenPiPay_notifyurl.html\\\",\\\"zp_back_url\\\":\\\"http://dd.jaapanapi.com/Pay_ZhenPiPay_callbackurl.html\\\",\\\"zp_pay_type\\\":\\\"ali2bank\\\",\\\"zp_sign\\\":\\\"e762cb09c91bd35046afb435ef3a737a\\\",\\\"zp_desc\\\":\\\"\\\",\\\"zp_extra\\\":\\\"\\\"}\",\"response\":\"{\\\"addOrderReturnData\\\":{\\\"mch_id\\\":\\\"11155925\\\",\\\"signkey\\\":\\\"ZmbxapRDwdmBIslhWgCPcKjFaTeFAIXI\\\",\\\"appid\\\":\\\"11155925\\\",\\\"appsecret\\\":\\\"ZmbxapRDwdmBIslhWgCPcKjFaTeFAIXI\\\",\\\"gateway\\\":\\\"https://api.zhenpipay.com/order/request\\\",\\\"notifyurl\\\":\\\"http://dd.jaapanapi.com/Pay_ZhenPiPay_notifyurl.html\\\",\\\"callbackurl\\\":\\\"http://dd.jaapanapi.com/Pay_ZhenPiPay_callbackurl.html\\\",\\\"unlockdomain\\\":\\\"\\\",\\\"amount\\\":100,\\\"bankcode\\\":\\\"935\\\",\\\"code\\\":\\\"AlipayScan\\\",\\\"orderid\\\":\\\"20190314160628524956\\\",\\\"out_trade_id\\\":\\\"47842\\\",\\\"subject\\\":\\\"47842\\\",\\\"datetime\\\":\\\"2019-03-14 16:06:28\\\",\\\"status\\\":\\\"success\\\"}}\"}","serviceName":"serviceName"}`
	var log1 = parsers.LogInfo{}
	log1, err := parsers.ParseLogInfo([]byte(oneLog))
	if err != nil {
		t.Error(err)
	}
	cnt, err := InsertOneLogMsg(log1.Msg)
	if err != nil {
		t.Error(err)
	}
	log.Println("insert后影响的行数是：", cnt)
}
