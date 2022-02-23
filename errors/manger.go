package errores

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jackc/pgconn"
)

// action Postgres DB error handling
type action struct {
	Message          string
	HttpResponseCode int
	Loggable         bool
}

var pgErrorMessage = map[string]action{
	"22001": {"error, verify that the fields have the correct length of characters", http.StatusBadRequest, false},
	"23505": {"error, duplicate record", http.StatusBadRequest, false},
	"23514": {"error, wrong data format, see the documentation", http.StatusBadRequest, false},
	"23503": {"error, the resource is being used by other registries", http.StatusBadRequest, false},
	"23000": {"error, the restricted operation, data integrity problem, see documentation", http.StatusBadRequest, false},
	"25000": {"error, could not complete operations", http.StatusInternalServerError, true},
	"26000": {"error, there was an internal problem, please report the incident to the respective technical team", http.StatusInternalServerError, true},
	"28000": {"error, restricted access", http.StatusUnauthorized, true},
	"2D000": {"error, invalid transaction", http.StatusInternalServerError, true},
}

func NewInternalDB(err error) error {
	var postgresErr *pgconn.PgError
	if errors.As(err, &postgresErr) {
		if postgresErr.Code == "23503" && strings.Contains(postgresErr.Message, "insert or update") {
			const message = "error, incompatible resource association, check the existence of records"
			return newErr(err, message, http.StatusBadRequest)
		}

		act, ok := pgErrorMessage[postgresErr.Code]
		if ok {
			if act.Loggable {
				return newErr(err, act.Message, act.HttpResponseCode)
			}

			return newErr(nil, act.Message, act.HttpResponseCode)
		}
	}

	return newErr(err, ErrDatabaseInternal, http.StatusInternalServerError)
}
