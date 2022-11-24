package handlers

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/shopspring/decimal"
)

func TestCalculateMarkup(t *testing.T) {

	var expectedPrice = decimal.NewFromFloat(19.5925)
	var expectedPrice2 = decimal.NewFromFloat(19.5925)

	var price decimal.Decimal = decimal.NewFromFloat(18.5)
	var markup decimal.Decimal = decimal.NewFromFloat(0.5)
	var markupCalculated = price.Div(decimal.NewFromFloat(100)).Mul(markup)
	var newPrice = price.Add(markupCalculated)

	assert.NotEqual(t, expectedPrice, newPrice)
	assert.Equal(t, expectedPrice2, newPrice)

}
