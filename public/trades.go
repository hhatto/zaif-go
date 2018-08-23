package public

type TradesMethod struct {
	Date         int     `json:"date"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Tid          int     `json:"tid"`
	CurrencyPair string  `json:"currency_pair"`
	TradeType    string  `json:"trade_type"`
}

func (api *ApiClient) Trades(currencyPair string) ([]TradesMethod, error) {
	var trades []TradesMethod
	err := api.GetRequest("trades", currencyPair, &trades)
	return trades, err
}
