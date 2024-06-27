package mysqluser

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iam-benyamin/hellofresh/entity/userentity"
	"github.com/iam-benyamin/hellofresh/pkg/errmsg"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
)

func scanUser(scanner mysql.Scanner) (userentity.User, error) {
	var user userentity.User

	err := scanner.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt)

	return user, err
}

func (d *DB) GetUserByID(ctx context.Context, userID string) (userentity.User, error) {
	const op = "mysqluser.GetUserByID"

	row := d.conn.Conn().QueryRowContext(ctx, `SELECT * FROM users WHERE id = ?`, userID)

	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return userentity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound)
		}

		return userentity.User{}, richerror.New(op).WithErr(err).WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return user, nil
}
