package mysqluser

//func TestDBGetUserByID(t *testing.T) {
//	// Create mock SQL connection
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			panic(err)
//		}
//	}(db)
//
//	// Create a new instance of DB with the mock connection
//	database := &DB{conn: db}
//
//	// Define the context
//	ctx := context.Background()
//
//	// Define the test scenarios
//	scenarios := []struct {
//		name     string
//		UserID   string
//		mockFunc func()
//		expected userentity.User
//		err      error
//	}{
//		{
//			name:   "success case - user found",
//			UserID: "123",
//			mockFunc: func() {
//				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
//					AddRow("123", "John", "Doe")
//				mock.ExpectQuery(`SELECT * FROM users WHERE id = ?`).
//					WithArgs("123").
//					WillReturnRows(rows)
//			},
//			expected: userentity.User{ID: "123", FirstName: "John", LastName: "Doe"},
//			err:      nil,
//		},
//		{
//			name:   "error case - user not found",
//			UserID: "999",
//			mockFunc: func() {
//				mock.ExpectQuery(`SELECT * FROM users WHERE id = ?`).
//					WithArgs("999").
//					WillReturnError(sql.ErrNoRows)
//			},
//			expected: userentity.User{},
//			err: richerror.New("mysqluser.GetUserByID").
//				WithErr(sql.ErrNoRows).
//				WithMessage(errmsg.ErrorMsgNotFound).
//				WithKind(richerror.KindNotFound),
//		},
//		{
//			name:   "error case - scan error",
//			UserID: "123",
//			mockFunc: func() {
//				rows := sqlmock.NewRows([]string{"id", "first_name", "last_name"}).
//					AddRow(nil, nil, nil)
//				mock.ExpectQuery(`SELECT * FROM users WHERE id = ?`).
//					WithArgs("123").
//					WillReturnRows(rows)
//			},
//			expected: userentity.User{},
//			err: richerror.New("mysqluser.GetUserByID").
//				WithErr(errors.New("scan error")).
//				WithMessage(errmsg.ErrorMsgCantScanQueryResult).
//				WithKind(richerror.KindUnexpected),
//		},
//	}
//
//	// Run the test scenarios
//	for _, scenario := range scenarios {
//		t.Run(scenario.name, func(t *testing.T) {
//			// Set up mock expectations
//			scenario.mockFunc()
//
//			// Call the GetUserByID method
//			user, err := database.GetUserByID(ctx, scenario.UserID)
//
//			// Assert the response
//			assert.Equal(t, scenario.expected, user)
//			if scenario.err != nil {
//				assert.Error(t, err)
//				assert.Equal(t, scenario.err.Error(), err.Error())
//			} else {
//				assert.NoError(t, err)
//			}
//
//			// Ensure all expectations were met
//			if err := mock.ExpectationsWereMet(); err != nil {
//				t.Errorf("there were unfulfilled expectations: %s", err)
//			}
//		})
//	}
//}
