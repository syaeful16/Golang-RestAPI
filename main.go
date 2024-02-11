package main

import (
	"go-restapi/database"
	"go-restapi/models"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	db, err := database.Connect()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM tb_student")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	var result []models.Student
// 	for rows.Next() {
// 		var each = student
// 		var err = rows.Scan(&each.ID, &each.Name, &each.Age, &each.Grade)

// 		if err != nil {
// 			fmt.Print(err.Error())
// 			return
// 		}

// 		result = append(result, each)
// 	}

// 	if err := rows.Err(); err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	for _, data := range result {
// 		fmt.Println(data.Name)
// 	}

// 	data := map[string]interface{}{
// 		"data": result,
// 	}

// 	jsonData, err := json.Marshal(data)

// 	w.Write(jsonData)

// }

func main() {
	var students []models.Student
	err := database.AllData("tb_student", &students)
	if err != nil {
		return
	}

	// Mencetak data mahasiswa
	// log.Println("test")
	// http.HandleFunc("/", index)

	// fmt.Println("starting web server at http://localhost:8080/")
	// http.ListenAndServe(":8080", nil)
}
