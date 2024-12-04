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
	originalError error
}

func (b tradingErrorImpl) Details() string {
	return b.details
}

func (b tradingErrorImpl) ErrorCode() string {
	return b.errorCode
}

func (b tradingErrorImpl) Error() string {
	return b.originalError.Error()
}

func (b tradingErrorImpl) OriginalError() error {
	return b.originalError
}

func (b tradingErrorImpl) SetOriginalError(originalError error) {
	b.originalError = originalError
}

func NewTradingError(errorCode string, details string) TradingError {
	return &tradingErrorImpl{
		errorCode: errorCode,
		details:   details,
	}
}
