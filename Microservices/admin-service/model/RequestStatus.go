package model

type RequestStatus string
const (
	SUBMITTED RequestStatus = "submitted"
	APPROVED RequestStatus = "approved"
	DENIED	RequestStatus = "denied"
	)