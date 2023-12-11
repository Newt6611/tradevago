package tri

import (
	"os"
	"os/exec"
)

type Side string

const (
    SELL Side = "SELL"
    BUY  Side = "BUY"
)

func ClearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
