/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"math/big"
	"strconv"
)

const (
	MOD97 = 97
)

func MatchProviderID(fromPath string, fromToken int) bool {
	if fromPath == "" || fromPath == "0" {
		return true
	}

	if fromToken == 0 {
		return true
	}

	fromPathInt, err := strconv.Atoi(fromPath)
	if err != nil {
		return false
	}

	return fromPathInt == fromToken
}

func AppendLeadingZero(number int) string {
	if number < 10 {
		return fmt.Sprintf("0%d", number)
	} else {
		return fmt.Sprintf("%d", number)
	}
}

func WrapQueryResult(count int64, r interface{}) *QueryResult {
	return &QueryResult{
		Amount:   count,
		Elements: r,
	}
}

func GenerateBankNumber(ctx *gin.Context, customerRegistrationNumber string) string {
	banksClearingNumber := "10900004"
	providerReckoningNumber := "6156"
	checkNumber := generateCheckNumber(ctx, fmt.Sprint(banksClearingNumber, providerReckoningNumber, customerRegistrationNumber), "2521")
	return fmt.Sprint(checkNumber, banksClearingNumber, providerReckoningNumber, customerRegistrationNumber)
}

func generateCheckNumber(ctx *gin.Context, bban, countryCode string) string {
	controlBase := fmt.Sprint(bban, countryCode, "00")
	var bignum, ok = new(big.Int).SetString(controlBase, 10)
	if !ok {
		e.HandleError(e.ErrInternalServerError, ctx)
		return ""
	}
	res := bignum.Mod(bignum, big.NewInt(MOD97))
	return strconv.FormatInt(98-res.Int64(), 10)
}
