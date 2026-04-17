// internal/database/pgerror.go
package database

import (
	"errors"
	"fmt"
	"note-app-api/internal/features/apperr"

	"github.com/jackc/pgx/v5/pgconn"
)

// HandlePostgresError acts as a central translator for all database errors.
func HandlePostgresError(err error, conflictMessage ...string) error {
	// If there is no error, just return nil
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		// Use a switch statement for easy scalability
		switch pgErr.Code {
		case "23505": // Unique Violation (Duplicate)

			msg := "The data already exists."

			// If the developer provided an optional message, grab the first one
			if len(conflictMessage) > 0 && conflictMessage[0] != "" {
				msg = conflictMessage[0] + ", " + msg
			}

			return apperr.NewConflict(msg, err)

		case "23503": // Foreign Key Violation (Bonus 10/10 feature!)
			// Example: Trying to save a Note to a FolderID that doesn't exist
			return apperr.NewBadRequest("Invalid reference. The linked data does not exist.", err)
		}
	}

	// If it's a random database crash or a code we haven't mapped,
	// wrap it in our safe 500 Internal Error.
	return apperr.NewInternal(fmt.Errorf("database operation failed: %w", err))
}
