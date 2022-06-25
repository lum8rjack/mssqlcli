package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

type DatabaseConnection struct {
	connection *sql.DB
	database   string
	debug      bool
	dbContext  context.Context
	host       string
	password   string
	port       int
	user       string
}

// Connect to the database
func (d *DatabaseConnection) Connect() error {
	// Make sure the connection details were set
	if d.host == "" {
		return errors.New("host not provided")
	} else if d.password == "" {
		return errors.New("password not provided")
	} else if d.user == "" {
		return errors.New("user not provided")
	} else if d.database == "" {
		return errors.New("database not provided")
	}

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", d.host, d.user, d.password, d.port, d.database)
	d.dbContext = context.Background()

	var connectionError error
	d.connection, connectionError = sql.Open("mssql", connectionString)
	if connectionError != nil {
		return connectionError
	}

	err := d.connection.PingContext(d.dbContext)
	if err != nil {
		e := fmt.Sprintf("Error checking db connection: %v", err)
		return errors.New(e)
	}

	fmt.Printf("[+] Successfully connect to %s:%d (database: %s) as '%s'\n", d.host, d.port, d.database, d.user)
	return nil
}

// Close the database connection
func (d *DatabaseConnection) Close() {
	d.connection.Close()
}
