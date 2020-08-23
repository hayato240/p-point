package usecase

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/hayato240/p-point/domain"
)

type mockUserRepository struct {
	UserRepository
	mockedAdd             func(u domain.User) (int, error)
	mockedFindById        func(int) (domain.User, error)
	mockedPoints          func(u domain.User) (int, error)
	mockedUpdateAmount    func(int, int) error // TODO(Sho): ここにmock用のメソッドを書き込む。
	mockedAddPointHistory func(int, int) error // TODO(Sho): ここにmock用のメソッドを書き込む。
}

func (m *mockUserRepository) Add(u domain.User) (int, error) {
	return m.mockedAdd(u)
}

func (m *mockUserRepository) FindById(i int) (domain.User, error) {
	return m.mockedFindById(i)
}

func (m *mockUserRepository) AddPoints(u domain.User) (int, error) {
	return m.mockedPoints(u)
}

func (m *mockUserRepository) UpdateAmount(tx *sql.Tx, newAmount int, userID int) error {
	return m.mockedUpdateAmount(newAmount, userID)
}

func (m *mockUserRepository) AddPointHistory(tx *sql.Tx, userID int, AddedAmount int) error {
	return m.mockedAddPointHistory(userID, AddedAmount)
}

func TestUserInteractor_Add(t *testing.T) {
	reqUser := domain.User{
		ID:     1,
		Amount: 100,
	}
	tests := []struct {
		name       string
		interactor *UserInteractor
		want       domain.User
		wantErr    bool
	}{
		{
			"success to add user",
			&UserInteractor{
				UserRepository: &mockUserRepository{
					mockedAdd: func(u domain.User) (int, error) {
						return reqUser.ID, nil
					},
					mockedFindById: func(i int) (domain.User, error) {
						return reqUser, nil
					},
				},
			},
			reqUser,
			false,
		},
		{
			"failed to add user",
			&UserInteractor{
				UserRepository: &mockUserRepository{
					mockedAdd: func(u domain.User) (int, error) {
						return 0, errors.New("add method failed")
					},
					mockedFindById: func(i int) (domain.User, error) {
						return reqUser, nil
					},
				},
			},
			domain.User{},
			true,
		},
		{
			"faild on FindById",
			&UserInteractor{
				UserRepository: &mockUserRepository{
					mockedAdd: func(u domain.User) (int, error) {
						return reqUser.ID, nil
					},
					mockedFindById: func(i int) (domain.User, error) {
						return domain.User{}, errors.New("findById failed")
					},
				},
			},
			domain.User{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := tt.interactor.Add(reqUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() got err = %v, wantErr = %v", err, tt.wantErr)
			}
			if (err == nil) && !reflect.DeepEqual(reqUser, gotUser) {
				t.Errorf("want: %+v", reqUser)
				t.Errorf("got %+v", gotUser)
			}
		})
	}
}

func TestUserInteractor_Show(t *testing.T) {
	reqUser := domain.User{
		ID:     1,
		Amount: 100,
	}
	tests := []struct {
		name       string
		interactor *UserInteractor
		want       domain.User
		wantErr    bool
	}{
		{
			"success to show user",
			&UserInteractor{
				UserRepository: &mockUserRepository{
					mockedAdd: func(u domain.User) (int, error) {
						return reqUser.ID, nil
					},
					mockedFindById: func(i int) (domain.User, error) {
						return reqUser, nil
					},
				},
			},
			reqUser,
			false,
		},
		{
			"fail to show user",
			&UserInteractor{
				UserRepository: &mockUserRepository{
					mockedFindById: func(i int) (domain.User, error) {
						return domain.User{ID: 0, Amount: 0}, errors.New("no such user")
					},
				},
			},
			domain.User{ID: 0, Amount: 0},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := tt.interactor.Show(reqUser.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() got err = %v, wantErr = %v", err, tt.wantErr)
			}
			if (err == nil) && !reflect.DeepEqual(reqUser, gotUser) {
				t.Errorf("want: %+v", reqUser)
				t.Errorf("got %+v", gotUser)
			}
		})
	}

}
func TestUserInteractor_AddPoints(t *testing.T) {
	reqUser := domain.User{
		ID:     1,
		Amount: 50,
	}
	resultUser := domain.User{
		ID:     1,
		Amount: 100,
	}
	tests := []struct {
		name       string
		interactor *UserInteractor
		want       domain.User
		wantErr    bool
	}{
		{
			"success to add points",
			&UserInteractor{
				UserRepository: &mockUserRepository{
					mockedAdd: func(u domain.User) (int, error) {
						return reqUser.ID, nil
					},
					mockedFindById: func(i int) (domain.User, error) {
						return reqUser, nil
					},
					mockedPoints: func(u domain.User) (int, error) { // TODO(Sho): ここにUpdateAmountMethodとAddPointHistoryメソッドを入れる。
						reqUser.Amount = reqUser.Amount + u.Amount
						return reqUser.ID, nil
					},
				},
			},
			resultUser,
			false,
		},
		{
			"failed to add points",
			&UserInteractor{
				UserRepository: &mockUserRepository{
					mockedAdd: func(u domain.User) (int, error) {
						return reqUser.ID, nil
					},
					mockedFindById: func(i int) (domain.User, error) {
						return reqUser, nil
					},
					mockedPoints: func(u domain.User) (int, error) {
						return reqUser.ID, errors.New("failed adding points")
					},
				},
			},
			reqUser,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := tt.interactor.Add(reqUser)
			pointResult, err := tt.interactor.AddPoints(domain.User{ID: gotUser.ID, Amount: 50})
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() got err = %v, wantErr = %v", err, tt.wantErr)
			}
			if (err == nil) && !reflect.DeepEqual(resultUser, pointResult) {
				t.Errorf("want: %+v", resultUser)
				t.Errorf("got %+v", pointResult)
			}
		})
	}

}
