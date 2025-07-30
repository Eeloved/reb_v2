package main

import (
    "database/sql"
    "fmt"
    "log"
    "sync"
    "time"

    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "myuser"
    password = "mypassword"
    dbname   = "mydb"
)

func main() {
    // Строка подключения
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Открываем соединение с базой
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Ошибка при подключении к БД:", err)
    }
    defer db.Close()

    // Проверка подключения
    if err = db.Ping(); err != nil {
        log.Fatal("Не удалось подключиться к БД:", err)
    }

    fmt.Println("✅ Подключено к PostgreSQL")

    // Группа для синхронизации горутин
    var wg sync.WaitGroup

    // Количество параллельных соединений
    connections := 5

    // Запускаем 5 горутин
    for i := 1; i <= connections; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            // Выполняем запрос в каждой горутине
            var now time.Time
            err := db.QueryRow("SELECT NOW()").Scan(&now)
            if err != nil {
                log.Printf("Горутина %d: ошибка запроса: %v", id, err)
                return
            }

            fmt.Printf("Горутина %d: текущее время из БД: %s\n", id, now.Format(time.RFC3339))
        }(i)
    }

    // Ждём завершения всех горутин
    wg.Wait()

    fmt.Println("✅ Все 5 параллельных запросов завершены")
}