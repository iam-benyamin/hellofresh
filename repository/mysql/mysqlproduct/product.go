package mysqlproduct

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iam-benyamin/hellofresh/entity/productentity"
	"github.com/iam-benyamin/hellofresh/pkg/errmsg"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
)

func scanProduct(scanner mysql.Scanner) (productentity.Product, error) {
	var product productentity.Product

	err := scanner.Scan(&product.ID, &product.Name, &product.Code, &product.Price, &product.CreatedAt)

	return product, err
}

func (d *DB) GetProductByProductCode(ctx context.Context, productCode string) (productentity.Product, error) {
	const op = "mysqlproduct.GetProductByProductCode"

	row := d.conn.Conn().QueryRowContext(ctx, `SELECT * FROM products WHERE product_code = ?`, productCode)

	product, err := scanProduct(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return productentity.Product{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

		return productentity.Product{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return product, nil
}
