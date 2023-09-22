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

func TestVacancies(t *testing.T) {
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

	// clear companies
	{
		const (
			// language=Postgres
			query = "TRUNCATE TABLE companies RESTART IDENTITY CASCADE"
		)

		_, err := connection.Exec(query)
		require.NoError(t, err)
	}

	// clear vacancies
	{
		const (
			// language=Postgres
			query = "TRUNCATE TABLE vacancies RESTART IDENTITY CASCADE"
		)

		_, err := connection.Exec(query)
		require.NoError(t, err)
	}

	var queries = dbs.New(connection)
	var ctx = context.Background()
	var now = time.Now().UTC().Truncate(time.Second)

	// create user
	{
		err := queries.UserNew(ctx, dbs.UserNewParams{
			Name:      "System",
			CreatedAt: now,
		})

		require.NoError(t, err)
	}

	{
		id, err := queries.CompanyNewAndGetID(ctx, dbs.CompanyNewAndGetIDParams{
			Alias:     "yaaws",
			Name:      "YAAWS",
			CreatedBy: 1,
			CreatedAt: now,
		})

		require.NoError(t, err)
		require.Equal(t, int64(1), id)
	}

	{
		err := queries.VacancyNew(ctx, dbs.VacancyNewParams{
			CompanyID: 1,
			Title:     "Senior Rust Developer",
			CreatedBy: 1,
			CreatedAt: now,
		})

		require.NoError(t, err)
	}

	{
		expected := []dbs.VacancyByCompaniesRow{
			{
				VacancyID:    1,
				VacancyTitle: "Senior Rust Developer",
				CompanyID:    1,
				CompanyAlias: "yaaws",
				CompanyName:  "YAAWS",
			},
		}

		actual, err := queries.VacancyByCompanies(ctx, []int64{1})
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}
}
