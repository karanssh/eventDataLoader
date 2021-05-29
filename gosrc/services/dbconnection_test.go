package services

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnectToDatabase(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "default test case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConnectToDatabase()
		})
	}
}

func TestEstablishConnection(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "default test case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EstablishConnection()
		})
	}
}
