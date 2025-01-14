package nowpaymentsmodel

import (
	"time"

	"gorm.io/gorm"
)

type NowPaymentCreatePayment struct {
	gorm.Model
	CreatePaymentId uint `gorm:"column:id:not null"`
	UserId          uint `gorm:"column:user_id;not null"`

	PaymentID              string    `gorm:"column:payment_id;not null"`
	PaymentStatus          string    `gorm:"column:payment_status;not null"`
	PayAddress             string    `gorm:"column:pay_address;not null"`
	PriceAmount            float32   `gorm:"column:price_amount;not null"`
	PriceCurrency          string    `gorm:"column:price_currency;not null"`
	PayAmount              float32   `gorm:"column:pay_amount;not null"`
	PayCurrency            string    `gorm:"column:pay_currency;not null"`
	OrderID                string    `gorm:"column:order_id;not null"`
	OrderDescription       string    `gorm:"column:order_description;not null"`
	IpnCallbackURL         string    `gorm:"column:ipn_callback_url;not null"`
	CreatedAt              time.Time `gorm:"column:created_at;not null"`
	UpdatedAt              time.Time `gorm:"column:updated_at;not null"`
	PurchaseID             string    `gorm:"column:purchase_id;not null"`
	AmountReceived         float32   `gorm:"column:amount_received;not null"`
	PayinExtraID           string    `gorm:"column:payin_extra_id;not null"`
	SmartContract          string    `gorm:"column:smart_contract;not null"`
	Network                string    `gorm:"column:network;not null"`
	NetworkPrecision       int       `gorm:"column:network_precision;not null"`
	TimeLimit              string    `gorm:"column:time_limit;not null"`
	BurningPercent         string    `gorm:"column:burning_percent;not null"`
	ExpirationEstimateDate time.Time `gorm:"column:expiration_estimate_date;not null"`
}
