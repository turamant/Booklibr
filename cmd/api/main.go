package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Объявляем строковую константу с номером версии приложения.
// Позже в книге мы будем генерировать это значение автоматически во время сборки,
// но пока просто сохраним его как жёстко заданную глобальную константу.
const version = "1.0.0"

// Определяем структуру config для хранения всех конфигурационных настроек приложения.
// Пока что здесь будут только номер порта, на котором будет слушать сервер,
// и имя текущей операционной среды (development, staging, production и т.д.).
// Эти настройки мы будем считывать из флагов командной строки при запуске приложения.
type config struct {
	port int
	env  string
}

// Определяем структуру application для хранения зависимостей наших обработчиков,
// вспомогательных функций и middleware. Пока она содержит только копию структуры
// config и логгер, но по мере разработки в неё будет добавляться всё больше данных.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Создаем экземпляр структуры config
	var cfg config

	// Считываем значения флагов port и env в структуру config.
	// По умолчанию используем порт 4000 и среду "development", если соответствующие
	// флаги не были переданы при запуске.
	flag.IntVar(&cfg.port, "port", 4000, "Порт для HTTP-сервера")
	flag.StringVar(&cfg.env, "env", "development", "Среда (development|staging|production)")
	flag.Parse()

	// Инициализируем новый логгер, который выводит сообщения в стандартный поток вывода,
	// добавляя к ним дату и время.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Создаем экземпляр структуры application, содержащей конфигурацию и логгер.
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Создаём HTTP-сервер с разумными настройками таймаутов, который слушает порт,
	// указанный в конфигурации, и использует наш multiplexer как обработчик.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Запускаем HTTP-сервер
	logger.Printf("запуск сервера в среде %s на адресе %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
