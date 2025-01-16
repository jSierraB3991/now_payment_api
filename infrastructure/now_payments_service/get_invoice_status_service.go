package nowpaymentsservice

func (s *NowPaymentService) GetInvoiceStatus(invoiceId string) error {
	var result interface{}
	err := s.httpClient.Get(s.apiUrl, ""+invoiceId, s.apiKey, &result)
	if err != nil {
		return err
	}
	return nil
}
