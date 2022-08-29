package errors_test

import (
	"fmt"
	"testing"

	"github.com/mudi25/monorepo-ci-cd/backend/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestNew(t *testing.T) {
	e := errors.New("internal", nil)
	assert.Equal(t, e.Reason, "err is nil")
	e = errors.New("internal", fmt.Errorf("db err"))
	assert.Equal(t, e.Reason, "db err")
	e1 := errors.New("invalid-argument", e)
	assert.Equal(t, e1, e)
}

func TestFromError(t *testing.T) {
	status, reason := errors.FromError(nil)
	assert.Equal(t, status, "success")
	assert.Equal(t, reason, "")
	// status, reason = errors.FromError(errors.New("internal", fmt.Errorf("db err")))
	// assert.Equal(t, status, "internal")
	// assert.Equal(t, reason, "db err")
	// status, reason = errors.FromError(fmt.Errorf("invlid error"))
	// assert.Equal(t, status, "unknown")
	// assert.Equal(t, reason, "invlid error")
}
func TestError(t *testing.T) {
	err := errors.New("invalid-argument", fmt.Errorf("invalid argument"))
	assert.Equal(t, err.Error(), "invalid-argument: invalid argument")
	err = nil
	assert.Equal(t, err.Error(), "api error is nil")

}
