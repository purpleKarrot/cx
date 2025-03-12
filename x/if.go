// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package x

func If[T any](c bool, t, f T) T {
	if c {
		return t
	}
	return f
}
