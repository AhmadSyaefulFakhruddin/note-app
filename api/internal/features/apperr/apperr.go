package apperr

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    // The HTTP Status Code (e.g., 400, 404, 409)
	Message string // The safe message sent to the React/Frontend client
	Err     error  // The raw, ugly internal error (for your server logs only)
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf(`%s: %s`, e.Message, e.Err.Error())
	}

	return e.Message
}

// --- HELPER FUNCTIONS ---
// These make writing your Repositories incredibly fast and clean

func NewConflict(message string, err error) *AppError {
	return &AppError{
		Code:    http.StatusConflict, // 409
		Message: message,
		Err:     err,
	}
}

func NewInternal(err error) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "An unexpected system error occurred. Please try again later.", // Safe for Toast!
		Err:     err,                                                            // Save the ugly error for the logs
	}
}

// NewBadRequest is used when the client sends invalid data.
// We allow a custom message here so the frontend knows exactly which field is wrong.
func NewBadRequest(message string, err error) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest, // HTTP 400
		Message: message,               // e.g., "Title is required" or "Invalid JSON format"
		Err:     err,                   // The raw binding error for our backend logs
	}
}
