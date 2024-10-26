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

func (b tradingErrorImpl) Details() string {
	if b.details != "" {
		return b.details
	}
	return ""
}

func (b tradingErrorImpl) ErrorCode() string {
	if b.errorCode != "" {
		return b.errorCode
	}
	return ""
}

func (b tradingErrorImpl) Error() string {
	if b.message != "" {
		return b.message
	}
	return ""
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

func (b tradingErrorImpl) SetDetails(details string) {
	if details != "" {
		b.details = details
	}
}

func NewTradingError(errorCode string, message string) TradingError {
	return &tradingErrorImpl{
		errorCode: errorCode,
		message:   message,
	}
}
