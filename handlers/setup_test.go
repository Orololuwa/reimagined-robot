package handlers

import (
	"os"
	"testing"

	"github.com/orololuwa/reimagined-robot/repository"
)


func NewTestingHandler() {
	r := &Repository{
		User: repository.NewUserTestingDBRepo(),
	}
	Repo = r
}

func TestMain(m *testing.M){
	NewTestingHandler()
	os.Exit(m.Run())
}