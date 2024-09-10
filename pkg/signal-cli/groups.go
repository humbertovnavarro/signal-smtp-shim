package signalcli

import (
	"encoding/json"
	"os/exec"
)

func Groups() (groups []*Group, err error) {
	cmd := exec.Command("signal-cli", "-o", "json", "listGroups")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	groups = make([]*Group, 0)
	json.Unmarshal(output, &groups)
	return
}
