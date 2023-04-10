package mysqlerr

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCastToDuplicateEntryError_False(t *testing.T) {
	tests := map[string]error{
		"common":           errors.New("my error"),
		"ErrBadFieldError": &mysql.MySQLError{Number: ErrBadFieldError},
	}

	for name, args := range tests {
		t.Run(name, func(t *testing.T) {
			got, ok := CastToDuplicateEntryError(args)

			assert.Nil(t, got)
			assert.False(t, ok)
		})
	}
}

func TestCastToDuplicateEntryError_True(t *testing.T) {
	tests := map[string]struct {
		args  string
		entry string
		key   string
	}{
		"PRIMARY": {
			args:  "Duplicate entry '12' for key 'PRIMARY'",
			entry: "12",
			key:   "PRIMARY",
		},
		"email": {
			args:  "Duplicate entry 'test@test.com' for key 'email'",
			entry: "test@test.com",
			key:   "email",
		},
		"uuid": {
			args:  "Duplicate entry '0f608670-1d43-4d01-b5dc-e3c8d19403cc' for key 'uuid'",
			entry: "0f608670-1d43-4d01-b5dc-e3c8d19403cc",
			key:   "uuid",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, ok := CastToDuplicateEntryError(&mysql.MySQLError{
				Number:  ErrDupEntry,
				Message: tt.args,
			})

			assert.NotNil(t, got)
			assert.True(t, ok)
			assert.Equal(t, tt.entry, got.Entry())
			assert.Equal(t, tt.key, got.Key())
		})
	}
}

func BenchmarkSscanf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var entry, key string

		fmt.Sscanf(
			"Duplicate entry '12' for key 'PRIMARY'",
			"Duplicate entry '%s' for key '%s'",
			&entry,
			&key,
		)
	}
}

func BenchmarkRegexp(b *testing.B) {
	r := regexp.MustCompile(`Duplicate entry '(.+)' for key '(.+)'`)

	for i := 0; i < b.N; i++ {
		r.FindStringSubmatch("Duplicate entry '12' for key 'PRIMARY'")
	}
}
