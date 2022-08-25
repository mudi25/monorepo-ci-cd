package logger_test

import (
	"fmt"
	"testing"

	"github.com/mudi25/monorepo-ci-cd/pkg/logger"
)

func TestMain(m *testing.M) {
	m.Run()
}
func TestError(t *testing.T) {
	logger.Error(fmt.Errorf("invalid argument"))
	logger.Error(nil)
}
func TestInfo(t *testing.T) {
	logger.Info("hello")
}
