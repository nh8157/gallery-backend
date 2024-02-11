package response

import "encoding/json"

const (
	JsonParseError  = "Unable to parse JSON."
	DynamoDbError   = "Internal server error."
	S3Error         = "Internal server error"
	FileDuplication = "File already exists"
)

type SuccessResponse struct {
	Msg string `json:"msg"`
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}

func ToString(errorResponse *ErrorResponse) string {
	jsonStr, err := json.Marshal(*errorResponse)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}
