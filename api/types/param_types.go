package types

type CoinBalanceRequest struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    int   // response code, usually 200
	Balance int64 // account balance
}
