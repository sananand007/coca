package call

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/adapter/identifier"
	"github.com/phodal/coca/core/models"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/call"
	identifierApp := new(identifier.JavaIdentifierApp)
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := NewJavaCallApp()
	callNodes := callApp.AnalysisPath(codePath, classes, iNodes)

	g.Expect(len(callNodes)).To(Equal(1))
}

func TestJavaCallListener_EnterConstructorDeclaration(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/suggest/factory"
	callNodes := getCallNodes(codePath)
	g.Expect(len(callNodes[0].Methods)).To(Equal(3))
}

func getCallNodes(codePath string) []models.JClassNode {
	identifierApp := new(identifier.JavaIdentifierApp)
	iNodes := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := NewJavaCallApp()

	callNodes := callApp.AnalysisPath(codePath, classes, iNodes)
	return callNodes
}

func TestLambda_Express(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/lambda"
	callNodes := getCallNodes(codePath)

	methodMap := make(map[string]models.JMethod)
	for _, c := range callNodes[1].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(methodMap["save"].MethodCalls[0].MethodName).To(Equal("of"))
	g.Expect(methodMap["findById"].MethodCalls[3].MethodName).To(Equal("toDomainModel"))
}

func TestInterface(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/grammar/interface"
	callNodes := getCallNodes(codePath)

	g.Expect(true).To(Equal(true))
	methodMap := make(map[string]models.JMethod)
	for _, c := range callNodes[0].Methods {
		methodMap[c.Name] = c
	}

	g.Expect(len(callNodes[0].Methods)).To(Equal(6))
	g.Expect(methodMap["count"].Name).To(Equal("count"))
}