package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/solo-io/build/pkg/constants"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/clicore"
)

var _ = Describe("build cli", func() {

	var (
		initialEnv   = []string{}
		confFilename = "../../examples/sample-solo-project.yaml"
	)

	var _ = BeforeEach(func() {
		initialEnv = os.Environ()
	})
	var _ = AfterEach(func() {
		applyEnv(initialEnv)
	})

	Context("basic args and flags", func() {
		type TestCase struct {
			description    string
			args           string
			buildId        string
			taggedVersion  string
			configFileName string
			cobraOut       string
			cobraErr       string
			consoleLogOut  string
			consoleLogErr  string
		}
		parseEnvTestCases := []TestCase{{
			description:    "should indicate true release",
			args:           "parse-env release",
			buildId:        "1234",
			taggedVersion:  "v1.2.3",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  constants.PrintEnvTrue + "\n",
			consoleLogErr:  "",
		}, {
			description:    "should get release version",
			args:           "parse-env image-tag",
			buildId:        "1234",
			taggedVersion:  "v1.2.3",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  "1.2.3\n",
			consoleLogErr:  "",
		}, {
			description:    "should get non-release version",
			args:           "parse-env image-tag",
			buildId:        "1234",
			taggedVersion:  "",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  "1234\n",
			consoleLogErr:  "",
		}, {
			description:    "should get container prefix for release",
			args:           "parse-env container-prefix",
			buildId:        "1234",
			taggedVersion:  "v1.2.3",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  "quay.io/solo-io\n",
			consoleLogErr:  "",
		}, {
			description:    "should get container prefix for test",
			args:           "parse-env container-prefix",
			buildId:        "1234",
			taggedVersion:  "",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  "gcr.io/solo-public-1010\n",
			consoleLogErr:  "",
		}, {
			description:    "should get version for release",
			args:           "parse-env version",
			buildId:        "1234",
			taggedVersion:  "v1.2.3",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  "1.2.3\n",
			consoleLogErr:  "",
		}, {
			description:    "should get version for test",
			args:           "parse-env version",
			buildId:        "1234",
			taggedVersion:  "",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  "1234\n",
			consoleLogErr:  "",
		}, {
			description:    "should get helm repo for release",
			args:           "parse-env helm-repo",
			buildId:        "1234",
			taggedVersion:  "v1.2.3",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  "gs://solo-helm/\n",
			consoleLogErr:  "",
		}, {
			description:    "should get helm repo for test",
			args:           "parse-env helm-repo",
			buildId:        "1234",
			taggedVersion:  "",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  "gs://solo-helm-test/\n",
			consoleLogErr:  "",
		}, {
			description:    "should validate without error",
			args:           "validate-operating-parameters FALSE 1234 gcr.io/solo-public-1010 1234 gs://solo-helm-test/",
			buildId:        "1234",
			taggedVersion:  "",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "",
			consoleLogOut:  "build parameters are valid\n",
			consoleLogErr:  "",
		}, {
			description:    "should validate with omission error",
			args:           "validate-operating-parameters FALSE 1234 gcr.io/solo-public-1010",
			buildId:        "1234",
			taggedVersion:  "",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "Error: did not receive the expected computed variables: expected 5 arguments, received 3",
			consoleLogOut:  "",
			consoleLogErr:  "",
		}, {
			description:    "should validate with value error",
			args:           "validate-operating-parameters FALSE 1234 gcr.io/solo-public-1010 454545 gs://solo-helm-test/",
			buildId:        "1234",
			taggedVersion:  "",
			configFileName: confFilename,
			cobraOut:       "",
			cobraErr:       "Error: did not receive the expected computed variables: image tag wants: 1234, got: 454545",
			consoleLogOut:  "",
			consoleLogErr:  "",
		}}
		It("should handle parse-env cases", func() {
			for _, tc := range parseEnvTestCases {
				By(fmt.Sprintf("Test case: %s, args: %s", tc.description, tc.args))
				Expect(os.Setenv(constants.EnvBuildId, tc.buildId)).NotTo(HaveOccurred())
				Expect(os.Setenv(constants.EnvTagVersion, tc.taggedVersion)).NotTo(HaveOccurred())
				Expect(os.Setenv(constants.EnvVarConfigFileName, tc.configFileName)).NotTo(HaveOccurred())
				out := appWithLoggerOutput(tc.args)
				Expect(out.CobraStdout).To(Equal(tc.cobraOut))
				Expect(out.CobraStderr).To(ContainSubstring(tc.cobraErr))
				Expect(out.LoggerConsoleStout).To(Equal(tc.consoleLogOut))
				Expect(out.LoggerConsoleStderr).To(Equal(tc.consoleLogErr))
			}
		})
		It("should return help messages without error", func() {
			_, _, err := appWithSimpleOutput("-h")
			Expect(err).NotTo(HaveOccurred())
			_, _, err = appWithSimpleOutput("help")
			Expect(err).NotTo(HaveOccurred())
			_, _, err = appWithSimpleOutput("--help")
			Expect(err).NotTo(HaveOccurred())
		})
	})
})

func appWithSimpleOutput(args string) (string, string, error) {
	co := appWithLoggerOutput(args)
	return co.CobraStdout, co.CobraStderr, nil
}

func appWithLoggerOutput(args string) clicore.CliOutput {
	cliOutput, err := AppConfig.RunForTest(args)
	Expect(err).NotTo(HaveOccurred())
	return cliOutput
}

func applyEnv(evs []string) {
	for _, ev := range evs {
		kv := strings.SplitN(ev, "=", 2)
		Expect(len(kv)).To(Equal(2))
		Expect(os.Setenv(kv[0], kv[1])).NotTo(HaveOccurred())
	}
}
