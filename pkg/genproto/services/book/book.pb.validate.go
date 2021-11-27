// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: services/book/book.proto

package pb_book

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

// Validate checks the field values on CreateReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateReqMultiError, or nil
// if none found.
func (m *CreateReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Author

	if len(errors) > 0 {
		return CreateReqMultiError(errors)
	}
	return nil
}

// CreateReqMultiError is an error wrapping multiple validation errors returned
// by CreateReq.ValidateAll() if the designated constraints aren't met.
type CreateReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateReqMultiError) AllErrors() []error { return m }

// CreateReqValidationError is the validation error returned by
// CreateReq.Validate if the designated constraints aren't met.
type CreateReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateReqValidationError) ErrorName() string { return "CreateReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateReqValidationError{}

// Validate checks the field values on GetReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in GetReqMultiError, or nil if none found.
func (m *GetReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetReqMultiError(errors)
	}
	return nil
}

// GetReqMultiError is an error wrapping multiple validation errors returned by
// GetReq.ValidateAll() if the designated constraints aren't met.
type GetReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetReqMultiError) AllErrors() []error { return m }

// GetReqValidationError is the validation error returned by GetReq.Validate if
// the designated constraints aren't met.
type GetReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetReqValidationError) ErrorName() string { return "GetReqValidationError" }

// Error satisfies the builtin error interface
func (e GetReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetReqValidationError{}

// Validate checks the field values on Book with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Book) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Book with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BookMultiError, or nil if none found.
func (m *Book) ValidateAll() error {
	return m.validate(true)
}

func (m *Book) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Author

	if len(errors) > 0 {
		return BookMultiError(errors)
	}
	return nil
}

// BookMultiError is an error wrapping multiple validation errors returned by
// Book.ValidateAll() if the designated constraints aren't met.
type BookMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookMultiError) AllErrors() []error { return m }

// BookValidationError is the validation error returned by Book.Validate if the
// designated constraints aren't met.
type BookValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookValidationError) ErrorName() string { return "BookValidationError" }

// Error satisfies the builtin error interface
func (e BookValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBook.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookValidationError{}