package unused

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/adapter/call"
	"github.com/phodal/coca/core/adapter/identifier"
	"sync"
	"testing"
)

func TestRenameMethodApp(t *testing.T) {
	g := NewGomegaWithT(t)

	var wg sync.WaitGroup

	codePath := "../../../../_fixtures/refactor/unused"
	configPath := "../../../../_fixtures/refactor/rename.config"
	identifierApp := new(identifier.JavaIdentifierApp)
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	wg.Add(1)
	callApp := call.NewJavaCallApp()
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	wg.Add(1)
	go func() {

		RenameMethodApp(callNodes, configPath).Start()
		defer wg.Done()

		newnodes := callApp.AnalysisPath(codePath, classes, identifiers)
		g.Expect(newnodes[0].Methods[0].Name).To(Equal("demoA"))

	}()

	wg.Add(1)
	go func() {
		wg.Wait()

		configPath2 := "../../../../_fixtures/refactor/rename_back.config"
		RenameMethodApp(callNodes, configPath2).Start()
		defer wg.Done()

		renameBackNodes := callApp.AnalysisPath(codePath, classes, identifiers)
		g.Expect(renameBackNodes[0].Methods[0].Name).To(Equal("demo"))

		wg.Wait()
	}()

}
