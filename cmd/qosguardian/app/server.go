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
// Description: define the app command of qosguardian

package app

import (
	"fmt"

	"github.com/spf13/cobra"
	log "k8s.io/klog/v2"

	"isula.org/qosguardian/cmd/qosguardian/app/options"
	"isula.org/qosguardian/pkg/config"
	"isula.org/qosguardian/pkg/utils"
)

func NewQoSGuardianCommand() *cobra.Command {
	server := options.NewQoSGuardianServer()

	cmd := &cobra.Command{
		Use:   "qosguardian",
		Short: "qosguardian is used for Interference Analysis of Co-Located Container Workloads",
		Long: `By analyzing the collected metrics of online workload, and modeling the performance
interference accord to the co-located environment, detecting or predicting the performance
of online workloads under the conditions of dynamic load, shared resource competition, and
interference mode, further find the source of interference, guiding cluster resource planning,
scheduling strategy optimization, and rational allocation of nodes in offline resources, and
maximize the throughput of offline services under the premise of ensuring the real-time and
stability of online services.`,
		Version:            "0.0.1",
		DisableFlagParsing: false,
		SilenceUsage:       true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(server.Config)
			if err := run(server); err != nil {
				return fmt.Errorf("failed to run QoSGuardian Server")
			}
			return nil
		},
	}

	server.AddFlags(cmd.Flags())
	return cmd
}

func run(server *options.QoSGuardianServer) error {
	signalCh := utils.SetupSignalHandler()
	go func() {
		log.Info("QoS Guardian Server Config: ", server.Config)
		qoSGuardianConfig, err := config.ParseQoSGuardianConfig(server.Config)
		if err != nil {
			log.Error("failed to parse server config: ", server.Config)
		}
		log.Info("QoSGuardianConfig", qoSGuardianConfig)
	}()
	<-signalCh
	return nil
}
