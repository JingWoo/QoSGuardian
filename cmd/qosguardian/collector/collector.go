// Copyright (c) Huawei Technologies Co., Ltd. 2022. All rights reserved.
// rubik licensed under the Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//     http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v2 for more details.
// Author: wujing
// Create: 2022-05-18
// Description: define the collect subcommand of qosguardian

package collector

import (
		"fmt"

		"github.com/spf13/cobra"
)

func validateParams() error {
	return nil
}

var CollectorCmd = &cobra.Command {
		Use: "collector",
		Aliases: []string{"col"},
		Short: "An agent for collecting, analyzing container performance metrics",
		Long: `The Cloud Profiler Collector collects container performance metrics,
analyzes the distribution rules of metrics, extracts features, and analyzes
the correlation between metrics, and reversely optimize the collection strategy`,
		// silence usage when an error occurs.
		/*
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
				// fs := cmd.Flags()
				if err := validateParams(); err != nil {
						return err
				}
				return nil
		},
		*/
		Args: func(cmd *cobra.Command, args []string) error {
				for _, arg := range args {
						if len(arg) > 0 {
								return fmt.Errorf("%q are illegal arguments for %q", args, cmd.CommandPath())
						}
				}
				return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Hello World!")
		},
}

// func init() {
	// collectorCmd.Flags().BoolVarP(&onlyDigits, "", "d", false, "Count only digits")
	// root.RootCmd.AddCommand(collectorCmd)
// }
