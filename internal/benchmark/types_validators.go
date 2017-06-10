// Code generated by gencheck
// DO NOT EDIT!

package benchmark

import (
	"errors"
	"net"
	"strings"
	"time"

	"github.com/abice/gencheck"
)

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s SingleString) Validate() error {

	var vErrors gencheck.ValidationErrors

	// BEGIN Entry Validations
	// required
	if s.Entry == "" {
		return append(vErrors, gencheck.NewFieldError("SingleString", "Entry", "required", errors.New("is required")))
	}
	// END Entry Validations

	return nil
}

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s TestAll) Validate() error {

	var vErrors gencheck.ValidationErrors

	// BEGIN Required Validations
	// required
	if s.Required == "" {
		return append(vErrors, gencheck.NewFieldError("TestAll", "Required", "required", errors.New("is required")))
	}
	// END Required Validations

	// BEGIN Len Validations
	// len
	if !(len(s.Len) == 10) {
		return append(vErrors, gencheck.NewFieldError("TestAll", "Len", "len", errors.New("length mismatch")))
	}
	// END Len Validations

	// BEGIN Min Validations
	// min
	if len(s.Min) < 5 {
		return append(vErrors, gencheck.NewFieldError("TestAll", "Min", "min", errors.New("length failed check for min=5")))
	}
	// END Min Validations

	// BEGIN Max Validations
	// max
	if len(s.Max) > 100 {
		return append(vErrors, gencheck.NewFieldError("TestAll", "Max", "max", errors.New("length failed check for max=100")))
	}
	// END Max Validations

	// BEGIN CIDR Validations
	// required
	if s.CIDR == "" {
		return append(vErrors, gencheck.NewFieldError("TestAll", "CIDR", "required", errors.New("is required")))
	}

	// cidr
	_, _, CIDRerr := net.ParseCIDR(s.CIDR)
	if CIDRerr != nil {
		return append(vErrors, gencheck.NewFieldError("TestAll", "CIDR", "cidr", CIDRerr))
	}
	// END CIDR Validations

	// BEGIN LteTime Validations
	// lte
	tLteTime := time.Now().UTC()
	if s.LteTime.After(tLteTime) {
		return append(vErrors, gencheck.NewFieldError("TestAll", "LteTime", "lte", errors.New("is after now")))
	}
	// END LteTime Validations

	// BEGIN GteTime Validations
	// gte
	tGteTime := time.Now().UTC()
	if s.GteTime.Before(tGteTime) {
		return append(vErrors, gencheck.NewFieldError("TestAll", "GteTime", "gte", errors.New("is before now")))
	}
	// END GteTime Validations

	// BEGIN Gte Validations
	// gte
	if s.Gte < 1.2345 {
		return append(vErrors, gencheck.NewFieldError("TestAll", "Gte", "gte", errors.New("failed check for gte=1.2345")))
	}
	// END Gte Validations

	// BEGIN NotNil Validations
	// required
	if s.NotNil == nil {
		return append(vErrors, gencheck.NewFieldError("TestAll", "NotNil", "required", errors.New("is required")))
	}
	// END NotNil Validations

	// BEGIN Contains Validations
	// contains
	if !strings.Contains(s.Contains, "fox") {
		return append(vErrors, gencheck.NewFieldError("TestAll", "Contains", "contains", errors.New("Contains did not contain fox")))
	}
	// END Contains Validations

	// BEGIN Hex Validations
	// hex
	if err := gencheck.IsHex(&s.Hex); err != nil {
		return append(vErrors, gencheck.NewFieldError("TestAll", "Hex", "hex", err))
	}
	// END Hex Validations

	// BEGIN UUID Validations
	// uuid
	if err := gencheck.IsUUID(&s.UUID); err != nil {
		return append(vErrors, gencheck.NewFieldError("TestAll", "UUID", "uuid", err))
	}
	// END UUID Validations

	// BEGIN MinInt Validations
	// min
	if s.MinInt < 12345 {
		return append(vErrors, gencheck.NewFieldError("TestAll", "MinInt", "min", errors.New("failed check for min=12345")))
	}
	// END MinInt Validations

	// BEGIN MaxInt Validations
	// max
	if s.MaxInt > 12345 {
		return append(vErrors, gencheck.NewFieldError("TestAll", "MaxInt", "max", errors.New("failed check for max=12345")))
	}
	// END MaxInt Validations

	// BEGIN Dive Validations
	// required
	if s.Dive == nil {
		return append(vErrors, gencheck.NewFieldError("TestAll", "Dive", "required", errors.New("is required")))
	}

	// dive
	if s.Dive != nil {
		if err := gencheck.Validate(s.Dive); err != nil {
			return append(vErrors, gencheck.NewFieldError("TestAll", "Dive", "dive", err))
		}
	}
	// END Dive Validations

	return nil
}

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s TestContainsAny) Validate() error {

	var vErrors gencheck.ValidationErrors

	// BEGIN Any Validations
	// containsany
	if !strings.ContainsAny(s.Any, "@#!") {
		return append(vErrors, gencheck.NewFieldError("TestContainsAny", "Any", "containsany", errors.New("Any did not contain any of @#!")))
	}
	// END Any Validations

	return nil
}

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s TestDive) Validate() error {

	var vErrors gencheck.ValidationErrors

	// BEGIN Value Validations
	// required
	if s.Value == nil {
		return append(vErrors, gencheck.NewFieldError("TestDive", "Value", "required", errors.New("is required")))
	}

	// dive
	if s.Value != nil {
		if err := gencheck.Validate(s.Value); err != nil {
			return append(vErrors, gencheck.NewFieldError("TestDive", "Value", "dive", err))
		}
	}
	// END Value Validations

	return nil
}

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s TestHex) Validate() error {

	var vErrors gencheck.ValidationErrors

	// BEGIN Value Validations
	// hex
	if err := gencheck.IsHex(&s.Value); err != nil {
		return append(vErrors, gencheck.NewFieldError("TestHex", "Value", "hex", err))
	}
	// END Value Validations

	return nil
}

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s TestMap) Validate() error {

	var vErrors gencheck.ValidationErrors

	// BEGIN Value Validations
	// contains
	if _, foundValue := s.Value["test"]; !foundValue {
		return append(vErrors, gencheck.NewFieldError("TestMap", "Value", "contains", errors.New("Value did not contain test")))
	}
	// END Value Validations

	return nil
}

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s TestString) Validate() error {

	var vErrors gencheck.ValidationErrors

	// BEGIN Required Validations
	// required
	if s.Required == "" {
		return append(vErrors, gencheck.NewFieldError("TestString", "Required", "required", errors.New("is required")))
	}
	// END Required Validations

	// BEGIN Len Validations
	// len
	if !(len(s.Len) == 10) {
		return append(vErrors, gencheck.NewFieldError("TestString", "Len", "len", errors.New("length mismatch")))
	}
	// END Len Validations

	// BEGIN Min Validations
	// min
	if len(s.Min) < 5 {
		return append(vErrors, gencheck.NewFieldError("TestString", "Min", "min", errors.New("length failed check for min=5")))
	}
	// END Min Validations

	// BEGIN Max Validations
	// max
	if len(s.Max) > 100 {
		return append(vErrors, gencheck.NewFieldError("TestString", "Max", "max", errors.New("length failed check for max=100")))
	}
	// END Max Validations

	return nil
}

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s TestUUID) Validate() error {

	var vErrors gencheck.ValidationErrors

	// BEGIN UUID Validations
	// required
	if s.UUID == "" {
		return append(vErrors, gencheck.NewFieldError("TestUUID", "UUID", "required", errors.New("is required")))
	}

	// uuid
	if err := gencheck.IsUUID(&s.UUID); err != nil {
		return append(vErrors, gencheck.NewFieldError("TestUUID", "UUID", "uuid", err))
	}
	// END UUID Validations

	return nil
}
