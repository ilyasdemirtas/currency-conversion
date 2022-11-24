package handlers

import (
	"arf/currency-conversion/api/app"
	"arf/currency-conversion/api/internal/models"
	"arf/currency-conversion/api/internal/repository"
	token "arf/currency-conversion/api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type ExchangeTransactionInput struct {
	BaseCurrency    string          `json:"base_currency" binding:"required"`
	CounterCurrency string          `json:"counter_currency" binding:"required"`
	Amount          decimal.Decimal `json:"amount" binding:"required"`
}

func CreateExchangeTransaction(c *gin.Context) {
	var input ExchangeTransactionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep := repository.NewR(app.GetDbConn())

	offer, err := GetExchangeRateOfferByCurrencyNames(input.BaseCurrency, input.CounterCurrency)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalPrice := offer.Price.Mul(input.Amount)

	hasBalance := rep.CheckUserBalanceByCurrency(userId, offer.CounterCurrencyId, totalPrice)
	if hasBalance == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance for: " + input.CounterCurrency})
		return
	}

	exchangeTransaction := models.ExchangeTransaction{
		BaseCurrencyId:    offer.BaseCurrencyId,
		CounterCurrencyId: offer.CounterCurrencyId,
		Amount:            input.Amount,
		Price:             totalPrice,
		MarkupRate:        offer.MarkupRate,
		UserId:            userId,
	}

	result, err := rep.CreateExchangeTransaction(exchangeTransaction)

	if err != nil && !result {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	baseCurrencyWallet, _ := rep.WalletAccountByUserAndCurrency(userId, offer.BaseCurrencyId)
	counterCurrencyWallet, _ := rep.WalletAccountByUserAndCurrency(userId, offer.CounterCurrencyId)

	baseCurrencyWallet.Balance = baseCurrencyWallet.Balance.Add(input.Amount)
	counterCurrencyWallet.Balance = counterCurrencyWallet.Balance.Sub(totalPrice)

	rep.UpdateUserWallet(baseCurrencyWallet)
	rep.UpdateUserWallet(counterCurrencyWallet)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": exchangeTransaction})

}
