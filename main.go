package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App { // Using cli package for making this cli tool
    Name: "HealthChecker", // Name of the tool
    Usage: "Tool to check website status", // Usage/Use-case of the tool
    Flags: []cli.Flag { // Flags required from the user
      &cli.StringFlag {
        Name: "domain", // Name of the flag
        Usage: "Domain name to check,", 
        Required: true, // Always true 
      },
      &cli.StringFlag {
        Name: "port", // the port number to connect to
        Aliases: []string{"p"},
        Usage: "Port number to check.", 
        Required: false, // Not a neccessary input from user side
      },
    },
    Action: func(c *cli.Context) error {  
      port := c.String("port") 
      if c.String("port") == "" { // In case of which the user does not input the port
        port = "80" // Default port set to 80
      }
      status := Check(c.String("domain"), port) // Calling the Check function
      fmt.Println(status) // Printing the return values from the function call
      return nil // if there are no errors the action returns nil error
    },
  }
  err := app.Run(os.Args) // executing the tool
  if err!=nil {
    log.Fatal(err) // in case there is any error it gets printed
  }
}
