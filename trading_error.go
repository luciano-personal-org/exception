package exception

type TradingError interface {
	Details() string
	ErrorCode() string
	Error() string
	OriginalError() error
	// SetOriginalError(err error)
	// SetDetails(details string)
}

type tradingErrorImpl struct {
	details       string
	errorCode     string
	message       string
	originalError error
}

func (b tradingErrorImpl) Details() string {
	return b.details
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

func NewTradingError(errorCode string, message string, details string, originalError error) TradingError {
	return &tradingErrorImpl{
		errorCode:     errorCode,
		message:       message,
		details:       details,
		originalError: originalError,
	}
}
