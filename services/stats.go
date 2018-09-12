package services

import (
	"mutant/db"
		"database/sql"
			"math"
)

func List() map[string]float64 {
	db := db.DbConn()

	defer db.Close()

	var totalHumans = 0.0
	var totalMutants = 0.0
	var ratio = 0.0

	mutants, err := db.Query("select count(dna) from mutants where is_mutant = 1")

	if err != nil {
		panic(err.Error())
	}
	defer mutants.Close()

	humans, err := db.Query("select count(dna)from mutants where is_mutant = 0")

	if err != nil {
		panic(err.Error())
	}
	defer humans.Close()

	totalMutants = float64(checkCount(mutants))
	totalHumans = float64(checkCount(humans))

	if totalHumans > 0 && totalMutants > 0 {
		ratio = math.Abs(totalMutants / totalHumans)
	}

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