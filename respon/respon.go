package respon

type BaseRespon struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Id      int         `json:"id"`
}
