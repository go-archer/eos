package upgrade

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

const eosCmd = "https://github.com/go-helios/eos.git"

var CmdUpgrade = &cobra.Command{
	Use:     "upgrade",
	Short:   "更新eos命令",
	Long:    "更新eos命令",
	Example: "eos upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("go install %s", eosCmd)
		cmd := exec.Command("go", "install", eosCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("go install %s 错误\n", err)
		}
		fmt.Printf("\n eos 更新成功!\n")
	},
}
