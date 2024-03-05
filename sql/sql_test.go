package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestInsertExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customers (id, name) VALUES ('arra', 'Arra')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Customer successfully inserted")
}

func TestUpdateExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "UPDATE customers SET id = 'aryn', name = 'Aryn' WHERE id = 'dila'"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Customer successfully updated")
}

func TestDeleteExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "DELETE FROM customers WHERE id = 'arra'"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Customer successfully deleted")
}

func TestSelectQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customers"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}

	fmt.Println("Customer successfully selected")
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, created_at, birth_date, is_student FROM customers"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var isStudent bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birthDate, &isStudent)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		if email.Valid {
			fmt.Println("email:", email.String)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		fmt.Println("createdAt:", createdAt)
		if birthDate.Valid {
			fmt.Println("birthDate:", birthDate.Time)
		}
		fmt.Println("isStudent:", isStudent)
	}
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin' #"
	password := "password"

	query := "SELECT username FROM users WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login succeed", username)
	} else {
		fmt.Println("Login failed")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin' #"
	password := "password"

	query := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login succeed", username)
	} else {
		fmt.Println("Login failed")
	}
}

func TestExecSQLParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "bima'; DROP TABLE user; #"
	password := "bima"

	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("User successfully inserted")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "bima@mail.com"
	comment := "first comment"

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Comment with id successfully inserted", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "bima" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"

	for i := 0; i < 10; i++ {
		email := "bima" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment id", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
