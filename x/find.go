// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package x

func FindIf[T any](slice []T, predicate func(*T) bool) *T {
	for i := range slice {
		r := &slice[i]
		if predicate(r) {
			return r
		}
	}
	return nil
}
