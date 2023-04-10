package mysqlerr

import (
	"regexp"

	"github.com/go-sql-driver/mysql"
)

var duplicateEntryRegexp = regexp.MustCompile(`Duplicate entry '(.+)' for key '(.+)'`)

type DuplicateEntryError struct {
	*mysql.MySQLError
	entry string
	key   string
}

func NewDuplicateEntryError(err *mysql.MySQLError) *DuplicateEntryError {
	res := &DuplicateEntryError{MySQLError: err}

	matches := duplicateEntryRegexp.FindStringSubmatch(err.Message)
	if len(matches) == 3 {
		res.entry = matches[1]
		res.key = matches[2]
	}

	return res
}

func (e *DuplicateEntryError) Entry() string {
	return e.entry
}

func (e *DuplicateEntryError) Key() string {
	return e.key
}

func CastToDuplicateEntryError(err error) (*DuplicateEntryError, bool) {
	mErr, ok := CastToMySQLError(err)
	if !ok {
		return nil, false
	}

	if mErr.Number != ErrDupEntry && mErr.Number != ErrDupEntryWithKeyName {
		return nil, false
	}

	return NewDuplicateEntryError(mErr), true
}
