package studentController

import (
	"database/sql"
	"go-restapi/database"
	"go-restapi/helper"
	"go-restapi/models"
	"log"
)

func Index() ([]models.Student, error) {
	db, err := database.Connect()
	helper.BasicHandler(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tb_student")
	helper.BasicHandler(err)
	defer rows.Close()

	var result []models.Student
	for rows.Next() {
		var each = models.Student{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Age, &each.Grade)

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		result = append(result, each)

	}

	return result, nil
}

func Show(id any) (models.Student, error) {
	db, err := database.Connect()
	helper.BasicHandler(err)
	defer db.Close()

	var result = models.Student{}
	err = db.QueryRow("SELECT * FROM tb_student WHERE id = ?", id).Scan(&result.ID, &result.Name, &result.Age, &result.Grade)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Tidak ada data yang ditemukan untuk ID", id)
		} else {
			log.Println("Kesalahan saat melakukan query:", err)
		}
		return result, err
	}

	return result, err
}
