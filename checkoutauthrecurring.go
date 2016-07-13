package paypal

type CheckoutAuthRecurring struct {
  CheckoutAuth
  Price string `url:"PAYMENTREQUEST_0_AMT"`
}