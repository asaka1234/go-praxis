package go_praxis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestDeposit(t *testing.T) {

	//构造client
	cli := NewClient(nil, MERCHANT_ID, MERCHANT_SECRET, APPLICATION_KEY, API_VERSION, API_LOCALE, SANDBOX_URL)

	//启动webserver
	go StartWebServer(cli)

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%s, %s\n", resp.RedirectURL, resp.Session.AuthToken)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	<-quit

}

func StartWebServer(cli *Client) {
	// 1. 创建 Gin 引擎（默认带有 Logger 和 Recovery 中间件）
	r := gin.Default()

	// 2. 定义路由
	r.POST("/deposit/back", func(c *gin.Context) {
		var req PraxisBackReq
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Printf("---1-------\n")
			return
		}
		gtAuth := c.GetHeader(SIGN_HEAD_NAME) //收到的req也是有签名的

		//------开始处理------------

		err := cli.CashierCallback(req, gtAuth, func(PraxisBackReq) error { return nil })
		if err != nil {
			fmt.Printf("---2-------\n")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	r.GET("/deposit/page", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	r.Run(":8080") //监听8080端口
}

func GenDepositRequestDemo() PraxisDepositReq {
	return PraxisDepositReq{
		Currency:        "USD",                                     //币种
		Amount:          100,                                       //这个是用 分 为单位的. 有的currency是100分,有的1000分. 具体要看 https://doc.praxiscashier.com/integration_docs/latest/overview/data_formats#currency_fraction
		Cid:             "12",                                      //Unique customer id in your system. 业务系统里的唯一客户id
		NotificationURL: "http://18.162.184.178:8080/deposit/back", //回调通知接口
		ReturnURL:       "http://18.162.184.178:8080/deposit/page", //回调通知接口
		OrderID:         "tkal-1560610956",                         //业务系统内的唯一订单id
		//Timestamp:       time.Now().Unix(),                                      //seconds
		CustomerData: PraxisDepositCustomerData{ //也必填
			Country:   "GB",
			FirstName: "Casper",
			LastName:  "Li",
			DOB:       "1985-01-02", // Date of birth (format: YYYY-MM-DD)
			Email:     "demo@gmail.com",
			Phone:     "44201113223", // Should include country code
			//Zip:       "WC2N 5DU",    // Postal/ZIP code
			//State:     "JS",          // State/Province
			//City:      "London",           // City
			//Address:   "Random st., 12/3", // Street address
			//Profile   int    `json:"profile"` // Street address
		},
		//option
		//CustomerToken: "123456", //客户身份id的HASH
		/*
			PaymentMethod: "",
			Gateway:       "",
			ValidationURL: "",
		*/
	}
}
