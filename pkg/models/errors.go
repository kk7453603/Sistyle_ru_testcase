package models

type ValidationError struct {
	Message string   `json:"message"`
	Members []string `json:"members"`
}

type ErrorData struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}

type ErrorResponse struct {
	Code             string            `json:"code"`
	Message          string            `json:"message"`
	Details          string            `json:"details"`
	Data             ErrorData         `json:"data"`
	ValidationErrors []ValidationError `json:"validationErrors"`
}

type APIError struct {
	Error ErrorResponse `json:"error"`
}
