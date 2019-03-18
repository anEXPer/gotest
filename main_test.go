package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	It("builds", func() {
		_, err := gexec.Build("github.com/anEXPer/gotest")
		Expect(err).NotTo(HaveOccurred())
	})
})
