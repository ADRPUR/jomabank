package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	MDL = "MDL"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, MDL:
		return true
	}
	return false
}
