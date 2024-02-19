package postgresql

import (
	"github/Ndraaa15/fitness-tracker-api/internal/api/auth/repository"

	"github.com/jmoiron/sqlx"
)

type store struct {
	db *sqlx.DB
}

type storeClient struct {
	tx *sqlx.Tx
}

func NewStore(newDB *sqlx.DB) repository.AuthStoreImpl {
	return &store{
		db: newDB,
	}
}

func (s *store) NewStoreClient(useTx bool) (repository.AuthStoreClientImpl, error) {
	var (
		tx  *sqlx.Tx
		err error
	)

	if useTx {
		tx, err = s.db.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &storeClient{
		tx: tx,
	}, nil
}

func (sc *storeClient) Commit() error {
	if sc.tx != nil {
		return sc.tx.Commit()
	}

	return nil
}

func (sc *storeClient) Rollback() error {
	if sc.tx != nil {
		return sc.tx.Rollback()
	}

	return nil
}
