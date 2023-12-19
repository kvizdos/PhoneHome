package server

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

func sendNotification(callerName string) {
	err := beeep.Notify("📞 Call Received", fmt.Sprintf("%s Phoned Home!", callerName), "assets/information.png")
	if err != nil {
		panic(err)
	}
}
