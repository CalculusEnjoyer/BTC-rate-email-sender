package main

import (
	"encoding/csv"
	"errors"
	"os"
)

func addEmailToCsv(email string) *RequestError {
	file, err1 := os.OpenFile("emails.csv", os.O_APPEND|os.O_RDWR, os.ModePerm)
	emailArr := []string{email}
	csvwriter := csv.NewWriter(file)
	defer file.Close()
	defer csvwriter.Flush()
	badRead := RequestError{StatusCode: 500, Err: errors.New("can not access data base")}

	if err1 == nil {
		cvsReader, _ := csv.NewReader(file).ReadAll()
		for i := range cvsReader {
			if cvsReader[i][0] == email {
				return &RequestError{
					StatusCode: 409,
					Err:        errors.New("this email already exist in database"),
				}
			}
		}
		if err2 := csvwriter.Write(emailArr); err2 != nil {
			return &badRead
		} else {
			return nil
		}
	} else {
		return &badRead
	}
}
