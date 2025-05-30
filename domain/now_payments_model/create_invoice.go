package nowpaymentsmodel

import (
	"time"

	nowpaymentlibs "github.com/jSierraB3991/now_payment_api/domain/now_payment_libs"
	"gorm.io/gorm"
)

type NowPaymentCreateInvoice struct {
	gorm.Model
	CreateInvoiceId     uint                               `gorm:"column:id;not null"`
	UserId              uint                               `gorm:"column:user_id;not null"`
	NowPaymentId        string                             `gorm:"column:now_payment_id;not null"`
	OrderID             string                             `gorm:"column:order_id;not null"`
	OrderDescription    string                             `gorm:"column:order_description;not null"`
	PriceAmount         string                             `gorm:"column:price_amount;not null"`
	PriceCurrency       string                             `gorm:"column:price_currency;not null"`
	PayCurrency         string                             `gorm:"column:pay_currency;not null"`
	IpnCallbackURL      string                             `gorm:"column:ipn_callback_url;not null"`
	InvoiceURL          string                             `gorm:"column:invoice_url;not null"`
	SuccessURL          string                             `gorm:"column:success_url"`
	CancelURL           string                             `gorm:"column:cancel_url"`
	CreatedAtMy         time.Time                          `gorm:"column:create_at_my;not null"`
	UpdatedAtMy         time.Time                          `gorm:"column:update_at_my;not null"`
	Status              nowpaymentlibs.CreateInvoiceStatus `gorm:"column:status_pay;not null;default:'CREATE_INVOICE_STATUS_WAIT_PAY'"`
	NowPaymentPaymentId uint                               `gorm:"column:now_payment_payment_id"`
}
