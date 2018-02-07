# Copyright 2016 The CAS Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

HAS_FILEB0X := $(shell which fileb0x)

all: algebrite.bundle-for-browser.js
ifndef HAS_FILEB0X
	go get -u github.com/UnnoTed/fileb0x
endif
	fileb0x fileb0x.json
	go build

algebrite.bundle-for-browser.js:
	curl https://raw.githubusercontent.com/davidedc/Algebrite/master/dist/algebrite.bundle-for-browser.js > algebrite.bundle-for-browser.js

clean:
	rm -f *.js
	rm -f cas
