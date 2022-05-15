package main

import (
	"github.com/fi-GarO/CTC-2022/05/cmd/client"
	"github.com/fi-GarO/CTC-2022/05/cmd/server"
	"github.com/fi-GarO/CTC-2022/05/pkg/util"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "05",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cmd.AddCommand(server.Cmd(), client.Cmd())

	util.ExitOnError(cmd.Execute())
}
