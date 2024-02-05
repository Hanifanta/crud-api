package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	} else {
		tx.Commit()
	}
}
