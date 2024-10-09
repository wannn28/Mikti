package main

import (
    "database/sql"
    "net/http"

    "github.com/labstack/echo/v4"
    _ "modernc.org/sqlite"
)

type Task struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var db *sql.DB

func main() {
    var err error

    // Inisialisasi database SQLite
    db, err = sql.Open("sqlite", "./tasks.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Buat tabel jika belum ada
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)`)
    if err != nil {
        panic(err)
    }

    // Inisialisasi instance Echo
    e := echo.New()

    // Definisikan rute CRUD
    e.POST("/tasks", createTask)       // Buat tugas
    e.GET("/tasks", getTasks)           // Ambil semua tugas
    e.GET("/tasks/:id", getTask)        // Ambil tugas berdasarkan ID
    e.PUT("/tasks/:id", updateTask)     // Perbarui tugas
    e.DELETE("/tasks/:id", deleteTask)  // Hapus tugas

    // Mulai server di localhost:8080
    e.Logger.Fatal(e.Start(":8080"))
}

// Buat tugas
func createTask(c echo.Context) error {
    task := new(Task)
    if err := c.Bind(task); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    result, err := db.Exec("INSERT INTO tasks (name) VALUES (?)", task.Name)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    id, _ := result.LastInsertId()
    task.ID = int(id)
    return c.JSON(http.StatusCreated, task)
}

// Ambil semua tugas
func getTasks(c echo.Context) error {
    rows, err := db.Query("SELECT id, name FROM tasks")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    defer rows.Close()

    var tasks []Task
    for rows.Next() {
        var task Task
        if err := rows.Scan(&task.ID, &task.Name); err != nil {
            return c.JSON(http.StatusInternalServerError, err)
        }
        tasks = append(tasks, task)
    }

    return c.JSON(http.StatusOK, tasks)
}

// Ambil satu tugas berdasarkan ID
func getTask(c echo.Context) error {
    id := c.Param("id")
    var task Task

    err := db.QueryRow("SELECT id, name FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Name)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, echo.Map{"error": "Tugas tidak ditemukan"})
        }
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, task)
}

// Perbarui tugas
func updateTask(c echo.Context) error {
    id := c.Param("id")
    task := new(Task)
    if err := c.Bind(task); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    _, err := db.Exec("UPDATE tasks SET name = ? WHERE id = ?", task.Name, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, echo.Map{"message": "Tugas berhasil diperbarui"})
}

// Hapus tugas
func deleteTask(c echo.Context) error {
    id := c.Param("id")

    _, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusOK, echo.Map{"message": "Tugas berhasil dihapus"})
}
