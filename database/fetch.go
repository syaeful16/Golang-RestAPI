package database

import (
	"fmt"
	"go-restapi/helper"
	"reflect"
)

func AllData(table string, data interface{}) error {
	// Connect ke database
	db, err := Connect()
	helper.BasicHandler(err)
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM %s", table)

	rows, err := db.Query(query)
	helper.BasicHandler(err)
	defer rows.Close()

	// Mendapatkan informasi tentang struktur dari data
	sliceValue := reflect.ValueOf(data).Elem()
	structType := sliceValue.Type().Elem()
	fieldsCount := structType.NumField()

	// Membuat slice untuk menyimpan nilai-nilai dari baris-baris hasil query
	values := make([]interface{}, fieldsCount)
	for i := 0; i < fieldsCount; i++ {
		values[i] = reflect.New(structType.Field(i).Type).Interface()
	}

	// Iterasi melalui setiap baris hasil query
	for rows.Next() {
		// Membuat instance baru dari struct untuk menyimpan nilai dari baris saat ini
		newStruct := reflect.New(structType).Elem()

		// Scan nilai-nilai dari baris ke dalam values
		if err := rows.Scan(values...); err != nil {
			return err
		}

		// Menyalin nilai-nilai dari values ke dalam struct
		for i := 0; i < fieldsCount; i++ {
			newStruct.Field(i).Set(reflect.ValueOf(values[i]).Elem())
		}

		// Menambahkan struct ke dalam slice
		sliceValue.Set(reflect.Append(sliceValue, newStruct))
	}

	fmt.Println(sliceValue)

	return nil
}
