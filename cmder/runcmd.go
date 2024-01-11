package cmder

import (
	"fmt"
	"os/exec"

	lua "github.com/yuin/gopher-lua"
)

func RunCommand(command string) (string, error) {
	cmd := exec.Command(command)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("err: failed to execute command: %s\n%w", command, err)
	}
	return string(output), nil
}

func RunLuaScript(strScript string) {
	// 创建Lua状态机
	state := lua.NewState()
	defer state.Close()

	err := state.DoString(strScript)
	if err != nil {
		panic(err)
	}
}
