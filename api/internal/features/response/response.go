package response

import (
	"errors"
	"log/slog"
	"net/http"
	"note-app-api/internal/features/apperr"
)

type Meta struct {
	TotalItems   int `json:"totalItems"`
	TotalPages   int `json:"totalPages"`
	CurrentPage  int `json:"currentPage"`
	ItemsPerPage int `json:"itemsPerPage"`
}

type ApiResponse[T any] struct {
	Status  string `json:"status"`         // Always "success" or "error"
	Message string `json:"message"`        // For toast notifications
	Data    T      `json:"data,omitempty"` // The actual payload
	Meta    *Meta  `json:"meta,omitempty"` // Pointer so it gets hidden if it's nil
}

func SuccessResponse[T any](data T, message string) ApiResponse[T] {
	return ApiResponse[T]{
		Status:  "success",
		Data:    data,
		Message: message,
	}
}

func SuccessWithMeta[T any](data T, message string, meta Meta) ApiResponse[T] {
	return ApiResponse[T]{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    &meta,
	}
}

func ErrorResponse[T any](message string) ApiResponse[T] {
	return ApiResponse[T]{
		Status:  "error",
		Message: message,
	}
}

func Error(err error) (int, ApiResponse[any]) {
	var appErr *apperr.AppError

	if errors.As(err, &appErr) {
		// 1. LOG THE UGLY ERROR TO THE TERMINAL FOR THE BACKEND DEVELOPER
		if appErr.Code >= 500 {
			// For 500 errors, we want a red ERROR log
			slog.Error("Internal System Failure", "details", appErr.Err)
		} else {
			// For 400/409 errors (like duplicate tags), it's just a user mistake, so a WARN log is fine
			slog.Warn("Client Request Failed", "details", appErr.Err)
		}

		// 2. SEND THE SAFE MESSAGE TO THE FRONTEND USER
		return appErr.Code, ApiResponse[any]{
			Status:  "error",
			Message: appErr.Message, // This is the safe string!
		}
	}

	// Fallback for completely unknown errors
	slog.Error("Unknown Server Error", "details", err)
	return http.StatusInternalServerError, ApiResponse[any]{
		Status:  "error",
		Message: "An unexpected internal server error occurred.",
	}
}
