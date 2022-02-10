package credential

// JsTicketHandle js ticket获取
type JsTicketHandle interface {
	// GetTicket 获取ticket
	GetTicket(accessToken string, optional ...interface{}) (ticket string, err error)
}
