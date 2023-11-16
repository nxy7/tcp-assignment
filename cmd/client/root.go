package client

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nxy7/tcp-assignment/internal/client"
	"github.com/nxy7/tcp-assignment/internal/command"
	"github.com/spf13/cobra"
)

var ClientCmd = &cobra.Command{
	Use:     "client",
	Aliases: []string{"c"},
	Short:   "Connect to server",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.MakeClient("localhost:22222")
		if err != nil {
			panic(err)
		}
		r := bufio.NewReader(os.Stdin)
		var cd command.Command
		fmt.Println("Going into client loop")
		for {
			fmt.Println("Write command, commands available are: 'list' and 'write'")
			t, err := r.ReadString('\n')
			if err != nil {
				panic(err)
			}
			if t == "list\n" {
				fmt.Println("listing")
				cd = &command.List{}
				c.ExecuteCommand(cd)
			} else if t == "write\n" {
				fmt.Println("writing")
				cd = &command.WriteMessage{Message: "TestMessage"}
				c.ExecuteCommand(cd)
			} else {
				fmt.Println("Unknown command", t)
			}
		}
	},
}

func init() {
	ClientCmd.Flags().StringP("addr", "a", "localhost:22222", "")
}
