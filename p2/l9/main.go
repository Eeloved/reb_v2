package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Обработчик для корня (/)
func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет с главной страницы!")
}

// Обработчик для /about
func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это страница 'О нас'")
}

// Обработчик для /contact
func handleContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Контактная информация здесь.")
}

func handleLongPing(w http.ResponseWriter, r *http.Request) {
	// Создаем контекст с таймаутом 1 секунда
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	// Канал для имитации долгой работы
	done := make(chan struct{})

	go func() {
		// Имитация долгой работы (например, 5 секунд)
		time.Sleep(2 * time.Second)
		close(done)
	}()

	// Ожидаем завершения или таймаута
	select {
	case <-ctx.Done():
		// Таймаут или отмена контекста
		http.Error(w, "Request Timeout", http.StatusRequestTimeout)
	case <-done:
		// Успешная обработка
		fmt.Fprintln(w, "Обработка завершена")
	}
}
func main() {
	// Регистрируем обработчики
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/contact", handleContact)
	http.HandleFunc("/longping", handleLongPing)

	// Получаем текущую директорию
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка получения текущей директории:", err)
		return
	}

	// Обработчик /source/ для отдачи файлов
	fs := http.FileServer(http.Dir(currentDir))
	http.Handle("/source/", http.StripPrefix("/source/", fs))

	// Запускаем сервер
	fmt.Println("Сервер запущен на http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
