// SPDX-FileCopyrightText: 2025 Daniel Pfeifer <daniel@pfeifer-mail.de>
// SPDX-License-Identifier: ISC

package x

import (
	"fmt"
	"os"
	"os/exec"
)

func Run(cmd *exec.Cmd, verbose bool) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if verbose {
		fmt.Printf("\nExecuting command: %v\n\n", cmd)
	}

	return cmd.Run()
}
