package service

import (
	"errors"
	"fmt"

	"github.com/docker/swarm-v2/api"
	"github.com/docker/swarm-v2/cmd/swarmctl/common"
	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:     "remove <service ID>",
		Short:   "Remove a service",
		Aliases: []string{"rm"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("service ID missing")
			}
			c, err := common.Dial(cmd)
			if err != nil {
				return err
			}

			service, err := getService(common.Context(cmd), c, args[0])
			if err != nil {
				return err
			}
			_, err = c.RemoveService(common.Context(cmd), &api.RemoveServiceRequest{ServiceID: service.ID})
			if err != nil {
				return err
			}
			fmt.Println(args[0])
			return nil
		},
	}
)