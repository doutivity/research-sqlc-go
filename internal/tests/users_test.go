package tests

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/doutivity/research-sqlc-go/internal/storage/postgres/dbs"

	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

const (
	dataSourceName = "postgresql://yaroslav:AnySecretPassword!!@postgres:5432/yaaws?sslmode=disable&timezone=UTC"
)

func TestUsers(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	connection, err := sql.Open("postgres", dataSourceName)
	require.NoError(t, err)
	defer connection.Close()

	require.NoError(t, connection.Ping())

	// clear users
	{
		const (
			// language=Postgres
			query = "TRUNCATE TABLE users RESTART IDENTITY CASCADE"
		)

		_, err := connection.Exec(query)
		require.NoError(t, err)
	}

	var queries = dbs.New(connection)

	// db.go
	_ = queries.WithTx

	// models.go
	var _ dbs.User

	// users.sql.go
	_ = queries.UserNew
	_ = queries.UserNewAndGet
	_ = queries.UserGetByID
	_ = queries.UsersCount
	_ = queries.Users

	var ctx = context.Background()

	{
		count, err := queries.UsersCount(ctx)
		require.NoError(t, err)
		require.Equal(t, int64(0), count)
	}

	var now = time.Now().UTC().Truncate(time.Second)

	var (
		expected = dbs.User{
			ID:        1,
			Name:      "System",
			CreatedAt: now,
		}
	)

	// create
	{
		actual, err := queries.UserNewAndGet(ctx, dbs.UserNewAndGetParams{
			Name:      "System",
			CreatedAt: now,
		})

		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}

	// get by id
	{

		actual, err := queries.UserGetByID(ctx, 1)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}

	// get count
	{
		count, err := queries.UsersCount(ctx)
		require.NoError(t, err)
		require.Equal(t, int64(1), count)
	}

	// get users
	{
		users, err := queries.Users(ctx, dbs.UsersParams{
			Offset: 0,
			Limit:  32,
		})
		require.NoError(t, err)
		require.Equal(t, []dbs.User{expected}, users)
	}
}
