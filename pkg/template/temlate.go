package template

import (
	"context"
	"html/template"
	"net/http"
)

func ExecuteTemplate(ctx context.Context, w http.ResponseWriter, files []string, data map[string]interface{}) error {
	ts, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	err = ts.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
