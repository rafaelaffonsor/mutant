package main

import (
	"strings"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/mutant", isMutant).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}

type dnaStruct struct {
	Dna []string
}

func isMutant(w http.ResponseWriter, r *http.Request) {
	//matrix := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	//matrix := []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}
	decoder := json.NewDecoder(r.Body)

	var data dnaStruct

	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	matrix := data.Dna
	n := len(matrix)

	dna := ""
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i < n-3 {
				if j < n-3 {
					strD := ""
					gen := strings.Split(matrix[i], "")
					strD = strD + gen[j];
					strD += strings.Split(matrix[i+1], "")[j+1]
					strD += strings.Split(matrix[i+2], "")[j+2]
					strD += strings.Split(matrix[i+3], "")[j+3]
					dna += strD
				}

				strV := ""
				gen := strings.Split(matrix[i], "")
				strV = strV + gen[j];
				strV += strings.Split(matrix[i+1], "")[j]
				strV += strings.Split(matrix[i+2], "")[j]
				strV += strings.Split(matrix[i+3], "")[j]
				dna += strV
			}

			if (j < n-3) {
				strH := ""
				gen := strings.Split(matrix[i], "")
				strH += gen[j];
				strH += gen[j+1]
				strH += gen[j+2]
				strH += gen[j+3]
				dna += strH
			}
		}
	}

	result := checkDna(dna)

	if result {
		json.NewEncoder(w).Encode(result)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(result)
	}
}

func checkDna(dna string) bool  {
	count := 0
	accepted := [4]string{"AAAA", "GGGG", "CCCC", "TTTT"}

	for _,gene := range accepted {
		if strings.Contains(dna, gene) {
			count++
		}
	}

	if count > 1 {
		return true
	}

	return false
}