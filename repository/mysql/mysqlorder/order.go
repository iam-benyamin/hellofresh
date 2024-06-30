package mysqlorder

import (
	"context"
	"math/rand"
	"time"

	"github.com/iam-benyamin/hellofresh/param/orderparam"
	"github.com/iam-benyamin/hellofresh/pkg/errmsg"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
)

func (d *DB) SaveOrder(ctx context.Context, co orderparam.SaveOrder) (string, error) {
	const op = "mysqlorder.saveOrder"
	newID := generateNewID()

	_, err := d.conn.Conn().ExecContext(ctx,
		`INSERT INTO orders(id, user_id, product_code, customer_full_name, product_name, total_amount) values(?, ?, ?, ?, ?, ?)`,
		newID, co.UserID, co.ProductCode, co.CustomerFullName, co.ProductName, co.TotalAmount)
	if err != nil {
		return "", richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindNotFound)
	}

	return newID, nil
}

func generateNewID() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 12)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}
