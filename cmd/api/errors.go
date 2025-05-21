package main

import (
    "fmt"
    "net/http"
)

// Метод logError() — универсальный помощник для логирования ошибок.
// Позже мы улучшим его, чтобы он сохранял структурированную информацию,
// включая HTTP-метод и URL запроса.
func (app *application) logError(r *http.Request, err error) {
    app.logger.Println(err)
}

// Метод errorResponse() — универсальный помощник для отправки JSON-ошибок клиенту
// с указанным HTTP-статусом. Обратите внимание, что мы используем тип interface{}
// для параметра message — это даёт нам больше гибкости в том, какие данные мы можем
// включить в ответ.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
    env := envelope{"error": message}

    // Используем writeJSON() для отправки ответа. Если эта операция завершится ошибкой,
    // логируем её и отправим клиенту пустой ответ с кодом 500 Internal Server Error.
    err := app.WriteJson(w, status, env, nil)
    if err != nil {
        app.logError(r, err)
        w.WriteHeader(500)
    }
}

// serverErrorResponse() используется, когда приложение сталкивается с неожиданной
// внутренней проблемой. Он записывает подробное сообщение об ошибке в лог и отправляет
// клиенту ответ с кодом 500 Internal Server Error и общим сообщением об ошибке.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
    app.logError(r, err)
    message := "the server encountered a problem and could not process your request"
    app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse() отправляет клиенту ответ с кодом 404 Not Found и соответствующее
// JSON-сообщение.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
    message := "the requested resource could not be found"
    app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse() отправляет клиенту ответ с кодом 405 Method Not Allowed
// и информацией о том, какой метод не поддерживается.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
    message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
    app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}