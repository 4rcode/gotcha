// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gooroo

import (
	"errors"
	"fmt"
	"path"
)

const traceFormat = "\n  at %s:%d"

func newError(
	err error, file string, line int,
) error {
	return &Error{
		err,
		fmt.Sprintf(
			traceFormat,
			path.Base(file),
			line)}
}

// Error TODO
type Error struct {
	internal error
	trace    string
}

// As TODO
func (e *Error) As(target interface{}) bool {
	return errors.As(e.internal, target)
}

// Error TODO
func (e *Error) Error() string {
	return e.internal.Error() + e.trace
}

// Is TODO
func (e *Error) Is(target error) bool {
	return errors.Is(e.internal, target)
}

// Unwrap TODO
func (e *Error) Unwrap() error {
	return e.internal
}
