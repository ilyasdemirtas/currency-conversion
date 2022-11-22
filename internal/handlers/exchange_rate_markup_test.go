package handlers

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestExchangeOfferWithMarkup(t *testing.T) {

	var expectedPrice = decimal.NewFromFloat(18.5925)

	var price decimal.Decimal = decimal.NewFromFloat(18.5)
	var markup decimal.Decimal = decimal.NewFromFloat(0.5)
	var markupCalculated = price.Div(decimal.NewFromFloat(100)).Mul(markup)
	var newPrice = price.Add(markupCalculated)

	if !newPrice.Equal(expectedPrice) {
		t.Errorf("price + ((price / 100) * markup) = %q; want %q", newPrice, expectedPrice)
	}

}
