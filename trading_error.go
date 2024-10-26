package exception

type TradingError interface {
	ErrorCode() string
	Error() string
	OriginalError() error
	SetOriginalError(err error)
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
	if b.originalError != nil {
		return b.originalError
	}
	return nil
}

func (b tradingErrorImpl) SetOriginalError(err error) {
	if err != nil {
		b.originalError = err
	}
}

func NewTradingError(errorCode string, message string) TradingError {
	return &tradingErrorImpl{
		errorCode: errorCode,
		message:   message,
	}
}
