package clivault_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClivault(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Clivault Suite")
}
