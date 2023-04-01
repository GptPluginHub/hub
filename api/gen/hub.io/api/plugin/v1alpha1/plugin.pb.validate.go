// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: hub.io/api/plugin/v1alpha1/plugin.proto

package v1alpha1

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

	types "hub.io/api/types"
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

	_ = types.OrderBy(0)
)

// Validate checks the field values on Plugin with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Plugin) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Plugin with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PluginMultiError, or nil if none found.
func (m *Plugin) ValidateAll() error {
	return m.validate(true)
}

func (m *Plugin) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Domain

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for AuthType

	// no validation rules for LogoUrl

	// no validation rules for ContactEmail

	// no validation rules for Organization

	// no validation rules for ApiType

	// no validation rules for ApiUrl

	for idx, item := range m.GetLabel() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PluginValidationError{
						field:  fmt.Sprintf("Label[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PluginValidationError{
						field:  fmt.Sprintf("Label[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PluginValidationError{
					field:  fmt.Sprintf("Label[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for State

	// no validation rules for InstallNum

	// no validation rules for Score

	// no validation rules for Heat

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return PluginMultiError(errors)
	}

	return nil
}

// PluginMultiError is an error wrapping multiple validation errors returned by
// Plugin.ValidateAll() if the designated constraints aren't met.
type PluginMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PluginMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PluginMultiError) AllErrors() []error { return m }

// PluginValidationError is the validation error returned by Plugin.Validate if
// the designated constraints aren't met.
type PluginValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PluginValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PluginValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PluginValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PluginValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PluginValidationError) ErrorName() string { return "PluginValidationError" }

// Error satisfies the builtin error interface
func (e PluginValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlugin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PluginValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PluginValidationError{}

// Validate checks the field values on PluginLabel with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PluginLabel) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PluginLabel with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PluginLabelMultiError, or
// nil if none found.
func (m *PluginLabel) ValidateAll() error {
	return m.validate(true)
}

func (m *PluginLabel) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	if len(errors) > 0 {
		return PluginLabelMultiError(errors)
	}

	return nil
}

// PluginLabelMultiError is an error wrapping multiple validation errors
// returned by PluginLabel.ValidateAll() if the designated constraints aren't met.
type PluginLabelMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PluginLabelMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PluginLabelMultiError) AllErrors() []error { return m }

// PluginLabelValidationError is the validation error returned by
// PluginLabel.Validate if the designated constraints aren't met.
type PluginLabelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PluginLabelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PluginLabelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PluginLabelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PluginLabelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PluginLabelValidationError) ErrorName() string { return "PluginLabelValidationError" }

// Error satisfies the builtin error interface
func (e PluginLabelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPluginLabel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PluginLabelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PluginLabelValidationError{}

// Validate checks the field values on ListPluginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListPluginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPluginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPluginRequestMultiError, or nil if none found.
func (m *ListPluginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPluginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	// no validation rules for OrderBy

	// no validation rules for SortBy

	// no validation rules for FuzzyName

	if len(errors) > 0 {
		return ListPluginRequestMultiError(errors)
	}

	return nil
}

// ListPluginRequestMultiError is an error wrapping multiple validation errors
// returned by ListPluginRequest.ValidateAll() if the designated constraints
// aren't met.
type ListPluginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPluginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPluginRequestMultiError) AllErrors() []error { return m }

// ListPluginRequestValidationError is the validation error returned by
// ListPluginRequest.Validate if the designated constraints aren't met.
type ListPluginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPluginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPluginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPluginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPluginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPluginRequestValidationError) ErrorName() string {
	return "ListPluginRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPluginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPluginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPluginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPluginRequestValidationError{}

// Validate checks the field values on ListPluginResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPluginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPluginResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPluginResponseMultiError, or nil if none found.
func (m *ListPluginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPluginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListPluginResponseValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListPluginResponseValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListPluginResponseValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetItem() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListPluginResponseValidationError{
						field:  fmt.Sprintf("Item[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListPluginResponseValidationError{
						field:  fmt.Sprintf("Item[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPluginResponseValidationError{
					field:  fmt.Sprintf("Item[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListPluginResponseMultiError(errors)
	}

	return nil
}

// ListPluginResponseMultiError is an error wrapping multiple validation errors
// returned by ListPluginResponse.ValidateAll() if the designated constraints
// aren't met.
type ListPluginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPluginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPluginResponseMultiError) AllErrors() []error { return m }

// ListPluginResponseValidationError is the validation error returned by
// ListPluginResponse.Validate if the designated constraints aren't met.
type ListPluginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPluginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPluginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPluginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPluginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPluginResponseValidationError) ErrorName() string {
	return "ListPluginResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListPluginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPluginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPluginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPluginResponseValidationError{}