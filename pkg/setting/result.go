package setting

type Result[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    T      `json:"data"`
}
