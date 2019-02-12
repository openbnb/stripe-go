// Package client provides a Stripe client for invoking APIs across all resources
package client

import (
	stripe "github.com/openbnb/stripe-go/v55"
	"github.com/openbnb/stripe-go/v55/account"
	"github.com/openbnb/stripe-go/v55/applepaydomain"
	"github.com/openbnb/stripe-go/v55/balance"
	"github.com/openbnb/stripe-go/v55/bankaccount"
	"github.com/openbnb/stripe-go/v55/bitcoinreceiver"
	"github.com/openbnb/stripe-go/v55/bitcointransaction"
	"github.com/openbnb/stripe-go/v55/card"
	"github.com/openbnb/stripe-go/v55/charge"
	"github.com/openbnb/stripe-go/v55/countryspec"
	"github.com/openbnb/stripe-go/v55/coupon"
	"github.com/openbnb/stripe-go/v55/customer"
	"github.com/openbnb/stripe-go/v55/discount"
	"github.com/openbnb/stripe-go/v55/dispute"
	"github.com/openbnb/stripe-go/v55/ephemeralkey"
	"github.com/openbnb/stripe-go/v55/event"
	"github.com/openbnb/stripe-go/v55/exchangerate"
	"github.com/openbnb/stripe-go/v55/fee"
	"github.com/openbnb/stripe-go/v55/feerefund"
	"github.com/openbnb/stripe-go/v55/file"
	"github.com/openbnb/stripe-go/v55/filelink"
	"github.com/openbnb/stripe-go/v55/invoice"
	"github.com/openbnb/stripe-go/v55/invoiceitem"
	"github.com/openbnb/stripe-go/v55/issuerfraudrecord"
	"github.com/openbnb/stripe-go/v55/issuing/authorization"
	issuingcard "github.com/openbnb/stripe-go/v55/issuing/card"
	"github.com/openbnb/stripe-go/v55/issuing/cardholder"
	issuingdispute "github.com/openbnb/stripe-go/v55/issuing/dispute"
	"github.com/openbnb/stripe-go/v55/issuing/transaction"
	"github.com/openbnb/stripe-go/v55/loginlink"
	"github.com/openbnb/stripe-go/v55/order"
	"github.com/openbnb/stripe-go/v55/orderreturn"
	"github.com/openbnb/stripe-go/v55/paymentintent"
	"github.com/openbnb/stripe-go/v55/paymentsource"
	"github.com/openbnb/stripe-go/v55/payout"
	"github.com/openbnb/stripe-go/v55/plan"
	"github.com/openbnb/stripe-go/v55/product"
	"github.com/openbnb/stripe-go/v55/radar/valuelist"
	"github.com/openbnb/stripe-go/v55/radar/valuelistitem"
	"github.com/openbnb/stripe-go/v55/recipient"
	"github.com/openbnb/stripe-go/v55/refund"
	"github.com/openbnb/stripe-go/v55/reporting/reportrun"
	"github.com/openbnb/stripe-go/v55/reporting/reporttype"
	"github.com/openbnb/stripe-go/v55/reversal"
	"github.com/openbnb/stripe-go/v55/sigma/scheduledqueryrun"
	"github.com/openbnb/stripe-go/v55/sku"
	"github.com/openbnb/stripe-go/v55/source"
	"github.com/openbnb/stripe-go/v55/sourcetransaction"
	"github.com/openbnb/stripe-go/v55/sub"
	"github.com/openbnb/stripe-go/v55/subitem"
	terminalconnectiontoken "github.com/openbnb/stripe-go/v55/terminal/connectiontoken"
	terminallocation "github.com/openbnb/stripe-go/v55/terminal/location"
	terminalreader "github.com/openbnb/stripe-go/v55/terminal/reader"
	"github.com/openbnb/stripe-go/v55/token"
	"github.com/openbnb/stripe-go/v55/topup"
	"github.com/openbnb/stripe-go/v55/transfer"
	"github.com/openbnb/stripe-go/v55/usagerecord"
	"github.com/openbnb/stripe-go/v55/webhookendpoint"
)

// API is the Stripe client. It contains all the different resources available.
type API struct {
	// Account is the client used to invoke /accounts APIs.
	Account *account.Client
	// ApplePayDomains is the client used to invoke /apple_pay/domains APIs.
	ApplePayDomains *applepaydomain.Client
	// Balance is the client used to invoke /balance and transaction-related APIs.
	Balance *balance.Client
	// BankAccounts is the client used to invoke bank account related APIs.
	BankAccounts *bankaccount.Client
	// BitcoinReceivers is the client used to invoke /bitcoin/receivers APIs.
	BitcoinReceivers *bitcoinreceiver.Client
	// BitcoinTransactions is the client used to invoke /bitcoin/transactions APIs.
	BitcoinTransactions *bitcointransaction.Client
	// Cards is the client used to invoke card related APIs.
	Cards *card.Client
	// Charges is the client used to invoke /charges APIs.
	Charges *charge.Client
	// CountrySpec is the client used to invoke /country_specs APIs.
	CountrySpec *countryspec.Client
	// Coupons is the client used to invoke /coupons APIs.
	Coupons *coupon.Client
	// Customers is the client used to invoke /customers APIs.
	Customers *customer.Client
	// Discounts is the client used to invoke discount related APIs.
	Discounts *discount.Client
	// Disputes is the client used to invoke /disputes APIs.
	Disputes *dispute.Client
	// EphemeralKeys is the client used to invoke /ephemeral_keys APIs.
	EphemeralKeys *ephemeralkey.Client
	// Events is the client used to invoke /events APIs.
	Events *event.Client
	// ExchangeRates is the client used to invoke /exchange_rates APIs.
	ExchangeRates *exchangerate.Client
	// Fees is the client used to invoke /application_fees APIs.
	Fees *fee.Client
	// FeeRefunds is the client used to invoke /application_fees/refunds APIs.
	FeeRefunds *feerefund.Client
	// Files is the client used to invoke the /files APIs.
	Files *file.Client
	// FileLinks is the client used to invoke the /file_links APIs.
	FileLinks *filelink.Client
	// Invoices is the client used to invoke /invoices APIs.
	Invoices *invoice.Client
	// InvoiceItems is the client used to invoke /invoiceitems APIs.
	InvoiceItems *invoiceitem.Client
	// IssuerFraudRecords is the client used to invoke /issuer_fraud_records APIs.
	IssuerFraudRecords *issuerfraudrecord.Client
	// IssuingAuthorizations is the client used to invoke /issuing/authorizations APIs.
	IssuingAuthorizations *authorization.Client
	// IssuingCardholders is the client used to invoke /issuing/cardholders APIs.
	IssuingCardholders *cardholder.Client
	// IssuingCards is the client used to invoke /issuing/cards APIs.
	IssuingCards *issuingcard.Client
	// IssuingDisputes is the client used to invoke /issuing/disputes APIs.
	IssuingDisputes *issuingdispute.Client
	// IssuingTransactions is the client used to invoke /issuing/transactions APIs.
	IssuingTransactions *transaction.Client
	// LoginLinks is the client used to invoke login link related APIs.
	LoginLinks *loginlink.Client
	// Orders is the client used to invoke /orders APIs.
	Orders *order.Client
	// OrderReturns is the client used to invoke /order_returns APIs.
	OrderReturns *orderreturn.Client
	// PaymentIntents is the client used to invoke /payment_intents APIs.
	PaymentIntents *paymentintent.Client
	// PaymentSource is used to invoke customer sources related APIs.
	PaymentSource *paymentsource.Client
	// Payouts is the client used to invoke /payouts APIs.
	Payouts *payout.Client
	// Plans is the client used to invoke /plans APIs.
	Plans *plan.Client
	// Products is the client used to invoke /products APIs.
	Products *product.Client
	// RadarValueLists is the client used to invoke /radar/value_lists APIs.
	RadarValueLists *valuelist.Client
	// RadarValueListItems is the client used to invoke /radar/value_list_items APIs.
	RadarValueListItems *valuelistitem.Client
	// Recipients is the client used to invoke /recipients APIs.
	Recipients *recipient.Client
	// Refunds is the client used to invoke /refunds APIs.
	Refunds *refund.Client
	// ReportRuns is the client used to invoke /reporting/report_runs APIs.
	ReportRuns *reportrun.Client
	// ReportTypes is the client used to invoke /reporting/report_types APIs.
	ReportTypes *reporttype.Client
	// Reversals is the client used to invoke /transfers/reversals APIs.
	Reversals *reversal.Client
	// SigmaScheduledQueryRuns is the client used to invoke /sigma/scheduled_query_runs APIs.
	SigmaScheduledQueryRuns *scheduledqueryrun.Client
	// Skus is the client used to invoke /skus APIs.
	Skus *sku.Client
	// Sources is the client used to invoke /sources APIs.
	Sources *source.Client
	// SourceTransactions is the client used to invoke source transaction related APIs.
	SourceTransactions *sourcetransaction.Client
	// Subscriptions is the client used to invoke /subscriptions APIs.
	Subscriptions *sub.Client
	// SubscriptionItems is the client used to invoke subscription's items related APIs.
	SubscriptionItems *subitem.Client
	// TerminalConnectionTokens is the client used to invoke /terminal/connectiontokens related APIs.
	TerminalConnectionTokens *terminalconnectiontoken.Client
	// TerminalLocations is the client used to invoke /terminal/locations related APIs.
	TerminalLocations *terminallocation.Client
	// TerminalReaders is the client used to invoke /terminal/readers related APIs.
	TerminalReaders *terminalreader.Client
	// Tokens is the client used to invoke /tokens APIs.
	Tokens *token.Client
	// Topups is the client used to invoke /tokens APIs.
	Topups *topup.Client
	// Transfers is the client used to invoke /transfers APIs.
	Transfers *transfer.Client
	// UsageRecords is the client used to invoke usage record related APIs.
	UsageRecords *usagerecord.Client
	// WebhookEndpoints is the client used to invoke usage record related APIs.
	WebhookEndpoints *webhookendpoint.Client
}

// Init initializes the Stripe client with the appropriate secret key
// as well as providing the ability to override the backend as needed.
func (a *API) Init(key string, backends *stripe.Backends) {
	if backends == nil {
		backends = &stripe.Backends{
			API:     stripe.GetBackend(stripe.APIBackend),
			Uploads: stripe.GetBackend(stripe.UploadsBackend),
		}
	}

	a.Account = &account.Client{B: backends.API, Key: key}
	a.ApplePayDomains = &applepaydomain.Client{B: backends.API, Key: key}
	a.Balance = &balance.Client{B: backends.API, Key: key}
	a.BankAccounts = &bankaccount.Client{B: backends.API, Key: key}
	a.BitcoinReceivers = &bitcoinreceiver.Client{B: backends.API, Key: key}
	a.BitcoinTransactions = &bitcointransaction.Client{B: backends.API, Key: key}
	a.Cards = &card.Client{B: backends.API, Key: key}
	a.Charges = &charge.Client{B: backends.API, Key: key}
	a.CountrySpec = &countryspec.Client{B: backends.API, Key: key}
	a.Coupons = &coupon.Client{B: backends.API, Key: key}
	a.Customers = &customer.Client{B: backends.API, Key: key}
	a.Discounts = &discount.Client{B: backends.API, Key: key}
	a.Disputes = &dispute.Client{B: backends.API, Key: key}
	a.EphemeralKeys = &ephemeralkey.Client{B: backends.API, Key: key}
	a.ExchangeRates = &exchangerate.Client{B: backends.API, Key: key}
	a.Events = &event.Client{B: backends.API, Key: key}
	a.Fees = &fee.Client{B: backends.API, Key: key}
	a.FeeRefunds = &feerefund.Client{B: backends.API, Key: key}
	a.Files = &file.Client{B: backends.Uploads, Key: key}
	a.FileLinks = &filelink.Client{B: backends.API, Key: key}
	a.Invoices = &invoice.Client{B: backends.API, Key: key}
	a.InvoiceItems = &invoiceitem.Client{B: backends.API, Key: key}
	a.IssuerFraudRecords = &issuerfraudrecord.Client{B: backends.API, Key: key}
	a.IssuingAuthorizations = &authorization.Client{B: backends.API, Key: key}
	a.IssuingCardholders = &cardholder.Client{B: backends.API, Key: key}
	a.IssuingCards = &issuingcard.Client{B: backends.API, Key: key}
	a.IssuingDisputes = &issuingdispute.Client{B: backends.API, Key: key}
	a.IssuingTransactions = &transaction.Client{B: backends.API, Key: key}
	a.LoginLinks = &loginlink.Client{B: backends.API, Key: key}
	a.Orders = &order.Client{B: backends.API, Key: key}
	a.OrderReturns = &orderreturn.Client{B: backends.API, Key: key}
	a.PaymentIntents = &paymentintent.Client{B: backends.API, Key: key}
	a.PaymentSource = &paymentsource.Client{B: backends.API, Key: key}
	a.Payouts = &payout.Client{B: backends.API, Key: key}
	a.Plans = &plan.Client{B: backends.API, Key: key}
	a.Products = &product.Client{B: backends.API, Key: key}
	a.RadarValueLists = &valuelist.Client{B: backends.API, Key: key}
	a.RadarValueListItems = &valuelistitem.Client{B: backends.API, Key: key}
	a.Recipients = &recipient.Client{B: backends.API, Key: key}
	a.Refunds = &refund.Client{B: backends.API, Key: key}
	a.ReportRuns = &reportrun.Client{B: backends.API, Key: key}
	a.ReportTypes = &reporttype.Client{B: backends.API, Key: key}
	a.Reversals = &reversal.Client{B: backends.API, Key: key}
	a.SigmaScheduledQueryRuns = &scheduledqueryrun.Client{B: backends.API, Key: key}
	a.Skus = &sku.Client{B: backends.API, Key: key}
	a.Sources = &source.Client{B: backends.API, Key: key}
	a.SourceTransactions = &sourcetransaction.Client{B: backends.API, Key: key}
	a.Subscriptions = &sub.Client{B: backends.API, Key: key}
	a.SubscriptionItems = &subitem.Client{B: backends.API, Key: key}
	a.TerminalConnectionTokens = &terminalconnectiontoken.Client{B: backends.API, Key: key}
	a.TerminalLocations = &terminallocation.Client{B: backends.API, Key: key}
	a.TerminalReaders = &terminalreader.Client{B: backends.API, Key: key}
	a.Tokens = &token.Client{B: backends.API, Key: key}
	a.Topups = &topup.Client{B: backends.API, Key: key}
	a.Transfers = &transfer.Client{B: backends.API, Key: key}
	a.UsageRecords = &usagerecord.Client{B: backends.API, Key: key}
	a.WebhookEndpoints = &webhookendpoint.Client{B: backends.API, Key: key}
}

// New creates a new Stripe client with the appropriate secret key
// as well as providing the ability to override the backends as needed.
func New(key string, backends *stripe.Backends) *API {
	api := API{}
	api.Init(key, backends)
	return &api
}
