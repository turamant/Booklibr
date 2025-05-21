package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	// Создаем новый экземпляр маршрутизатора httprouter.
	router := httprouter.New()

	// Преобразуем notFoundResponse() в http.Handler через http.HandlerFunc(),
    // и устанавливаем его как кастомный обработчик для 404 Not Found.
    router.NotFound = http.HandlerFunc(app.notFoundResponse)

    // То же самое делаем с methodNotAllowedResponse() для обработки 405 Method Not Allowed.
    router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Регистрируем соответствующие методы, URL-паттерны и обработчики для наших
	// конечных точек с помощью метода HandlerFunc(). Обратите внимание, что
	// http.MethodGet и http.MethodPost — это константы, равные строкам "GET" и "POST".
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/books", app.createBookHandler)
	router.HandlerFunc(http.MethodGet, "/v1/books/:id", app.showBookHandler)

	// Возвращаем экземпляр маршрутизатора.
	return router
}
