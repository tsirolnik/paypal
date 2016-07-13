package paypal

import (
	"errors"
	"fmt"
)

// CheckoutAuth is a struct used to get a SetExpressCheckout, which is, a payment authorization call.
type CheckoutAuth struct {
	Client          `url:",squash"`
	Method          string  `url:"METHOD"`
	BillingType     string  `url:"L_BILLINGTYPE0"`
	Shipping        string  `url:"NOSHIPPING"`
	Price           float64 `url:"PAYMENTREQUEST_0_AMT"`
	Currency        string  `url:"PAYMENTREQUEST_0_CURRENCYCODE,omitempty"`
	ItemName        string  `url:"L_PAYMENTREQUEST_0_NAME0,omitempty"`
	ItemQty         int     `url:"L_PAYMENTREQUEST_0_QTY0,omitempty"`
	ItemAmount      float64 `url:"L_PAYMENTREQUEST_0_AMT0,omitempty"`
	Category        string  `url:"L_PAYMENTREQUEST_0_ITEMCATEGORY0,omitempty"`
	ConfirmShipping string  `url:"REQCONFIRMSHIPPING,omitempty"`
	NoShipping      string  `url:"NOSHIPPING,omitempty"`
	UserEmail       string  `url:"EMAIL,omitempty"`
	Description     string  `url:"L_BILLINGAGREEMENTDESCRIPTION0"`
	ReturnURL       string  `url:"returnUrl"`
	CancelURL       string  `url:"cancelUrl"`
	token           string
}

// NewCheckoutAuth Returns a new CheckoutAuth with the correct method type
func NewCheckoutAuth(pa *CheckoutAuth) *CheckoutAuth {
	pa.Method = MethodCheckoutAuth
	return pa
}

// Token Returns the generated auth token
func (pa *CheckoutAuth) Token() string {
	return pa.token
}

// RedirectURL Gets the redirect url used to actually let the user pay
func (pa *CheckoutAuth) RedirectURL() (string, error) {
	reqValues, err := pa.Request(pa)
	if err != nil {
		return "", err
	}
	token := reqValues.Get("TOKEN")
	if token == "" {
		return "", errors.New("Not token found.")
	}
	pa.token = token
	return fmt.Sprintf(pa.redirectURL, reqValues.Get("TOKEN")), nil
}
