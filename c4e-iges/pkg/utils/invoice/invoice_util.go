package invoiceutil

import (
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"math"
)

func FormatFloat(number float64) string {
	return fmt.Sprintf("%v", number)
}

func FormatFloat2(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

func RoundFloat(value float64) float64 {
	return math.Round(value*100) / 100
}

func FormatReadType(readType invoice.ReadType) string {
	switch readType {
	case invoice.DirectRead:
		return "F"
	case invoice.ProvidedRead:
		return "O"
	case invoice.EstimatedRead:
		return "S"
	case invoice.RemoteRead:
		return "Z"
	default:
		return ""
	}
}
