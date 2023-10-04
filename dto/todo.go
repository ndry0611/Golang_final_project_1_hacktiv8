package dto


type NewTodoRequest struct {
	Title string
	Done  bool
}

type NewTodoResponse struct {
	StatusCode int `json:"status"`
	Message string `json:"message"`
	Result string `json:"result"`
	Data interface{} `json:"data"`
}