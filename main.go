package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	version = "0.1"
)

func (d *DatabaseConnection) interactive() {
	// Continuous loop to take user input
	var i string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("mssqlcli > ")
		// Get user input
		scanner.Scan()
		i = scanner.Text()

		// Check commands
		l := strings.ToLower(i)
		if l == "currentuser" {
			d.GetCurrentUser()
		} else if l == "exit" {
			break
		} else if l == "help" {
			h := printOptions()
			fmt.Printf(h)
			fmt.Printf("%s\t\t:\t%s\n", "exit", "Exit the program")
		} else if l == "isadmin" {
			d.IsSysadmin()
		} else if l == "systemuser" {
			d.GetSystemUser()
		} else if l == "version" {
			d.GetVersion()
		} else {
			// Run the command provided
			q := strings.TrimSuffix(i, "\n")
			d.RawQuery(q)
		}
	}
}

func printOptions() string {
	var r string
	a1 := fmt.Sprintf("%s\t:\t%s\n", "currentuser", "Get the current user")
	a2 := fmt.Sprintf("%s\t\t:\t%s\n", "isadmin", "Check if you are running as a sysadmin")
	a3 := fmt.Sprintf("%s\t:\t%s\n", "systemuser", "Get the system user")
	a4 := fmt.Sprintf("%s\t\t:\t%s\n", "version", "Get the version of the database server")

	r = a1 + a2 + a3 + a4

	return r
}

func printDefaults(s string) {
	fmt.Printf("mssqlcli v%s\n\n", version)
	flag.PrintDefaults()
	fmt.Println("\n" + s)
	os.Exit(0)
}

func main() {
	var db DatabaseConnection

	interact := fmt.Sprintf("%s\t:\t%s\n", "interact", "Interactive mode")
	check := fmt.Sprintf("%s\t\t:\t%s\n", "check", "Only check if the credentials work")

	flag.StringVar(&db.database, "database", "master", "Database to connect to")
	flag.BoolVar(&db.debug, "debug", false, "Print information such as the SQL commands being executed")
	flag.StringVar(&db.host, "host", "", "Host to connect to")
	flag.StringVar(&db.password, "password", "", "Password")
	flag.IntVar(&db.port, "port", 1433, "Port to connect to")
	flag.StringVar(&db.user, "user", "", "User to connect with")
	method := flag.String("method", "interact", "Run a specific method\n"+check+interact+printOptions())
	flag.Parse()

	// Check args
	if db.host == "" {
		printDefaults("You must provide a server to connect to")
	} else if db.user == "" {
		printDefaults("You must provide a username to connect with")
	} else if db.password == "" {
		printDefaults("You must provide the user's password")
	}

	// Connect to the database
	err := db.Connect()
	if err != nil {
		log.Fatalln(fmt.Errorf("error connecting: %v", err))
	}

	// Defer closing when we are done
	defer db.Close()

	m := strings.ToLower(*method)
	if m == "check" {
		// do nothing since we already tried connecting
	} else if m == "currentuser" {
		fmt.Printf("[+] Curent user: ")
		db.GetCurrentUser()
	} else if m == "interact" {
		db.interactive()
	} else if m == "isadmin" {
		fmt.Printf("[+] Is sysadmin: ")
		db.IsSysadmin()
	} else if m == "systemuser" {
		fmt.Printf("[+] System user: ")
		db.GetCurrentUser()
	} else if m == "version" {
		fmt.Printf("[+] Version: ")
		db.GetVersion()
	} else {
		printDefaults("Invalid method provided")
	}
}
