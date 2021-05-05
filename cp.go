package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/progrium/hostctl/providers"
	"github.com/spf13/cobra"
)

func init() {
	Hostctl.AddCommand(cpCmd)
}

var cpCmd = &cobra.Command{
	Use:   "cp <name> <src> [<dst>]",
	Short: "Copy files to host",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 && defaultName == "" {
			cmd.Usage()
			os.Exit(1)
		}
		name := args[0]
		src := args[1]
		dst := "."
		if len(args) > 2 {
			dst = args[2]
		}
		provider, err := providers.Get(providerName, true)
		fatal(err)
		host := provider.Get(namespace + name)
		if host == nil {
			os.Exit(1)
		}
		fatal(scpExec(host.IP, src, dst))
	},
}

func scpExec(ip, src, dst string) error {
	binary, err := exec.LookPath("scp")
	if err != nil {
		return fmt.Errorf("Unable to find scp")
	}
	args := []string{"scp", src, fmt.Sprintf("%s@%s:%s", sshUser, ip, dst)}
	return syscall.Exec(binary, args, os.Environ())
}
