package cocago

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
	"testing"
)

func Test_DataStructProperty(t *testing.T) {
	g := NewGomegaWithT(t)

	results := ProcessFile("testdata/data_struct_property.code")
	g.Expect(len(results.Members)).To(Equal(1))
	properties := results.DataStructures[0].Properties

	g.Expect(len(properties)).To(Equal(5))
	g.Expect(properties[0].Name).To(Equal("FullName"))
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/data_struct_property.json")).To(Equal(true))
}

func Test_DataStructWithFuncType(t *testing.T) {
	g := NewGomegaWithT(t)

	results := ProcessFile("testdata/struct_with_func.code")
	g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/struct_with_func.json")).To(Equal(true))
}
//
//func Test_DataStructWithFuncDecl(t *testing.T) {
//	g := NewGomegaWithT(t)
//
//	results := ProcessFile("testdata/struct_with_func_decl.code")
//	fmt.Println(results)
//	//g.Expect(cocatest.JSONFileBytesEqual(results, "testdata/struct_with_func_decl.json")).To(Equal(true))
//	g.Expect(true).To(Equal(true))
//}