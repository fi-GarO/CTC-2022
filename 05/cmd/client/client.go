package client

import (
	"context"
	"fmt"
	"github.com/fi-GarO/CTC-2022/05/pkg/api"
	"github.com/fi-GarO/CTC-2022/05/pkg/util"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var endpoint string

func newClient() api.Client {
	conn, err := grpc.Dial(
		endpoint,
		grpc.WithReturnConnectionError(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	util.ExitOnError(err)
	return api.NewGrpcClient(api.NewApiClient(conn))
}

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "client",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	f := cmd.PersistentFlags()

	f.StringVarP(&endpoint, "endpoint", "e", "localhost:8080", "endpoint")

	cmd.AddCommand(getCmd(), putCmd(), delCmd())

	return cmd
}

func getCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "get key",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]

			cli := newClient()
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			res, err := cli.Get(ctx, key)
			util.ExitOnError(err)
			fmt.Println(res)
		},
	}
}

func putCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "put key value",
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			key, val := args[0], args[1]
			cli := newClient()
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			err := cli.Put(ctx, key, val)
			util.ExitOnError(err)
		},
	}
}

func delCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "del key",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]

			cli := newClient()
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			_, err := cli.Delete(ctx, key)
			util.ExitOnError(err)
			fmt.Println("Deleted")
		},
	}
}
