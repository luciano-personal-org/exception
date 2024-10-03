package exception

type TradingError interface {
	ErrorCode() string
	Error() string
}

type tradingErrorImpl struct {
	errorCode string
	message   string
}

func (b tradingErrorImpl) ErrorCode() string {
	return b.errorCode
}

func (b tradingErrorImpl) Error() string {
	return b.message
}

func NewTradingError(errorCode string, message string) TradingError {
	return &tradingErrorImpl{
		errorCode: errorCode,
		message:   message,
	}
}
