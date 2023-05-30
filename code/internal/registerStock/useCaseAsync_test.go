package registerStock

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"marian.com/interesante-go/code/internal"
	"marian.com/interesante-go/code/test/platform/storagemock"
	"testing"
)

// Parameterized example
func TestUsecase_Execute(t *testing.T) {
	entity, _ := internal.NewEntity(10, "Pant")

	tests := map[string]struct {
		input internal.Entity
		want  string
		err   error
	}{
		"Valid case": {input: entity, want: "Pant", err: nil},
		"Not found":  {input: entity, want: "Pant", err: errors.New("error")},
	}

	//repo := repository.NewEntityRepository()
	repo := new(storagemock.MyMockedRepo)
	repo.On("GetEntities").Return([]internal.Entity{entity}, nil)
	repo.On("SaveEntity", mock.Anything).Return([]internal.Entity{entity}, nil)

	useCase := NewUseCase(repo)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			result, err := useCase.Execute(context.Background(), tc.input)

			if err != nil && tc.err == nil {
				t.Fatalf("Not expected any errors and got %v", err)
			}

			if err == nil && tc.err != nil {
				//t.Errorf("Expected an error and got nil")
				//t.Fatalf("Expected an error and got nil")
			}

			if result != tc.want {
				t.Fatalf("expected %v, got: %v", tc.want, result)
			}

			assert.True(t, true, "True is true!")
		})
	}
}

func Test_UseCase_Execute_RepositoryError(t *testing.T) {
	entity, err := internal.NewEntity(10, "Pant")

	repoMock := new(storagemock.MyMockedRepo)
	repoMock.On("GetEntities").Return([]internal.Entity{entity}, errors.New("Unexpected error"))

	useCase := NewUseCase(repoMock)

	_, err = useCase.Execute(context.Background(), entity)

	repoMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_UseCase_Execute_Succeed(t *testing.T) {
	entity, err := internal.NewEntity(10, "Pant")
	require.NoError(t, err)

	repoMock := new(storagemock.MyMockedRepo)
	repoMock.On("GetEntities").Return([]internal.Entity{entity}, nil)
	repoMock.On("SaveEntity", mock.Anything).Return(nil)

	courseService := NewUseCase(repoMock)

	_, err = courseService.Execute(context.Background(), entity)

	repoMock.AssertExpectations(t)
	assert.NoError(t, err)
}
