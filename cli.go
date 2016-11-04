package main

import (
  "github.com/urfave/cli"
  "fmt"
  "os"
)

func main() {
  app := cli.NewApp()
  app.Name = "pomogoro"
  app.Usage = "Track those tasks and FINISH THEM"
  app.Action = func(c *cli.Context) error {
    fmt.Println("Tracking....")
    return nil
  }

  app.Run(os.Args)
}
