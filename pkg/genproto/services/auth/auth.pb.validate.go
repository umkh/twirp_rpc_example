// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: services/auth/auth.proto

package pb_auth

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

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReqMultiError, or nil
// if none found.
func (m *LoginReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetEmail()); l < 5 || l > 20 {
		err := LoginReqValidationError{
			field:  "Email",
			reason: "value length must be between 5 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 5 || l > 20 {
		err := LoginReqValidationError{
			field:  "Password",
			reason: "value length must be between 5 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginReqMultiError(errors)
	}
	return nil
}

// LoginReqMultiError is an error wrapping multiple validation errors returned
// by LoginReq.ValidateAll() if the designated constraints aren't met.
type LoginReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReqMultiError) AllErrors() []error { return m }

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

// Validate checks the field values on LoginRes with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRes with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginResMultiError, or nil
// if none found.
func (m *LoginRes) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	// no validation rules for UpdateToken

	if len(errors) > 0 {
		return LoginResMultiError(errors)
	}
	return nil
}

// LoginResMultiError is an error wrapping multiple validation errors returned
// by LoginRes.ValidateAll() if the designated constraints aren't met.
type LoginResMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResMultiError) AllErrors() []error { return m }

// LoginResValidationError is the validation error returned by
// LoginRes.Validate if the designated constraints aren't met.
type LoginResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResValidationError) ErrorName() string { return "LoginResValidationError" }

// Error satisfies the builtin error interface
func (e LoginResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResValidationError{}

// Validate checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterReqMultiError, or
// nil if none found.
func (m *RegisterReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FullName

	// no validation rules for Email

	// no validation rules for Phone

	// no validation rules for Sex

	// no validation rules for Password

	if len(errors) > 0 {
		return RegisterReqMultiError(errors)
	}
	return nil
}

// RegisterReqMultiError is an error wrapping multiple validation errors
// returned by RegisterReq.ValidateAll() if the designated constraints aren't met.
type RegisterReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReqMultiError) AllErrors() []error { return m }

// RegisterReqValidationError is the validation error returned by
// RegisterReq.Validate if the designated constraints aren't met.
type RegisterReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReqValidationError) ErrorName() string { return "RegisterReqValidationError" }

// Error satisfies the builtin error interface
func (e RegisterReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReqValidationError{}

// Validate checks the field values on ActivateReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ActivateReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActivateReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ActivateReqMultiError, or
// nil if none found.
func (m *ActivateReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ActivateReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SecretKey

	if len(errors) > 0 {
		return ActivateReqMultiError(errors)
	}
	return nil
}

// ActivateReqMultiError is an error wrapping multiple validation errors
// returned by ActivateReq.ValidateAll() if the designated constraints aren't met.
type ActivateReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActivateReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActivateReqMultiError) AllErrors() []error { return m }

// ActivateReqValidationError is the validation error returned by
// ActivateReq.Validate if the designated constraints aren't met.
type ActivateReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActivateReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActivateReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActivateReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActivateReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActivateReqValidationError) ErrorName() string { return "ActivateReqValidationError" }

// Error satisfies the builtin error interface
func (e ActivateReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActivateReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActivateReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActivateReqValidationError{}

// Validate checks the field values on ForgotReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ForgotReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForgotReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ForgotReqMultiError, or nil
// if none found.
func (m *ForgotReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ForgotReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetEmail()); l < 5 || l > 20 {
		err := ForgotReqValidationError{
			field:  "Email",
			reason: "value length must be between 5 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ForgotReqMultiError(errors)
	}
	return nil
}

// ForgotReqMultiError is an error wrapping multiple validation errors returned
// by ForgotReq.ValidateAll() if the designated constraints aren't met.
type ForgotReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForgotReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForgotReqMultiError) AllErrors() []error { return m }

// ForgotReqValidationError is the validation error returned by
// ForgotReq.Validate if the designated constraints aren't met.
type ForgotReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForgotReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForgotReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForgotReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForgotReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForgotReqValidationError) ErrorName() string { return "ForgotReqValidationError" }

// Error satisfies the builtin error interface
func (e ForgotReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForgotReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForgotReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForgotReqValidationError{}
