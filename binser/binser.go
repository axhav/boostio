// Copyright 2018 The go-boostio Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binser provides types to read and write binary archives from the C++
// Boost Serialization library.
package binser // import "github.com/go-boostio/boostio/binser"

//go:generate go run ./testdata/gen-binary-archive.go

import (
	"github.com/pkg/errors"
)

var (
	ErrNotBoost      = errors.New("binser: not a Boost binary archive")
	ErrInvalidHeader = errors.New("binser: invalid Boost binary archive header")
)

// Header describes a binary boost archive.
type Header struct {
	Version uint16
	Flags   uint64
}

// TypeDescr describes an on-disk binary boost archive type.
type TypeDescr struct {
	Version uint32
	Flags   uint8
}

// Unmarshaler is the interface implemented by types that can unmarshal a binary
// Boost description of themselves.
type Unmarshaler interface {
	UnmarshalBoost(r *Reader) error
}
