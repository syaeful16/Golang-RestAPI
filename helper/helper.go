package helper

import "log"

func BasicHandler(err error) error {
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
