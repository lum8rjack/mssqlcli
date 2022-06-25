# mssqlcli

## Overview
Program to help enumerate and exploit MSSQL servers. This tools was made as an alternative to Impacket's mssqlclient.py program so the user does not have to fight with Python dependencies.

```bash
./mssqlcli
mssqlcli v0.1

  -database string
        Database to connect to (default "master")
  -debug
        Print information such as the SQL commands being executed
  -host string
        Host to connect to
  -method string
        Run a specific method
        check           :       Only check if the credentials work
        interact        :       Interactive mode
        currentuser     :       Get the current user
        isadmin         :       Check if you are running as a sysadmin
        systemuser      :       Get the system user
        version         :       Get the version of the database server
         (default "interact")
  -password string
        Password
  -port int
        Port to connect to (default 1433)
  -user string
        User to connect with

You must provide a server to connect to
```

## Requirements
The program relies on the following modules:
- github.com/microsoft/go-mssqldb

## Setup
Make sure go is installed and then run the following:
```bash
git clone https://github.com/lum8rjack/mssqlcli
cd mssqlcli
make
```

A docker-compose file is also provided if you would like to spin up a simple MSSQL server to test with.

## Examples

Getting the current database version
```bash
./mssqlcli -user sa -password "Password123" -host 127.0.0.1 -method version
[+] Successfully connect to 127.0.0.1:1433 (database: master) as 'sa'
[+] Version: Microsoft SQL Server 2019 (RTM-CU16-GDR) (KB5014353) - 15.0.4236.7 (X64) 
        May 29 2022 15:55:47 
        Copyright (C) 2019 Microsoft Corporation
        Developer Edition (64-bit) on Linux (Ubuntu 20.04.4 LTS) <X64>
```

Running in interactive mode to execute the built-in commands and raw SQL queries.
```bash
./mssqlcli -user sa -password "Password123" -host 127.0.0.1
[+] Successfully connect to 127.0.0.1:1433 (database: master) as 'sa'
mssqlcli > help
currentuser     :       Get the current user
isadmin         :       Check if you are running as a sysadmin
systemuser      :       Get the system user
version         :       Get the version of the database server
exit            :       Exit the program
mssqlcli > currentuser
dbo
mssqlcli > isadmin
1
mssqlcli > systemuser
sa
mssqlcli > SELECT name FROM sys.databases;
master
tempdb
model
msdb
mssqlcli > exit
```

## References / Credit

Using Go to connect and query databases:
- https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go?view=azuresql
- https://blog.logrocket.com/using-sql-database-golang/

Run MSSQL in a Docker container:
- https://blog.logrocket.com/docker-sql-server/

Impacket toolset that contains their mssqlclient.py program:
- https://github.com/SecureAuthCorp/impacket

