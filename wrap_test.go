// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gooroo

import (
	"errors"
	"testing"

	"github.com/4rcode/gorgon"
)

func TestWrap(t *testing.T) {
	assert := gorgon.New(t)

	var err error

	assert.Equal(false, Wrap(nil))
	assert.Equal(false, Wrap(&err))

	err = errors.New("error")
	wrap := err

	assert.Equal(true, Wrap(&wrap))

	assert.Error(err, wrap)
	assert.Equal(
		"error\n  at wrap_test.go:25",
		wrap.Error())
}

func TestWrapf(t *testing.T) {
	assert := gorgon.New(t)

	var err error

	assert.Equal(false, Wrapf(nil, ""))
	assert.Equal(false, Wrapf(&err, ""))

	err = errors.New("error")
	wrapped := err

	assert.Equal(
		true,
		Wrapf(&wrapped, "wrapped %s:", "error"))
	assert.Error(err, wrapped)
	assert.Equal(
		"wrapped error: error\n  at wrap_test.go:46",
		wrapped.Error())
}
