package go_praxis

// 提现的回调处理(传入一个处理函数)
func (cli *Client) WithdrawCallback(req PraxisBackReq, processor func(PraxisBackReq) error) error {
	//验证签名
	//TODO

	//开始处理
	return processor(req)
}
