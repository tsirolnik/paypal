package paypal

// RecurringPaymentsProfile Represents a new recurring payments profile
// https://developer.paypal.com/docs/classic/api/merchant/CreateRecurringPaymentsProfile_API_Operation_NVP/
type RecurringPaymentsProfile struct {
	Client            `url:",squash"`
	Method            string  `url:"METHOD"`
	Token             string  `url:"TOKEN"`

	PayerID           string  `url:"PAYERID"`
	ProfileStartDate  string  `url:"PROFILESTARTDATE"`
	Description       string  `url:"DESC"`
	CountryCode       string  `url:"COUNTRYCODE"`

	Price             float64 `url:"AMT"`
	BillingPeriod     string  `url:"BILLINGPERIOD"`
	BillingFrequency  int     `url:"BILLINGFREQUENCY"`
	Currency          string  `url:"CURRENCYCODE"`
	MaxFailedPayments int     `url:"MAXFAILEDPAYMENTS"`
}

// NewRecurringPaymentsProfile Create a new RecurringPaymentsProfile instance
func NewRecurringPaymentsProfile(rpp *RecurringPaymentsProfile) *RecurringPaymentsProfile {
	rpp.Method = MethodRecurringPaymentsProfile
	return rpp
}

// Create Returns a map[string]string containing profileID and profileStatus
func (rpp *RecurringPaymentsProfile) Create() (map[string]string, error) {
	reqValues, err := rpp.Request(rpp)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"profileID":     reqValues.Get("PROFILEID"),
		"profileStatus": reqValues.Get("PROFILESTATUS"),
	}, nil
}
