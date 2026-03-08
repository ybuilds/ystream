package controllers

import "net/http"

const (
	OK  = http.StatusOK
	ISE = http.StatusInternalServerError
	UN  = http.StatusUnauthorized
	NF  = http.StatusNotFound
	BR  = http.StatusBadRequest
	CR  = http.StatusCreated
)
