package model

import "lordtt13/GO_Prods/TODO-API/views"

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}
