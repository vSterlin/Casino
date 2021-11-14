package payment

type PaymentController struct {
	ps *PaymentService
}

func NewPaymentController(ps *PaymentService) *PaymentController {
	return &PaymentController{ps: ps}
}
