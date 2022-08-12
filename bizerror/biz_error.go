package bizerror

import (
	"encoding/json"
	"errors"
	"go-start/bizcode"
)

type BizError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	err     error
}

func (biz BizError) Error() string {
	return biz.err.Error()
}

func Wrap(err error) BizError {
	return BizError{bizcode.ServerError.Code(), err.Error(), err}
}

func Biz(code bizcode.BizCode) BizError {
	return BizError{code.Code(), code.String(), errors.New(code.String())}
}

func Of(code int, message string) BizError {
	return BizError{code, message, errors.New(message)}
}

func Parse(a any) (*BizError, error) {
	jsonStr, jsonErr := json.Marshal(a)
	if jsonErr != nil {
		return nil, jsonErr
	}

	bizError := BizError{}
	jsonErr = json.Unmarshal(jsonStr, &bizError)
	if jsonErr != nil {
		return nil, jsonErr
	}

	bizError.err = errors.New(bizError.Message)
	return &bizError, nil
}

func (biz BizError) UnWrap() error {
	return biz.err
}
