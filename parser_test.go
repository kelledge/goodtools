package goodtools

import (
	"testing"
	"github.com/prataprc/goparsec"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func printNodes(ns parsec.ParsecNode) {


	t, ok := ns.([]parsec.ParsecNode)
	if ok {
		fmt.Println("RECURSE")
		for n := range t {
			printNodes(n)
		}
	}

	n, ok := ns.(*parsec.Terminal)
	if ok {
		fmt.Printf("'%s'", n)
	}

}

func TestStandardFlag(t *testing.T) {
	s := parsec.NewScanner([]byte("[!]"))

	parser := standardFlag()
	nodes, _ := parser(s)

	assert.Equal(t, nodes, nodes, "not equal")
}

func TestTranslationFlag(t *testing.T) {

	type testResult struct {

	}
	// "[T+Ger90%_TranX]", "[T+Fre100V2_generation9]", "[T-Fre100V1_generation9]", "[T+Eng]", "[T+Eng1.0_DKK]"
	s := parsec.NewScanner([]byte("[T+Ger]"))

	parser := translationFlag()
	nodes, _ := parser(s)

	assert.Equal(t, nodes, nodes, "not equal")
}

func TestVersionFlag(t *testing.T) {
	s := parsec.NewScanner([]byte("(V1.1)"))

	parser := versionFlag()
	nodes, _ := parser(s)

	assert.Equal(t, nodes, nodes, "not equal")
}

func TestMultiLanguageFlag(t *testing.T) {
	s := parsec.NewScanner([]byte("(M3)"))

	parser := multiLanguageFlag()
	nodes, _ := parser(s)

	assert.Equal(t, nodes, nodes, "not equal")
}

func TestZeroOrMoreFlags(t *testing.T) {
	s := parsec.NewScanner([]byte("(U) (V1.0) [!][T+Ger]"))

	parser := zeroOrMoreFlags()
	nodes, _ := parser(s)

	assert.Equal(t, nodes, nodes, "not equal")
}

func TestGoodNameParser(t *testing.T) {
	s := parsec.NewScanner([]byte("Final Fantasy III (U) (V1.1) [T+Ger1.00_Star].smc"))

	parser := goodNameParser()
	nodes, _ := parser(s)

	//fmt.Printf("%s\n", nodes)

	printNodes(nodes)

	assert.Equal(t, nodes, nil, "not equal")
}