// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// func main() {
// 	// Example connection string
// 	// Format: postgres://username:password@host:port/database
// 	databaseURL := "postgres://postgres:password@localhost:5432/mydb"

// 	// Use environment variable in production
// 	if os.Getenv("DATABASE_URL") != "" {
// 		databaseURL = os.Getenv("DATABASE_URL")
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// Create connection pool
// 	pool, err := pgxpool.New(ctx, databaseURL)
// 	if err != nil {
// 		log.Fatalf("Unable to create connection pool: %v\n", err)
// 	}
// 	defer pool.Close()

// 	// Test connection
// 	err = pool.Ping(ctx)
// 	if err != nil {
// 		log.Fatalf("Unable to connect to database: %v\n", err)
// 	}

// 	fmt.Println("Successfully connected to PostgreSQL!")

// 	// Example query
// 	var now time.Time
// 	err = pool.QueryRow(ctx, "SELECT NOW()").Scan(&now)
// 	if err != nil {
// 		log.Fatalf("Query failed: %v\n", err)
// 	}

// 	fmt.Println("Current time from DB:", now)
// }
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func main() {
	ctx := context.Background()

	databaseURL := "postgres://postgres:password@localhost:5432/mydb"
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// CREATE
	userID, err := CreateUser(ctx, pool, "John Doe", "john@example.com")
	if err != nil {
		log.Fatal("Create error:", err)
	}
	fmt.Println("Created user with ID:", userID)

	// READ (Get by ID)
	user, err := GetUserByID(ctx, pool, userID)
	if err != nil {
		log.Fatal("Read error:", err)
	}
	fmt.Println("Fetched user:", user)

	// UPDATE
	err = UpdateUser(ctx, pool, userID, "John Updated", "john_updated@example.com")
	if err != nil {
		log.Fatal("Update error:", err)
	}
	fmt.Println("User updated")

	// DELETE
	err = DeleteUser(ctx, pool, userID)
	if err != nil {
		log.Fatal("Delete error:", err)
	}
	fmt.Println("User deleted")
}

// CREATE
func CreateUser(ctx context.Context, pool *pgxpool.Pool, name, email string) (int, error) {
	var id int
	err := pool.QueryRow(ctx,
		`INSERT INTO users (name, email)
		 VALUES ($1, $2)
		 RETURNING id`,
		name, email).Scan(&id)

	return id, err
}

// READ
func GetUserByID(ctx context.Context, pool *pgxpool.Pool, id int) (*User, error) {
	var user User

	err := pool.QueryRow(ctx,
		`SELECT id, name, email, created_at
		 FROM users
		 WHERE id = $1`,
		id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UPDATE
func UpdateUser(ctx context.Context, pool *pgxpool.Pool, id int, name, email string) error {
	_, err := pool.Exec(ctx,
		`UPDATE users
		 SET name = $1, email = $2
		 WHERE id = $3`,
		name, email, id)

	return err
}

// DELETE
func DeleteUser(ctx context.Context, pool *pgxpool.Pool, id int) error {
	_, err := pool.Exec(ctx,
		`DELETE FROM users WHERE id = $1`,
		id)

	return err
}
