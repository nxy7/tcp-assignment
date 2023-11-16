package serve

import (
	"fmt"

	"github.com/nxy7/tcp-assignment/internal/hub"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		h := hub.NewHub()
		port := ":22222"
		fmt.Println("Starting to serve traffic on port ", port)
		err := h.Serve(port)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	ServeCmd.Flags().StringP("port", "p", ":22222", "")
}
