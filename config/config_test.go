package config

import (
    "os"
    "testing"

    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
)

// Mock for godotenv package
type MockGodotenv struct {
    mock.Mock
}

func (m *MockGodotenv) Load() error {
    args := m.Called()
    return args.Error(0)
}

func TestLoadEnv(t *testing.T) {

    mockGodotenv := new(MockGodotenv)

	t.Run(".env file found", func(t *testing.T) {
		mockGodotenv.On("Load").Return(nil)
		LoadEnv(func() error {
			return mockGodotenv.Load()
		})
		mockGodotenv.AssertCalled(t, "Load")

	})


	t.Run(".env file not found", func(t *testing.T) {
		mockGodotenv.On("Load").Return(os.ErrNotExist)
		LoadEnv(func() error {
			return mockGodotenv.Load()
		})
		mockGodotenv.AssertCalled(t, "Load")
	})

}

func TestGetPort(t *testing.T) {
    
	t.Run("PORT environment variable is set", func(t *testing.T) {
		os.Setenv("PORT", "9090")
		port := GetPort()
		require.Equal(t, "9090", port)
	})

	t.Run("PORT environment variable is not set", func(t *testing.T) {
		os.Unsetenv("PORT")
		port := GetPort()
		require.Equal(t, "8080", port)
	})
}