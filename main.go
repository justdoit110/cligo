package main

import (
  "fmt"
  "github.com/urfave/cli/v2"
  "os"
)

func main() {
  app := cli.App{
    Name: "greet",
    Usage: "say it again",
    Action: mainAction,
  }

  app.Run(os.Args)
}

func  mainAction(_ *cli.Context) error {
  fmt.Println("Allo ha")
  return nil
}
