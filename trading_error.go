package exception

type TradingError interface {
	ErrorCode() string
	Error() string
	OriginalError() error
}

type tradingErrorImpl struct {
	errorCode     string
	message       string
	originalError error
}

func (b tradingErrorImpl) ErrorCode() string {
	return b.errorCode
}

func (b tradingErrorImpl) Error() string {
	return b.message
}

func (b tradingErrorImpl) OriginalError() error {
	return b.originalError
}

func NewTradingError(errorCode string, message string, originalError error) TradingError {
	return &tradingErrorImpl{
		errorCode:     errorCode,
		message:       message,
		originalError: originalError,
	}
}
