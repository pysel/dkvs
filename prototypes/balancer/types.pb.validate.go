// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: dkvs/balancer/types.proto

package balancer

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

// Validate checks the field values on Coverage with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Coverage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Coverage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CoverageMultiError, or nil
// if none found.
func (m *Coverage) ValidateAll() error {
	return m.validate(true)
}

func (m *Coverage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTicks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CoverageValidationError{
						field:  fmt.Sprintf("Ticks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CoverageValidationError{
						field:  fmt.Sprintf("Ticks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CoverageValidationError{
					field:  fmt.Sprintf("Ticks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CoverageMultiError(errors)
	}

	return nil
}

// CoverageMultiError is an error wrapping multiple validation errors returned
// by Coverage.ValidateAll() if the designated constraints aren't met.
type CoverageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CoverageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CoverageMultiError) AllErrors() []error { return m }

// CoverageValidationError is the validation error returned by
// Coverage.Validate if the designated constraints aren't met.
type CoverageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CoverageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CoverageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CoverageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CoverageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CoverageValidationError) ErrorName() string { return "CoverageValidationError" }

// Error satisfies the builtin error interface
func (e CoverageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCoverage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CoverageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CoverageValidationError{}

// Validate checks the field values on Tick with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Tick) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tick with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TickMultiError, or nil if none found.
func (m *Tick) ValidateAll() error {
	return m.validate(true)
}

func (m *Tick) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	// no validation rules for Covers

	if len(errors) > 0 {
		return TickMultiError(errors)
	}

	return nil
}

// TickMultiError is an error wrapping multiple validation errors returned by
// Tick.ValidateAll() if the designated constraints aren't met.
type TickMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TickMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TickMultiError) AllErrors() []error { return m }

// TickValidationError is the validation error returned by Tick.Validate if the
// designated constraints aren't met.
type TickValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TickValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TickValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TickValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TickValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TickValidationError) ErrorName() string { return "TickValidationError" }

// Error satisfies the builtin error interface
func (e TickValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTick.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TickValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TickValidationError{}
