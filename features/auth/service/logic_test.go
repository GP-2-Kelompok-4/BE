package service

// import (
// 	"errors"
// 	"testing"

// 	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/auth"
// 	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/mocks"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestLogin(t *testing.T) {
// 	repo := new(mocks.AuthRepository)
// 	t.Run("Success Login")
// }

// func TestCreate(t *testing.T) {
// 	repo := new(mocks.UserRepository)
// 	t.Run("Success Create user", func(t *testing.T) {
// 		inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
// 		inputData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
// 		repo.On("Create", inputRepo).Return(1, nil).Once()
// 		srv := New(repo)
// 		err := srv.Create(inputData)
// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed Create user, duplicate entry", func(t *testing.T) {
// 		inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
// 		inputData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
// 		repo.On("Create", inputRepo).Return(0, errors.New("failed to insert data, error query")).Once()
// 		srv := New(repo)
// 		err := srv.Create(inputData)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "failed to insert data, error query", err.Error())
// 		repo.AssertExpectations(t)
// 	})

// 	/*
// 		dalam testcase ini kita akan melakukan test saat kondisi validationnya error,
// 		sehingga pengkondisian (if) validation akan terpenuhi dan akan menjalankan perintah return.
// 		jadi fungsi Create yang ada di repository tidak akan dijalankan dalam test case ini. sooo kita tidak perlu memanggil mock "Create" repo.On

// 	t.Run("Failed Create user, name empty", func(t *testing.T) {
// 		// inputRepo := user.Core{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
// 		inputData := user.Core{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
// 		// repo.On("Create", inputRepo).Return(0, errors.New("failed to insert data, error query")).Once()
// 		srv := New(repo)
// 		err := srv.Create(inputData)
// 		assert.NotNil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// }
