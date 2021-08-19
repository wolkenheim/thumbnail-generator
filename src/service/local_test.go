package service

import (
	"github.com/spf13/afero"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestLocalFileService(t *testing.T){

	// arrange
	filename := "../takk/test.jpg"

	var appFs = afero.NewMemMapFs()
	logger := zaptest.NewLogger(t)
	s := NewLocalFileService(logger.Sugar(), appFs)

	// create file to be deleted
	file, _ := appFs.Create(filename)
	defer file.Close()

	// assert file exists
	_, err := appFs.Open(filename)
	if err != nil {
		t.Error("file should exists")
	}

	// act. test actual function
	s.DeleteFile(filename)

	// asser
	_, fileDoesNotExistErr := appFs.Open(filename)
	if fileDoesNotExistErr == nil {
		t.Error("file should be deleted")
	}

}