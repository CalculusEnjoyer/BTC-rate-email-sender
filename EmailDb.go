package main

import (
	"encoding/csv"
	"errors"
	"os"
)

func addEmailToCsv(email string) *RequestError {
	file, err1 := os.OpenFile("emails.csv", os.O_WRONLY|os.O_APPEND|os.O_RDWR, 0644)
	emailArr := []string{email}
	csvwriter := csv.NewWriter(file)
	defer file.Close()
	defer csvwriter.Flush()
	badRead := RequestError{StatusCode: 500, Err: errors.New("Can not access data base")}

	if err1 == nil {
		file, _ := os.Open("emails.csv")
		cvsReader, errorRead := csv.NewReader(file).ReadAll()
		if errorRead == nil {
			for i := range cvsReader {
				if cvsReader[i][0] == email {
					return &RequestError{
						StatusCode: 409,
						Err:        errors.New("This email already exist in database"),
					}
				}
			}
		}
		if err2 := csvwriter.Write(emailArr); err2 != nil {
			return &badRead
		}
	} else {
		file1, err2 := os.Create("emails.csv")
		defer file1.Close()
		csvwriter1 := csv.NewWriter(file1)
		defer csvwriter1.Flush()
		if err2 == nil {
			if err4 := csvwriter1.Write([]string{"emails"}); err4 != nil {
				return &badRead
			}
			if err3 := csvwriter1.Write(emailArr); err3 != nil {
				return &badRead
			}
		}
	}
	return nil
}
