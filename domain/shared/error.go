package shared

import "fmt"

type DomainError struct {
	Err        error
	ErrorType  string
	StatusCode int
	Details    string
	Message    string
}

func (r *DomainError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}
