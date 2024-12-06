package tests

import (
	"testing"

	"github.com/d-alejandro/go-code-examples/e2e/tests/suites/ordersuite"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

func TestRunSuites(t *testing.T) {
	t.Parallel()

	suite.RunSuite(t, new(ordersuite.OrderRunnerSuite))
}
