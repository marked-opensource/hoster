package utils_test

import (
	"fmt"
	"hoster/hoster/utils"
	"io/ioutil"
	"os"
	"testing"
)

func SetupFileUtils() (fUtils utils.IFileUtils, file *os.File, cleanup func()) {
	file, err := ioutil.TempFile(os.TempDir(), "*.sample.etc.hosts")
	if err != nil {
		panic(err)
	}

	fUtils = utils.NewFileUtils(file.Name())

	cleanup = func() {
		err := os.Remove(file.Name())
		if err != nil {
			panic(err)
		}
	}
	return
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
