package reconcilefunc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReconcilefunc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reconcilefunc Suite")
}
