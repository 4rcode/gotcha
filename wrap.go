// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gooroo

import (
	"fmt"
	"runtime"
)

const wrapper = " %w"

func isErrNil(err *error) bool {
	return err == nil || *err == nil
}

func wrap(err *error) {
	_, file, line, _ := runtime.Caller(2)
	*err = newError(*err, file, line)
}

// Wrap TODO
func Wrap(err *error) bool {
	if isErrNil(err) {
		return false
	}

	wrap(err)

	return true
}

// Wrapf returns a new decorator function that wraps the provided error in a new
// one, with the specified message formatted with the optional arguments.
func Wrapf(
	err *error, format string, args ...interface{},
) bool {
	if isErrNil(err) {
		return false
	}

	*err =
		fmt.Errorf(
			format+wrapper,
			append(args, *err)...)

	wrap(err)

	return true
}
