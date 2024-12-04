package exception

type TradingError interface {
	Details() string
	ErrorCode() string
	Error() string
	OriginalError() error
	SetOriginalError(err error)
	SetDetails(details string)
}

type tradingErrorImpl struct {
	details       string
	errorCode     string
	message       string
	originalError error
}

func (b *tradingErrorImpl) Details() string {
	return b.details
}

func (b *tradingErrorImpl) ErrorCode() string {
	return b.errorCode
}

func (b *tradingErrorImpl) Error() string {
	return b.message
}

func (b *tradingErrorImpl) OriginalError() error {
	return b.originalError
}

func (b *tradingErrorImpl) SetOriginalError(err error) {
	b.originalError = err
}

func (b *tradingErrorImpl) SetDetails(details string) {
	b.details = details
}

func NewTradingError(errorCode string, message string) TradingError {
	return &tradingErrorImpl{
		errorCode: errorCode,
		message:   message,
	}
}
