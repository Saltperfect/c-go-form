package controllers

import (
	"fmt"
	"net/http"
)

type ActionHandler struct{}

func NewActionHandler() *ActionHandler {
	return new(ActionHandler)
}

func (ah *ActionHandler) View(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter into view handler")
	w.Header().Add("hello", "hello")
}
