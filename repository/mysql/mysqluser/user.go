package mysqluser

import (
	"context"
	"database/sql"
	"errors"
	"github.com/iam-benyamin/hellofresh/entity"
	"github.com/iam-benyamin/hellofresh/pkg/errmsg"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
)

func scanUser(scanner mysql.Scanner) (entity.User, error) {
	var user entity.User

	err := scanner.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt)

	return user, err
}

func (d *DB) GetUserByID(ctx context.Context, UserID string) (entity.User, error) {
	const op = "mysqluser.GetUserByID"

	row := d.conn.Conn().QueryRow(`SELECT * FROM users WHERE id = ?`, UserID)

	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

		return entity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return user, nil
}
