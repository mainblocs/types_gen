// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: remove_staff.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RemoveStaff with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RemoveStaff) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveStaff with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RemoveStaffMultiError, or
// nil if none found.
func (m *RemoveStaff) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveStaff) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for RemoveStaffEmail

	// no validation rules for BusinessId

	if len(errors) > 0 {
		return RemoveStaffMultiError(errors)
	}

	return nil
}

// RemoveStaffMultiError is an error wrapping multiple validation errors
// returned by RemoveStaff.ValidateAll() if the designated constraints aren't met.
type RemoveStaffMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveStaffMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveStaffMultiError) AllErrors() []error { return m }

// RemoveStaffValidationError is the validation error returned by
// RemoveStaff.Validate if the designated constraints aren't met.
type RemoveStaffValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveStaffValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveStaffValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveStaffValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveStaffValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveStaffValidationError) ErrorName() string { return "RemoveStaffValidationError" }

// Error satisfies the builtin error interface
func (e RemoveStaffValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveStaff.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveStaffValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveStaffValidationError{}