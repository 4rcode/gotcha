# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

test:
	go test -cover -coverprofile=coverage.out -count 1 ./...

coverage:
	go tool cover -html=coverage.out
