// Initial proof-of-concept
package main

import (
	"fmt"
	"github.com/prataprc/goparsec"
)

func one2one(ns []parsec.ParsecNode) parsec.ParsecNode {
	if ns == nil || len(ns) == 0 {
		return nil
	}
	return ns[0]
}



func standardFlag() parsec.Parser {
	openSquare := parsec.Token(`\[`, "SO")
	closeSquare := parsec.Token(`\]`, "SC")

	flags := parsec.OrdChoice(one2one,
		parsec.Token(`!`, "GOOD"),
		parsec.Token(`a`, "ALTERNATIVE"),
		parsec.Token(`b`, "BAD"),
		parsec.Token(`f`, "FIXED"),
		parsec.Token(`h`, "HACKED"),
		parsec.Token(`o`, "OVERDUMPED"),
		parsec.Token(`p`, "PIRATED"),
		parsec.Token(`t`, "TRAINED"),
		parsec.Token(`!p`, "PENDING"),
	)

	standardFlagNode := func(ns []parsec.ParsecNode) parsec.ParsecNode {
		return ns[1]
	}

	return parsec.And(standardFlagNode, openSquare, flags, closeSquare)
}

func extension() parsec.Parser {
	return parsec.Token(`\.[0-9a-zA-Z]+`, "EXT")

}

func title() parsec.Parser {
	return parsec.TokenExact(`.*?`, "TITLE")
}

func main() {


}

func debugNode(ns []parsec.ParsecNode) parsec.ParsecNode {
	fmt.Printf("Node Length: %d\n", len(ns))
	fmt.Printf("Node Values: %s\n", ns)
	return ns
}
