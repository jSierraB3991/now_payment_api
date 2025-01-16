# NOW Payment API

A Golang package for integrating with the NOW Payments API, providing functionality for creating and managing payments and invoices.

## Installation

```bash
go get github.com/jSierraB3991/now_payment_api
```

## Configuration

To use the NOW Payment API service, you'll need to initialize it with your credentials and configuration:

```go
import nowPaymentService "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_service"

// Initialize the service
service := nowPaymentService.InitiateService(
    database,                // Your database instance
    "YOUR_API_KEY",         // NOW Payments API key
    "https://api.nowpayments.io/v1", // API base URL (or sandbox URL for testing)
    "YOUR_USERNAME",        // NOW Payments username
    "YOUR_PASSWORD"         // NOW Payments password
)
```

For sandbox testing, use: `https://api-sandbox.nowpayments.io/v1`

## Usage

### Service Implementation

The package provides an interface for interacting with the NOW Payments API. Here's how to implement it in your service:

```go
import nowpaymentsserviceinterface "github.com/jSierraB3991/now_payment_api/domain/now_payments_service_interface"

type YourService struct {
    nowPaymentService nowpaymentsserviceinterface.NowPaymentServiceInterface
    frontendUrl      string
    backendUrl       string
}

func NewYourService(
    nowPaymentService nowpaymentsserviceinterface.NowPaymentServiceInterface,
    frontendUrl string,
    backendUrl string,
) *YourService {
    return &YourService{
        nowPaymentService: nowPaymentService,
        frontendUrl:      frontendUrl,
        backendUrl:       backendUrl,
    }
}
```

### Creating an Invoice

To create a payment invoice:

```go
import "github.com/jSierraB3991/now_payment_api/domain/now_payments_request"

func (s *YourService) CreateInvoice(amountPrice float64, user User) (*nowpaymentsresponse.CreateInvoiceResponse, error) {
    request := nowpaymentsrequest.CreateInvoiceRequest{
        AmountPrice:      amountPrice,
        PriceCurrency:    "USD",
        OrderID:          orderID,
        OrderDescription: fmt.Sprintf("Payment of %v USD for %s", amountPrice, user.Email),
        SuccessURL:       s.frontendUrl + "/success",
        CancelURL:        s.frontendUrl + "/fail",
        IpnCallbackURL:   s.backendUrl + "/public/now_payments/ipn",
    }

    return s.nowPaymentService.CreateInvoice(request, user.ID)
}
```

## Available Methods

The `NowPaymentServiceInterface` provides the following methods:

- `CreatePayment(req CreatePaymentRequest, userId uint) (*CreatePaymentResponse, error)`
- `CreateInvoice(req CreateInvoiceRequest, userId uint) (*CreateInvoiceResponse, error)`
- `GetInvoiceDataService(page, limit int, isAscendente bool, userId uint) ([]NowPaymentCreateInvoice, error)`
- `GetPaymentStatus(paymentId string) (*GetPaymentStatusResponse, error)`
- `GetInvoiceStatus(invoiceId string) (*GetPaymentStatusResponse, error)`

## Error Handling

All methods return an error as the second return value. Always check for errors when calling these methods:

```go
response, err := service.CreateInvoice(request, userID)
if err != nil {
    // Handle error appropriately
    return nil, err
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Resumen de la licencia:

- Puedes copiar, distribuir y modificar este proyecto siempre que mantengas la misma licencia.
- Cualquier modificación realizada y distribuida debe estar bajo los mismos términos.
- No hay garantías sobre el uso del software.

Puedes encontrar una copia completa de la licencia en el archivo [LICENSE](./LICENSE) en el repositorio de este proyecto.
