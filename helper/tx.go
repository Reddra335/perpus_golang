package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {

	err := recover()

	if err != nil {
		ErrorRollback := tx.Rollback()
		ErrorT(ErrorRollback)
		panic(err)
	} else {

		ErrCommit := tx.Commit()
		ErrorT(ErrCommit)
	}
}
