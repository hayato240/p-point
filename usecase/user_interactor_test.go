package usecase

import (
	"github.com/p-point/domain"
	"reflect"
	"testing"
)

type mockUserRepository struct {
	UserRepository
	mockedAdd func(u domain.User) (int, error)
	mockedFindById func(int) (domain.User, error)
}

func (m *mockUserRepository) Add(u domain.User) (int, error) {
	return m.mockedAdd(u)
}

func (m *mockUserRepository) FindById(i int) (domain.User, error) {
	return m.mockedFindById(i)
}

func TestUserInteractor_Add(t *testing.T){
	reqUser := domain.User{
		ID:     1,
		Amount: 100,
	}
	tests := []struct{
		name string
		interactor *UserInteractor
		want domain.User
		wantErr bool
	}{
		{
			"success to add user",
			&UserInteractor{
				UserRepository: &mockUserRepository{
					mockedAdd: func(u domain.User) (int, error){
						return reqUser.ID, nil
					},
					mockedFindById: func(i int) (domain.User, error){
						return reqUser, nil
					},
				},
			},
			reqUser,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			gotUser, err := tt.interactor.Add(reqUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() got err = %v, wantErr = %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(reqUser, gotUser){
				t.Errorf("want: %+v", reqUser)
				t.Errorf("got %+v", gotUser)
			}
		})
	}
}
