package helper

func FailedResponse(msg string) map[string]any {
	return map[string]any{
		"message": msg,
	}
}

func SuccessResponse(msg string) map[string]any {
	return map[string]any{
		"message": msg,
	}
}

func SuccessWithDataResponse(msg string, data any) map[string]any {
	return map[string]any{
		"message": msg,
		"data":    data,
	}
}

func BadRequest(msg string) map[string]any {
	return map[string]any{
		"message": "Id must be integer",
	}
}
