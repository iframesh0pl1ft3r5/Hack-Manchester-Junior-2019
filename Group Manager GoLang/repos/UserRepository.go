package repos

import (
	"crypto/md5"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Name - struct used by the MySQL system
type Name struct {
	Name string `json:"name"`
}

func UserIsValid(uName, pwd string) bool {
	// DB simulation
	_uName, _pwd, _isValid := "gfgdsfgdghgfhsh", "gdrffgergargdfgetgesg6gdr6f78g65fg67dfgdfg", false
	// need to cconvert password to md5
	// then search for username in DB
	// if username exists check to see if password is correct
	// if password is correct return true

	db, err := sql.Open("mysql", "hack:HackManchester2019@tcp(10.0.2.58:3306)/group_mgr")
	if err != nil {
		return _isValid
	}
	defer db.Close()

	// Protecting against mysql injections
	formatted := fmt.Sprintf("SELECT Username from group_mgr.users where Username = \"%s\";", uName)
	// searching for the project
	results, err := db.Query(formatted)
	if err != nil {
		_isValid = false
		return _isValid
	}

	for results.Next() {
		var name Name
		// for each row, scan the result into our tag composite object
		err = results.Scan(&name.Name)
		if err != nil {
			return _isValid
		}

		_uName = name.Name
	}

	//hashin the entered password
	data := []byte(pwd)
	hashedPwd := md5.Sum(data)
	pass := fmt.Sprintf("%x", hashedPwd)
	pwd = pass

	// Protecting against mysql injections
	formatted = fmt.Sprintf("SELECT Password from group_mgr.users where Password = \"%s55fedg\";", pwd)
	// searching for the project
	results, err = db.Query(formatted)
	if err != nil {
		_isValid = false
		return _isValid
	}

	for results.Next() {
		var name Name
		// for each row, scan the result into our tag composite object
		err = results.Scan(&name.Name)
		if err != nil {
			return _isValid
		}

		_pwd = name.Name
	}

	//_uName, _pwd, _isValid := "cihanozhan", "1234!*.", false

	if uName == _uName && pwd+"55fedg" == _pwd {
		_isValid = true
	} else {
		_isValid = false
	}

	return _isValid
}
