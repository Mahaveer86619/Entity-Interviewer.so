package config

import (
        "database/sql"
        "fmt"
        "log"
        "os"

        _ "github.com/lib/pq"
)

var db *sql.DB

// ConnectDB establishes a connection to the PostgreSQL database.
func ConnectDB() (*sql.DB, error) {
        dsn := os.Getenv("DATABASE_URL")
        if dsn == "" {
                return nil, fmt.Errorf("environment variable DATABASE_URL is not set")
        }

        conn, err := sql.Open("postgres", dsn)
        if err != nil {
                return nil, fmt.Errorf("failed to open database connection: %w", err)
        }

        // Verify database connectivity.
        if err := conn.Ping(); err != nil {
                return nil, fmt.Errorf("database ping failed: %w", err)
        }

        log.Println("Database connection established successfully")
        db = conn

        return db, nil
}

// SetDBConnection sets the global database connection.
func SetDBConnection(conn *sql.DB) {
        db = conn
}

// GetDBConnection retrieves the global database connection.
func GetDBConnection() *sql.DB {
        return db
}

// CloseDBConnection closes the provided database connection.
func CloseDBConnection(conn *sql.DB) {
        if conn == nil {
                return
        }

        if err := conn.Close(); err != nil {
                log.Printf("Error closing database connection: %v", err)
        } else {
                log.Println("Database connection closed successfully")
        }
}

// CreateTables creates or verifies the required database tables.
func CreateTables(conn *sql.DB) error {
        queries := []string{
                `CREATE TABLE IF NOT EXISTS users (
                        id UUID PRIMARY KEY,
                        name TEXT NOT NULL,
                        email TEXT UNIQUE NOT NULL,
                        password TEXT NOT NULL,
                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW()
                );`,
                `CREATE TABLE IF NOT EXISTS job_descriptions (
                        id UUID PRIMARY KEY,
                        user_id UUID NOT NULL REFERENCES users(id),
                        description TEXT NOT NULL,
                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW()
                );`,
                `CREATE TABLE IF NOT EXISTS interview_questions (
                        id UUID PRIMARY KEY,
                        job_description_id UUID NOT NULL REFERENCES job_descriptions(id),
                        question TEXT NOT NULL,
                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW()
                );`,
                `CREATE TABLE IF NOT EXISTS tts_audio (
                        id UUID PRIMARY KEY,
                        question_id UUID NOT NULL REFERENCES interview_questions(id),
                        audio_path TEXT NOT NULL,
                        created_at TIMESTAMP DEFAULT NOW()
                );`,
                `CREATE TABLE IF NOT EXISTS sessions (
                        id UUID PRIMARY KEY,
                        user_id UUID REFERENCES users(id),
                        session_token TEXT UNIQUE NOT NULL,
                        expires_at TIMESTAMP NOT NULL,
                        created_at TIMESTAMP DEFAULT NOW()
                );`,
                `CREATE TABLE IF NOT EXISTS forgot_password (
                        id UUID PRIMARY KEY,
                        email TEXT UNIQUE NOT NULL,
                        code TEXT UNIQUE NOT NULL
                );`,
        }

        for _, query := range queries {
                _, err := conn.Exec(query)
                if err != nil {
                        return fmt.Errorf("failed to create/verify table: %w", err)
                }
        }
        log.Println("Database tables created/verified successfully")

        return nil
}