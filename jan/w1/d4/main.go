package main

import (
	"strings"
	"w1/d4/middleware"
)

func LoggingMiddleware(next middleware.Handler) middleware.Handler {
	return func(s string) string {
		println("Input:", s)
		result := next(s)
		println("Output:", result)
		return result
	}
}

// UppercaseMiddleware converts the result to uppercase
func UppercaseMiddleware(next middleware.Handler) middleware.Handler {
	return func(s string) string {
		println("UppercaseMiddleware: Before calling next")
		result := next(s)
		println("UppercaseMiddleware: After calling next, uppercasing result")
		return strings.ToUpper(result)
	}
}

func TrimMiddleware(next middleware.Handler) middleware.Handler {
	return func(s string) string {
		println("TrimMiddleware: Trimming input")
		trimmed := strings.TrimSpace(s)
		println("TrimMiddleware: Calling next with trimmed input")
		return next(trimmed)
	}
}

func main() {
	handler := middleware.Chain(LoggingMiddleware, UppercaseMiddleware, TrimMiddleware)
	handler("hello world")
}
