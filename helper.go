package mysqlerr

import "github.com/go-sql-driver/mysql"

// CheckCode check error number
func CheckCode(err error, code uint16) bool {
	mErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return false
	}

	return checkCode(mErr, code)
}

func checkCode(err *mysql.MySQLError, code uint16) bool {
	return err.Number == code
}

func IsMySQLError(err error) bool {
	_, ok := err.(*mysql.MySQLError)

	return ok
}

func CastToMySQLError(err error) (*mysql.MySQLError, bool) {
	mErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return nil, false
	}

	return mErr, true
}
