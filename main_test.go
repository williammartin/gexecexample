package main_test

import (
	"os/exec"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

func TestRun(t *testing.T) {
	g := NewGomegaWithT(t)

	path, err := gexec.Build("main.go")
	g.Expect(err).NotTo(HaveOccurred())

	command := exec.Command(path)

	stdio := gbytes.NewBuffer()
	session, err := gexec.Start(command, stdio, stdio)
	g.Expect(err).NotTo(HaveOccurred())

	g.Eventually(session).Should(gexec.Exit(0))
	g.Eventually(session).Should(gbytes.Say("Hello World"))

	gexec.CleanupBuildArtifacts()
}
