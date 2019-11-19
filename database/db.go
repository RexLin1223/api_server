package database

import (
	"database/sql"
	"fmt"

	// Driver of SQL
	_ "github.com/go-sql-driver/mysql"
)

// ConnectionInfo is detail info of remote database
type ConnectionInfo struct {
	address  string
	port     string
	username string
	password string
	dbName   string
}

// NewConnectionInfo is factor pattern of ConnectionInfo
func NewConnectionInfo(address string, port string, username string, password string, dbName string) *ConnectionInfo {
	return &ConnectionInfo{
		address:  address,
		port:     port,
		username: username,
		password: password,
		dbName:   dbName,
	}
}

// ISQLConnector is interface of database connector which handles all manipulation from user
type ISQLConnector interface {
	Open() error
	Close() error

	// Create(table string, params map[string]interface{})
	Create(string, map[string]interface{}) error

	//  Read(table string, constraint string)
	Read(string, string) (map[string]interface{}, error)

	// Update function will send select command for find target data out from particular {table}
	Update(string, string, interface{}) error

	// Delete function will send delete command for remove some target matches {constraint}
	Delete(string, string) error
}

// MySQLConnectorImpl is implmentation of IConnector
type MySQLConnectorImpl struct {
	info *ConnectionInfo
	db   *sql.DB
}

// NewMySQLLConnector is factor pattern of Connector
func NewMySQLLConnector(info *ConnectionInfo) ISQLConnector {
	conn := &MySQLConnectorImpl{
		info: info,
		db:   nil,
	}
	return conn
}

// Open connection with remote database
func (conn *MySQLConnectorImpl) Open() error {
	command := conn.info.username + ":" + conn.info.password + "@tcp(" + conn.info.address + ":" + conn.info.port + ")/" + conn.info.dbName
	db, err := sql.Open("mysql", command)
	if err != nil {
		return err
	}

	conn.db = db
	return nil
}

// Close connection with remote database
func (conn *MySQLConnectorImpl) Close() error {

	if conn.db != nil {
		err := conn.db.Close()
		return err
	}
	return nil
}

// Create will insert data into particular table()
func (conn *MySQLConnectorImpl) Create(table string, params map[string]interface{}) error {
	if conn.db != nil {
		err := conn.db.Ping()
		if err != nil {
			return err
		}
	} else {
		err := conn.Open()
		if err != nil {
			return err
		}
	}

	// Format  create command

	return nil
}

// Read function will send select command for find target data matches {constraint} from particular {table}
func (conn *MySQLConnectorImpl) Read(table string, constraint string) (map[string]interface{}, error) {
	command := "SELTECT * FROM " + conn.info.dbName + "." + table
	rows, err := conn.db.Query(command)
	if err != nil {
		return nil, err
	}
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return nil, nil
}

// Update function will send select command for find target data out from particular {table}
func (conn *MySQLConnectorImpl) Update(table string, constraint string, newValue interface{}) error {

	return nil
}

// Delete function will send delete command for remove some target matches {constraint}
func (conn *MySQLConnectorImpl) Delete(table string, constraint string) error {

	return nil
}
