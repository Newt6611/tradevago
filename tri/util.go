package tri

import (
	"os"
	"os/exec"

	"github.com/Newt6611/tradevago/pkg/api"
	"github.com/Newt6611/tradevago/pkg/notify"
)

type Backend struct {
    Api *api.Api
    Apiws *api.WSApi
    MsgBot notify.Notifier
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
