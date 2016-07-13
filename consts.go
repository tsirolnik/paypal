package paypal

const (
	SandboxEndpoint                = "https://api-3t.sandbox.paypal.com/nvp"
	SandboxRedirectURL             = "https://www.sandbox.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token=%s"
	ProductionEndpoint             = "https://api-3t.paypal.com/nvp"
	ProductionRedirectURL          = "https://www.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token=%s"
	MethodCheckoutAuth             = "SetExpressCheckout"
	MethodCheckoutData             = "GetExpressCheckoutDetails"
	MethodRecurringPaymentsProfile = "CreateRecurringPaymentsProfile"
	EnvSandbox                     = "sandbox"
	EnvProduction                  = "production"
	ShippingDisableShipping        = "1"
	ShippingEnableShipping         = "0"
)
