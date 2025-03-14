// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package x

func Map[T, R any](slice []T, fn func(T) R) []R {
	r := make([]R, len(slice))
	for i, t := range slice {
		r[i] = fn(t)
	}
	return r
}
