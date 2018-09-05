package token

import (
	"os"

	"github.com/CityOfZion/neo-storm/interop/storage"
	"github.com/CityOfZion/neo-storm/interop/util"
	"github.com/studentofjs/moonrock-neo/token/nep5"
	"github.com/subosito/gotenv"
)

const (
	decimals   = 8
	multiplier = 100000000
)

// OwnerAddress is the NEO address of the account owner
var OwnerAddress string

var owner = util.FromAddress(OwnerAddress)

// CreateToken initializes the Token Interface for the Smart Contract to operate with
func CreateToken() nep5.Token {
	return nep5.Token{
		Name:           "Moonrock Token",
		Symbol:         "MRT",
		Decimals:       decimals,
		Owner:          owner,
		TotalSupply:    11000000 * multiplier,
		CirculationKey: "TokenCirculation",
	}
}

// Main function = contract entry
func Main(operation string, args []interface{}) interface{} {
	token := CreateToken()

	if operation == "name" {
		return token.Name
	}
	if operation == "symbol" {
		return token.Symbol
	}
	if operation == "decimals" {
		return token.Decimals
	}

	// The following operations need ctx
	ctx := storage.GetContext()

	if operation == "totalSupply" {
		return token.GetSupply(ctx)
	}
	if operation == "balanceOf" {
		hodler := args[0].([]byte)
		return token.BalanceOf(ctx, hodler)
	}
	if operation == "transfer" && CheckArgs(args, 3) {
		from := args[0].([]byte)
		to := args[1].([]byte)
		amount := args[2].(int)
		return token.Transfer(ctx, from, to, amount)
	}

	return true
}

// CheckArgs checks args array against a length indicator
func CheckArgs(args []interface{}, length int) bool {
	if len(args) == length {
		return true
	}

	return false
}

func init() {
	gotenv.Load()
	OwnerAddress = os.Getenv("OWNER")
}
