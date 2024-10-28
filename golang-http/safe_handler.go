package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"function/function"
)

type ErrorResponse struct {
	Err ErrorBody `json:"error"`
}

type ErrorBody struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

func SafeHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			_, _, lineno := identifyPanic()

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			resp := ErrorResponse{
				Err: ErrorBody{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("function panicked at line %d: %v", lineno, r),
					Details: map[string]interface{}{
						"panic_message": fmt.Sprintf("%v", r),
						"panic_line":    lineno,
					},
				},
			}
			if bytes, err := json.Marshal(resp); err == nil {
				w.Write(bytes)
			}
		}
	}()

	w.Header().Add("X-Invoked", "true")

	function.Handle(w, r)
}

func identifyPanic() (filename, functionName string, lineno int) {
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		filename, lineno = fn.FileLine(pc)
		functionName = fn.Name()
		if !strings.HasPrefix(functionName, "runtime.") {
			break
		}
	}

	return
}
