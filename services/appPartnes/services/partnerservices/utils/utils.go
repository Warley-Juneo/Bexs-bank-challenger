package utils

func ValidatedCurrency(currency string) bool {
	allowedCurrencies := []string{"GBP", "EUR", "USD"}

	for _, allowedCurrency := range allowedCurrencies {
		if allowedCurrency == currency {
			return true
		}
	}
	return false
}