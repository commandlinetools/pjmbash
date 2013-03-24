// Copyright 2013 Michael A. Wright. All rights reserved.
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE.txt file.
package list

import (
	"log"
)

func List() []string {
	log.Printf("List")
	list := make([]string, 0)
	return list
}
