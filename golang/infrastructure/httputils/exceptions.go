package httputils

//Error mapping constants
const GeneralException int = 0

//Front messages error mapping
var errorMap = map[int]string{
	0: "Ocurrió un error en el servidor",
}

type Exception struct {
	Code         int    `json:"code"`
	Error        error  `json:"-"`
	Message      string `json:"message"`
	FrontMessage string `json:"front_message"`
}

func NewException(code int, err error) *Exception {
	return &Exception{
		Code:         code,
		Error:        err,
		Message:      err.Error(),
		FrontMessage: errorMap[code],
	}
}
