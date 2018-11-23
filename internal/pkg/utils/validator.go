// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package utils

// Validator util
type Validator struct {
}

func (v *Validator) in(item string, list []string) bool {
	for _, a := range list {
		if a == item {
			return true
		}
	}
	return false
}