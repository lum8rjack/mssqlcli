package main

import (
	"fmt"
)

// Get the current user
func (d *DatabaseConnection) GetCurrentUser() {
	sqlStatement := "SELECT CURRENT_USER;"
	d.RawQuery(sqlStatement)
}

// Get the version of the MSSQL database
func (d *DatabaseConnection) GetVersion() {
	sqlStatement := "SELECT @@version;"
	d.RawQuery(sqlStatement)
}

// Check if you are running as a sysadmin
func (d *DatabaseConnection) IsSysadmin() {
	sqlStatement := "SELECT IS_SRVROLEMEMBER('sysadmin');"
	d.RawQuery(sqlStatement)
}

// Run a raw SQL query
func (d *DatabaseConnection) RawQuery(q string) {

	if d.debug {
		fmt.Printf("Executing: %s\n", q)
	}

	stmt, err := d.connection.Prepare(q)
	if err != nil {
		fmt.Printf("Query err: %v", err)
		return
	}

	data, queryErr := stmt.Query()
	if queryErr != nil {
		fmt.Printf("Query err: %v", queryErr)
		return
	}

	for data.Next() {
		var result string

		nErr := data.Scan(&result)
		if nErr != nil {
			fmt.Printf("Error: %v", nErr)
		}

		fmt.Println(result)
	}
}

// Get the system user
func (d *DatabaseConnection) GetSystemUser() {
	sqlStatement := "SELECT SYSTEM_USER;"
	d.RawQuery(sqlStatement)
}
