package npminstall_test

import (
	"os"
	"slices"
	"testing"

	"github.com/onsi/gomega/format"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitNPMInstall(t *testing.T) {
	format.MaxLength = 0

	suite := spec.New("npm-install", spec.Report(report.Terminal{}))
	suite("Build", testBuild)
	suite("BuildProcessResolver", testBuildProcessResolver)
	suite("CIBuildProcess", testCIBuildProcess)
	suite("Detect", testDetect)
	suite("Environment", testEnvironment)
	suite("InstallBuildProcess", testInstallBuildProcess)
	suite("LinkedModuleResolver", testLinkedModuleResolver)
	suite("Linker", testLinker)
	suite("PackageManangerConfigurationManager", testPackageManagerConfigurationManager)
	suite("PruneBuildProcess", testPruneBuildProcess)
	suite("RebuildBuildProcess", testRebuildBuildProcess)
	suite("UpdateNpmCacheLayer", testUpdateNpmCache)
	suite.Run(t)
}

// addedEnv returns the entries a process adds on top of the inherited
// environment, so failing assertions don't print the CI environment, which
// can contain secrets.
func addedEnv(env []string) []string {
	inherited := os.Environ()
	if len(env) < len(inherited) || !slices.Equal(env[:len(inherited)], inherited) {
		return []string{"<inherited environment was not passed through unchanged>"}
	}

	return env[len(inherited):]
}
