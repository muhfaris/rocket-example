package mysqladapter

import (
	"context"
	"database/sql"
	"fmt"
)

type Client struct {
	conn *sql.DB
}

func (c *Client) WithTx(ctx context.Context, funcs ...func(tx *sql.Tx) (interface{}, error)) (interface{}, error) {
	tx, err := c.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	for _, fn := range funcs {
		_, err := fn(tx)
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return nil, fmt.Errorf("transaction error: %v, rollback error: %v", err, rollbackErr)
			}
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}
