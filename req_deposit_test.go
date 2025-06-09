package go_praxis

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestDeposit(t *testing.T) {

	//构造client
	cli := NewClient(nil, &PraxisInitParams{MERCHANT_ID, MERCHANT_SECRET, APPLICATION_KEY, API_VERSION, API_LOCALE, SANDBOX_URL, DepositBackUrl, DepositFeBackUrl, WithdrawBackUrl, WithdrawFeBackUrl})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("===>%+v\n", resp)

	fmt.Printf("resp:%s, %s\n", resp.RedirectURL, resp.Session.AuthToken)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	<-quit

}

func GenDepositRequestDemo() PraxisDepositReq {
	return PraxisDepositReq{
		Currency: "USD",             //币种
		Amount:   100,               //这个是用 分 为单位的. 有的currency是100分,有的1000分. 具体要看 https://doc.praxiscashier.com/integration_docs/latest/overview/data_formats#currency_fraction
		Cid:      "12",              //Unique customer id in your system. 业务系统里的唯一客户id
		OrderID:  "tkal-1560610953", //业务系统内的唯一订单id
		//Timestamp:       time.Now().Unix(),                                      //seconds
		CustomerData: PraxisDepositCustomerData{ //也必填
			Country:  "GB",
			LastName: "Li",
			DOB:      "19850102", // Date of birth (format: YYYY-MM-DD)
			Email:    "demo@gmail.com",
		},
	}
}
