package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMd5Service(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Md5Service Suite")
}
