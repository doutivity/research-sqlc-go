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

func TestReviews(t *testing.T) {
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

	// clear reviews
	{
		const (
			// language=Postgres
			query = "TRUNCATE TABLE reviews RESTART IDENTITY CASCADE"
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
		id1, err := queries.ReviewNew(ctx, dbs.ReviewNewParams{
			ParentID:  sql.NullInt64{},
			CompanyID: 1,
			Content:   "one",
			CreatedBy: 1,
			CreatedAt: now,
		})
		require.NoError(t, err)
		require.Equal(t, int64(1), id1)

		// nested
		{
			parent := sql.NullInt64{
				Int64: id1,
				Valid: true,
			}

			id2, err := queries.ReviewNew(ctx, dbs.ReviewNewParams{
				ParentID:  parent,
				CompanyID: 1,
				Content:   "two",
				CreatedBy: 1,
				CreatedAt: now,
			})
			require.NoError(t, err)
			require.Equal(t, int64(2), id2)

			id3, err := queries.ReviewNew(ctx, dbs.ReviewNewParams{
				ParentID:  parent,
				CompanyID: 1,
				Content:   "three",
				CreatedBy: 1,
				CreatedAt: now,
			})
			require.NoError(t, err)
			require.Equal(t, int64(3), id3)

			{
				id4, err := queries.ReviewNew(ctx, dbs.ReviewNewParams{
					ParentID: sql.NullInt64{
						Int64: id3,
						Valid: true,
					},
					CompanyID: 1,
					Content:   "four",
					CreatedBy: 1,
					CreatedAt: now,
				})
				require.NoError(t, err)
				require.Equal(t, int64(4), id4)

				{
					id5, err := queries.ReviewNew(ctx, dbs.ReviewNewParams{
						ParentID: sql.NullInt64{
							Int64: id4,
							Valid: true,
						},
						CompanyID: 1,
						Content:   "five",
						CreatedBy: 1,
						CreatedAt: now,
					})
					require.NoError(t, err)
					require.Equal(t, int64(5), id5)
				}
			}
		}
	}

	// root reviews
	{
		expected := []dbs.Review{
			{
				ID:        1,
				ParentID:  sql.NullInt64{},
				CompanyID: 1,
				Content:   "one",
				CreatedBy: 1,
				CreatedAt: now,
			},
		}

		actual, err := queries.ReviewsRootByCompany(ctx, 1)
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}

	// nested reviews
	{
		expected := []dbs.ReviewsNestedRow{
			{
				ID:        1,
				ParentID:  sql.NullInt64{},
				CompanyID: 1,
				Content:   "one",
				CreatedBy: 1,
				CreatedAt: now,
			},
			{
				ID: 2,
				ParentID: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				CompanyID: 1,
				Content:   "two",
				CreatedBy: 1,
				CreatedAt: now,
			},
			{
				ID: 3,
				ParentID: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				CompanyID: 1,
				Content:   "three",
				CreatedBy: 1,
				CreatedAt: now,
			},
			{
				ID: 4,
				ParentID: sql.NullInt64{
					Int64: 3,
					Valid: true,
				},
				CompanyID: 1,
				Content:   "four",
				CreatedBy: 1,
				CreatedAt: now,
			},
			{
				ID: 5,
				ParentID: sql.NullInt64{
					Int64: 4,
					Valid: true,
				},
				CompanyID: 1,
				Content:   "five",
				CreatedBy: 1,
				CreatedAt: now,
			},
		}

		actual, err := queries.ReviewsNested(ctx, dbs.ReviewsNestedParams{
			CompanyID: 1,
			ID:        1,
		})
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}
}
