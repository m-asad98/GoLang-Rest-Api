package api

//Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

//Coin balance response
type CoinBalanceResponse struct {
	//success code, usually 200
	Code int
	//Account balance
	Balance int64
}

//Error response
type Error struct {
	//Error code
	Code int
	//Error message
	Message string
}
