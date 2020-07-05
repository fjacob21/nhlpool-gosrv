package data

// Result Is the result of a request
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SUCCESS Result is a success
const SUCCESS = 0

// ERROR General error
const ERROR = -1

//ACCESSDENIED Player do not have the required access right
const ACCESSDENIED = -2

// NOTFOUND Result not found
const NOTFOUND = -3

// EXISTS Result not found
const EXISTS = -4
