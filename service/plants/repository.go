package plants

import "github.com/jackc/pgx/v4/pgxpool"

type Repository struct {
	pgx *pgxpool.Pool
}

func NewRepository(pgx *pgxpool.Pool) Repository {
	return Repository{
		pgx: pgx,
	}
}
