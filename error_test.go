// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gooroo

import (
	"errors"
	"testing"

	"github.com/4rcode/gorgon"
)

func TestTracedError(t *testing.T) {
	assert := gorgon.New(t)

	var err error = &mockedError{}

	wrapped := err

	WrapWith(&wrapped, "w3")
	WrapWith(&wrapped, "w2")
	WrapWith(&wrapped, "w1")

	assert.Error(err, wrapped)
	assert.Equal(
		"w1 w2 w3 mock"+
			"\n  at traced_error_test.go:21"+
			"\n  at traced_error_test.go:22"+
			"\n  at traced_error_test.go:23",
		wrapped.Error())

	var mockedErr *mockedError

	assert.
		That(
			errors.As(wrapped, &mockedErr)).
		Log("the wrapped error should be a mockedError")

	assert.Equal(err, mockedErr)

	unwrapped := wrapped

	for i := 0; i < 6; i++ {
		unwrapped = errors.Unwrap(unwrapped)
	}

	assert.Equal(err, unwrapped)
}
