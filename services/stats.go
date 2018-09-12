package services

import (
	"mutant/db"
		"database/sql"
			"math"
)

func List() map[string]float64 {
	db := db.DbConn()

	defer db.Close()

	mutants, err := db.Query("select count(dna) from mutants where is_mutant = 1")

	if err != nil {
		if err == sql.ErrNoRows {
			panic(err.Error())
		}
		panic(err.Error())
	}
	defer mutants.Close()

	humans, err := db.Query("select count(dna)from mutants where is_mutant = 0")

	if err != nil {
		if err == sql.ErrNoRows {
			panic(err.Error())
		}
		panic(err.Error())
	}
	defer humans.Close()

	totalMutants := float64(checkCount(mutants))
	totalHumans := float64(checkCount(humans))
	ratio := math.Abs(totalMutants / totalHumans)

	result := map[string]float64{
		"count_mutant_dna":  totalMutants,
		"count_human_dna": totalHumans,
		"ratio": ratio,
	}

	return result
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err:= rows.Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	}
	return count
}