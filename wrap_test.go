// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package gooroo

import (
	"fmt"
	"testing"

	"github.com/4rcode/gorgon"
)

func TestWrap(t *testing.T) {
	assert := gorgon.New(t)
	err := fmt.Errorf("error")
	wrapped := err

	WrapWith(&wrapped, "wrapped")

	assert.Error(err, wrapped)
	assert.Equal(
		"wrapped error\n  at wrap_test.go:20",
		wrapped.Error())
}
