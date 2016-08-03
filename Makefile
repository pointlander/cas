# Copyright 2016 The CAS Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

HAS_BINDATA := $(shell go-bindata -version 2>/dev/null)

all: algebrite.bundle-for-browser.js
ifndef HAS_BINDATA
	go get github.com/jteeuwen/go-bindata/...
endif
	go-bindata algebrite.bundle-for-browser.js
	go build

algebrite.bundle-for-browser.js:
	wget https://raw.githubusercontent.com/davidedc/Algebrite/master/dist/algebrite.bundle-for-browser.js

clean:
	rm -f *.js
	rm -f cas
