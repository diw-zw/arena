// Copyright 2026 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package llm

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

const rbgBinary = "arena-rbg"

// NewLLMCommand returns a cobra command that proxies all args to `arena-rbg llm ...`.
func NewLLMCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "llm",
		Short:              "LLM deployment management commands (powered by arena-rbg)",
		Long:               `Commands for managing LLM model deployments. Requires 'arena-rbg' binary to be installed.`,
		DisableFlagParsing: true,
		SilenceUsage:       true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRBG(args)
		},
	}
	return cmd
}

// runRBG delegates execution to `arena-rbg llm <args...>`.
func runRBG(args []string) error {
	rbgPath, err := exec.LookPath(rbgBinary)
	if err != nil {
		return fmt.Errorf("'%s' binary not found in PATH, please install it first: %w", rbgBinary, err)
	}

	subArgs := append([]string{"llm"}, args...)
	c := exec.Command(rbgPath, subArgs...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
