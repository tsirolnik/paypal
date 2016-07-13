package paypal

type CheckoutData struct {
	Client `url:",squash"`
	Method string `url:"METHOD"`
	Token  string `url:"TOKEN"`
}


// NewCheckoutData Returns a new CheckoutData with the correct method type
func NewCheckoutData(cd *CheckoutData) *CheckoutData {
	cd.Method = MethodCheckoutData
	return cd
}

// Data returns the checkout data that's containting -
// map[string]string {
//  token : "",
//  payerID : "",
//  billingStatus : "",
// }
func (cd *CheckoutData) Data() (map[string]string, error) {
	reqValues, err := cd.Request(cd)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"billingStatus": reqValues.Get("BILLINGAGREEMENTACCEPTEDSTATUS"),
		"payerID":       reqValues.Get("PAYERID"),
		"token":         reqValues.Get("TOKEN"),
	}, nil
}
