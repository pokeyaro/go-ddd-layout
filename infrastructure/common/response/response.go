package response

const (
	OkMsg  = "success"
	OkCode = 0

	FailMsg  = "failed"
	FailCode = 50000
)

type BaseResponse[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code   int    `json:"code"`
	Err    string `json:"error,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// Ok returns a BaseResponse with a success code and message.
func Ok() *BaseResponse[any] {
	return &BaseResponse[any]{
		Code: OkCode,
		Msg:  OkMsg,
		Data: (*any)(nil),
	}
}

// Fail returns an ErrorResponse with a failure code, error message, and optional detail message.
func Fail(err error, vs ...string) *ErrorResponse {
	var detail string
	if len(vs) == 0 {
		detail = FailMsg
	} else {
		detail = vs[0]
	}

	return &ErrorResponse{
		Code:   FailCode,
		Err:    err.Error(),
		Detail: detail,
	}
}

// Data returns a BaseResponse with a success code, message, and data.
func Data(data any) *BaseResponse[any] {
	return &BaseResponse[any]{
		Code: OkCode,
		Msg:  OkMsg,
		Data: data,
	}
}

// List returns a BaseResponse with a success code, message, and a map containing the given data list and its count.
func List[T any](data []T) *BaseResponse[any] {
	return &BaseResponse[any]{
		Code: OkCode,
		Msg:  OkMsg,
		Data: map[string]any{
			"items": data,
			"count": len(data),
		},
	}
}
