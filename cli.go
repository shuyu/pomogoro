package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "pomogoro"
	app.Usage = "Track those tasks and FINISH THEM"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Tracking....")
		timer := StartTimer(5)
		<-timer.C
		println("Timer is OOOOVVVEERRR")
		return nil
	}

	app.Run(os.Args)
}

func StartTimer(sec time.Duration) *time.Timer {
	return time.NewTimer(time.Second * sec)
}
