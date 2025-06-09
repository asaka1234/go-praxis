package go_praxis

import (
	"github.com/asaka1234/go-praxis/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params *PraxisInitParams

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *PraxisInitParams) *Client {
	return &Client{
		Params: params,

		ryClient:  resty.New(), //client实例
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetDebugModel(debugModel bool) {
	cli.debugMode = debugModel
}
