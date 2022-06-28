package main

import (
	"fmt"
)

// Check TDE status on all databases
func (d *DatabaseConnection) CheckTDE() {
	sqlStatement := "SELECT name, database_id, is_master_key_encrypted_by_server, is_encrypted FROM master.sys.databases;"
	d.RawQuery(sqlStatement)
}

// Disable xp_cmdshell stored procedure
func (d *DatabaseConnection) DisableXPCmdShell() {
	sqlStatement := "exec sp_configure 'xp_cmdshell', 0 ;RECONFIGURE;exec sp_configure 'show advanced options', 0 ;RECONFIGURE;"
	d.RawQuery(sqlStatement)
}

// Enable xp_cmdshell stored procedure
func (d *DatabaseConnection) EnableXPCmdShell() {
	sqlStatement := "exec sp_configure 'show advanced options',1;RECONFIGURE;exec sp_configure 'xp_cmdshell', 1;RECONFIGURE;"
	d.RawQuery(sqlStatement)
}

// Get the current user
func (d *DatabaseConnection) GetCurrentUser() {
	sqlStatement := "SELECT CURRENT_USER;"
	d.RawQuery(sqlStatement)
}

// Get the system user
func (d *DatabaseConnection) GetSystemUser() {
	sqlStatement := "SELECT SYSTEM_USER;"
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

// List linked servers
func (d *DatabaseConnection) LinkedServers() {
	sqlStatement := "exec sp_linkedservers;"
	d.RawQuery(sqlStatement)
}

// List all databases
func (d *DatabaseConnection) ListDatabases() {
	sqlStatement := "SELECT name FROM master.dbo.sysdatabases;"
	d.RawQuery(sqlStatement)
}

// List users you can impersonate
func (d *DatabaseConnection) ListImpersonations() {
	sqlStatement := "SELECT distinct b.name FROM sys.server_permissions a INNER JOIN sys.server_principals b ON a.grantor_principal_id = b.principal_id WHERE a.permission_name = 'IMPERSONATE';"
	d.RawQuery(sqlStatement)
}

// List all traces on the system
func (d *DatabaseConnection) ListTraces() {
	sqlStatement := "SELECT * FROM master.sys.traces;"
	d.RawQuery(sqlStatement)
}

// List all users
func (d *DatabaseConnection) ListUsers() {
	sqlStatement := "SELECT name FROM master..syslogins;"
	d.RawQuery(sqlStatement)
}

// Run a raw SQL query
func (d *DatabaseConnection) RawQuery(q string) {

	if d.debug {
		fmt.Printf("Executing: %s\n", q)
	}

	stmt, err := d.connection.Prepare(q)
	if err != nil {
		fmt.Printf("Query err: %v\n", err)
		return
	}
	defer stmt.Close()

	data, err := stmt.Query()
	if err != nil {
		fmt.Printf("Query err: %v\n", err)
		return
	}
	defer data.Close()

	colNames, err := data.Columns()
	if err != nil {
		fmt.Printf("Colmn name err: %v", err)
		return
	}

	if len(colNames) > 1 {
		for _, c := range colNames {
			fmt.Printf("%s\t", c)
		}
		fmt.Println()

		cols := make([]interface{}, len(colNames))
		colPtrs := make([]interface{}, len(colNames))
		for i := 0; i < len(colNames); i++ {
			colPtrs[i] = &cols[i]
		}

		var myMap = make(map[string]interface{})

		for data.Next() {
			err := data.Scan(colPtrs...)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}

			for i, col := range cols {
				myMap[colNames[i]] = col
			}

			for _, cname := range colNames {
				r := fmt.Sprintf("%v", myMap[cname])
				fmt.Printf("%s\t", r)
			}
			fmt.Println()
		}
	} else {
		for data.Next() {
			var results string
			nErr := data.Scan(&results)
			if nErr != nil {
				fmt.Printf("Error: %v", nErr)
			}

			fmt.Println(results)
		}
	}
}
