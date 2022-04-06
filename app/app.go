package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func GeneratorCli() (app *cli.App) {
	app = cli.NewApp()

	app.Name = "Cli of Servers"
	app.Usage = "Searching ips, servers, and more..."
	app.Description = "Best CLI to fetch public IP'S and Server Names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search ip's by hostname",
			Flags:  flags,
			Action: searchByIPs,
		},
		{
			Name:   "server",
			Usage:  "Search names of servers by hostname",
			Flags:  flags,
			Action: searchByServers,
		},
	}

	return
}

func searchByIPs(c *cli.Context) {
	host := c.String("host")

	fmt.Println("Searching...")
	IPs, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d IP's\n", len(IPs))

	for _, ip := range IPs {
		fmt.Println(ip)
	}

}

func searchByServers(c *cli.Context) {
	host := c.String("host")

	fmt.Println("Searching...")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d Names Servers\n", len(servers))

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
