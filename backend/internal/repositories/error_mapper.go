package repositories

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	uniqueViolation  = "23505"
	notNullViolation = "23502"
)

// errorMapper is private function that only
// exists on repositories package. errorMapper
// handles error returned from database
// to more readble error (only handles postgres for now)
func errorMapper(err error) error {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		// unique violation
		if pgErr.Code == uniqueViolation {
			if strings.Contains(pgErr.ConstraintName, "email") {
				return ErrEmailExists
			}

			return errors.New(fmt.Sprintf("unique violation detected on %s", pgErr.ConstraintName))
		}

		// not null constraint error
		if pgErr.Code == notNullViolation {
			return errors.New(fmt.Sprintf("not null violation detected on %s", pgErr.ConstraintName))
		}
	}

	// if no match, return the same error
	return err
}
