package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    // Формируем JSON-ответ в виде строки с фиксированным форматом.
    // Обратите внимание, что мы используем raw string literal (в обратных кавычках),
    // чтобы включить двойные кавычки без необходимости их экранировать.
    // Также используем глагол %q, чтобы автоматически добавить двойные кавычки вокруг значений.
    env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version": version,
		},
	}
	err := app.WriteJson(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)	
	}
}