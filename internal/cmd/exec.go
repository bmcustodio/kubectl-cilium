// Copyright 2020 bmcustodio
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/cache"

	"github.com/bmcustodio/kubectl-cilium/internal/constants"
	ciliumutils "github.com/bmcustodio/kubectl-cilium/internal/utils/cilium"
	nodeutils "github.com/bmcustodio/kubectl-cilium/internal/utils/kubernetes"
)

func init() {
	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Args:                  cobra.MinimumNArgs(1),
	DisableFlagsInUseLine: true,
	Short:                 "Execute a command in a particular Cilium agent (default: '/bin/bash').",
	Use:                   "exec (NODE|NAMESPACE/NAME) [COMMAND [args...]]",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			command []string
		)
		switch len(args) {
		case 1:
			command = []string{constants.DefaultCommand}
		default:
			command = args[1:]
		}
		return exec(args[0], command...)
	},
}

func exec(target string, command ...string) error {
	// Start by attempting to discover the namespace in which Cilium is installed.
	ns, err := ciliumutils.DiscoverCiliumNamespace(kubeClient)
	if err != nil {
		return err
	}
	// Try to understand if 'target' is the name of a node or a '<namespace>/<name>' key targeting a pod.
	tns, tn, err := cache.SplitMetaNamespaceKey(target)
	if err != nil {
		return fmt.Errorf("failed to parse %q as a target: %v", target, err)
	}
	var (
		nn string
	)
	switch {
	case tns == "":
		// Assume that 'target' is the name of a node.
		nn = tn
	default:
		// Lookup the name of the node where the pod referenced by 'target' is running.
		v, err := nodeutils.GetNodeNameFromPod(kubeClient, tns, tn)
		if err != nil {
			return err
		}
		nn = v
	}
	// Double-check whether the targeted node exists.
	ne, err := nodeutils.NodeExists(kubeClient, nn)
	if err != nil {
		return err
	}
	if !ne {
		return fmt.Errorf("node with name %q does not exist", nn)
	}
	// Try to get the name of the Cilium pod running in the targeted node.
	pn, err := ciliumutils.DiscoverCiliumPodInNode(kubeClient, ns, nn)
	if err != nil {
		return err
	}
	return nodeutils.Exec(kubeClient, kubeConfig, streams, ns, pn, true, true, command...)
}
