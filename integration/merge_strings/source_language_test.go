package merge_strings_test

import (
	"os"
	"path/filepath"

	. "github.com/maximilien/i18n4cf/integration/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("merge-strings -d dirName -source-lanuage sourceLanguage", func() {
	var (
		rootPath          string
		fixturesPath      string
		inputFilesPath    string
		expectedFilesPath string
	)

	BeforeEach(func() {
		dir, err := os.Getwd()
		Ω(err).ShouldNot(HaveOccurred())
		rootPath = filepath.Join(dir, "..", "..")

		fixturesPath = filepath.Join("..", "..", "test_fixtures", "merge_strings", "source_language")
		inputFilesPath = filepath.Join(fixturesPath, "input_files")
		expectedFilesPath = filepath.Join(fixturesPath, "expected_output")
	})

	Context("can combine multiple language files", func() {
		BeforeEach(func() {
			session := Runi18n("-merge-strings", "-v", "-d", filepath.Join(inputFilesPath), "-source-language", "fr")
			Ω(session.ExitCode()).Should(Equal(0))
		})

		AfterEach(func() {
			RemoveAllFiles(
				GetFilePath(inputFilesPath, "fr.all.json"),
			)
		})

		It("fr.all.json contains translations from both files", func() {
			CompareExpectedToGeneratedTraslationJson(
				GetFilePath(expectedFilesPath, "fr.all.json"),
				GetFilePath(inputFilesPath, "fr.all.json"),
			)
		})
	})

})
