package services

import (
	"mutant/db"
	"database/sql"
	"log"
)

type Mutant struct {
	ID   int    `json:"id"`
	Dna string `json:"dna"`
	IsMutant bool `json:"is_mutant"`
}

func Save(dna string, isMutant int)  {
	db := db.DbConn()

	defer db.Close()

	var mutant Mutant

	err := db.QueryRow("SELECT dna FROM mutants WHERE dna = ?", dna).Scan(&mutant.Dna)
	if err != nil {
		if err == sql.ErrNoRows {
			insertDna(dna, isMutant)
		}
	}
}

func insertDna(dna string, isMutant int) {
	db := db.DbConn()

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO mutants(dna, is_mutant) VALUES(?, ?)")

	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(dna, isMutant)

	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}