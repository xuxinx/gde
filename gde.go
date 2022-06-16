package gde

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Setenv(file string) {
	bs, err := exec.Command("bash", "-c", fmt.Sprintf("source %s; echo '<<<ENVIRONMENT>>>'; env", file)).CombinedOutput()
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(bytes.NewReader(bs))
	start := false
	for s.Scan() {
		if s.Text() == "<<<ENVIRONMENT>>>" {
			start = true
			continue
		}
		if start {
			kv := strings.SplitN(s.Text(), "=", 2)
			if len(kv) == 2 {
				os.Setenv(kv[0], kv[1])
			}
		}
	}
}
