package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hoster/hoster/tmocks"
	"io/ioutil"
	"os"
	"testing"
)

func SetupFileUtils() (fUtils IFileUtils, file *os.File, cleanup func()) {
	file, err := ioutil.TempFile(os.TempDir(), "*.sample.etc.hosts")
	if err != nil {
		panic(err)
	}

	fUtils = NewFileUtils(file.Name())

	cleanup = func() {
		err := os.Remove(file.Name())
		if err != nil {
			panic(err)
		}
	}
	return
}

const sampleEnvFile = `##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1	localhost
255.255.255.255	broadcasthost
::1             localhost
# Added by Docker Desktop
# To allow the same kube context to work on the host and the container:
127.0.0.1 kubernetes.docker.internal
# End of section`

const sampleEnvFileExistingRules = `##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1	localhost
255.255.255.255	broadcasthost
::1             localhost
# Added by Docker Desktop
# To allow the same kube context to work on the host and the container:
127.0.0.1 kubernetes.docker.internal
# End of section
# Added by Hoster
127.0.0.1 marek2901.github.io
# End of section
`

func TestFileUtils_EnsureHosterBlock(t *testing.T) {
	t.Run("Test empty file", func(t *testing.T) {
		utilsInterface, targetHostsFile, cleanup := SetupFileUtils()
		defer cleanup()
		err := utilsInterface.(*fileUtils).EnsureHosterBlock()
		assert.NoError(t, err)

		fileBytes, err := ioutil.ReadFile(targetHostsFile.Name())
		assert.NoError(t, err)
		assert.Contains(t, string(fileBytes), "# Added by Hoster\n# End of section")
	})

	t.Run("Test already existing non hoster rules file", func(t *testing.T) {
		utilsInterface, targetHostsFile, cleanup := SetupFileUtils()
		_, err := targetHostsFile.WriteString(sampleEnvFile)
		assert.NoError(t, err)
		defer cleanup()
		err = utilsInterface.(*fileUtils).EnsureHosterBlock()
		assert.NoError(t, err)

		fileBytes, err := ioutil.ReadFile(targetHostsFile.Name())
		assert.NoError(t, err)
		assert.Contains(t, string(fileBytes), "# Added by Hoster\n# End of section")
	})

	t.Run("Test already existing hoster rules file", func(t *testing.T) {
		utilsInterface, targetHostsFile, cleanup := SetupFileUtils()
		_, err := targetHostsFile.WriteString(sampleEnvFileExistingRules)
		assert.NoError(t, err)
		defer cleanup()
		err = utilsInterface.(*fileUtils).EnsureHosterBlock()
		assert.NoError(t, err)

		fileBytes, err := ioutil.ReadFile(targetHostsFile.Name())
		assert.NoError(t, err)
		assert.NotContains(t, string(fileBytes), "# Added by Hoster\n# End of section")
	})
}

func TestFileutils_RefreshRules(t *testing.T) {
	t.Run("Injects hakuna matata host rule", func(t *testing.T) {
		tests := map[string]func()(fUtils IFileUtils, file *os.File, cleanup func()){
			"empty": func() (fUtils IFileUtils, file *os.File, cleanup func()) {
				fUtils, file, cleanup = SetupFileUtils()
				return
			},
			"existing block": func() (fUtils IFileUtils, file *os.File, cleanup func()) {
				fUtils, file, cleanup = SetupFileUtils()
				_, err := file.WriteString(sampleEnvFileExistingRules)
				assert.NoError(t, err)
				return
			},
		}

		for testName, setupFunc := range tests {
			t.Run(testName, func(t *testing.T) {
				utilsInterface, targetHostsFile, cleanup := setupFunc()
				defer cleanup()
				err := utilsInterface.(*fileUtils).EnsureHosterBlock()
				assert.NoError(t, err)

				rulesMock := tmocks.EnabledRulesMockImplementation{}
				rulesMock.On("GetAll").Return([]string{"project1.hosfile"}, nil)
				rulesMock.On("ReadRule", "project1.hosfile").Return("127.0.0.1\thakuna.matata.org", nil)
				utilsInterface.(*fileUtils).enabledRules = rulesMock

				err = utilsInterface.RefreshRules()
				assert.NoError(t, err)

				fileBytes, err := ioutil.ReadFile(targetHostsFile.Name())
				assert.NoError(t, err)
				assert.Contains(t, string(fileBytes), "127.0.0.1\thakuna.matata.org")
				assert.NotContains(t, string(fileBytes), "127.0.0.1 marek2901.github.io")
			})
		}
	})
}

func TestFileUtils_InjectRule(t *testing.T) {
	utilsInterface, testFile, cleanup := SetupFileUtils()
	defer cleanup()

	fmt.Printf("%s\n", testFile.Name())

	fmt.Printf("%p\n", utilsInterface.ClearAll)
}

func TestFileUtils_ClearRule(t *testing.T) {
	utilsInterface, testFile, cleanup := SetupFileUtils()
	defer cleanup()

	fmt.Printf("%s\n", testFile.Name())

	fmt.Printf("%p\n", utilsInterface.ClearAll)
}

func TestFileUtils_ClearAll(t *testing.T) {
	utilsInterface, testFile, cleanup := SetupFileUtils()
	defer cleanup()

	fmt.Printf("%s\n", testFile.Name())

	fmt.Printf("%p\n", utilsInterface.ClearAll)
}
