package goodtools

import "github.com/prataprc/goparsec"


type Translation struct {

	Language string
	Details string
}

type Descriptor struct {
	Title string
	Country string
	Version float64

}

// [ ParsecNode , ... ] => ParsecNode
func one2one(ns []parsec.ParsecNode) parsec.ParsecNode {
	if ns == nil || len(ns) == 0 {
		return nil
	}
	return ns[0]
}

func multiLanguageFlag() parsec.Parser {
	openParens := parsec.Token(`\(`, "PO")
	closeParens := parsec.Token(`\)`, "PC")

	flag := parsec.Token(`M`, "MULTILANGUAGE PREFIX")
	version := parsec.Token(`[0-9]+`, "MULTILANGUAGE")

	versionNode := func(ns []parsec.ParsecNode) parsec.ParsecNode {
		return ns[2]
	}

	return parsec.And(versionNode, openParens, flag, version, closeParens)
}

// Recognize a single translation flag
func translationFlag() parsec.Parser {
	openSquare := parsec.Token(`\[`, "SO")
	closeSquare := parsec.Token(`\]`, "SC")

	flags := parsec.OrdChoice(one2one,
		parsec.Token(`T\+`, "CURRENT"),
		parsec.Token(`T\-`, "OBSOLETE"),
	)

	languages := parsec.OrdChoice(one2one,
		parsec.Token(`Eng`, "ENGLISH"),
		parsec.Token(`Ger`, "GERMAN"),
		parsec.Token(`Fer`, "FRENCH"),
	)

	extra := parsec.Token(`[^\[\]]+`, "EXTRA")

	// Drop square brackets
	translationNode := func(ns []parsec.ParsecNode) parsec.ParsecNode {
		return ns[1:4]
	}

	parser := parsec.And(translationNode,
		openSquare,
		flags,
		languages,
		parsec.Maybe(one2one, extra),
		closeSquare,
	)

	return parser
}

// Recognize a single country flag
func countryFlag() parsec.Parser {
	openParens := parsec.Token(`\(`, "PO")
	closeParens := parsec.Token(`\)`, "PC")

	countries := parsec.OrdChoice(one2one,
		parsec.Token(`A`, "AUSTRALIA"),
		parsec.Token(`AS`, "ASIA"),
		parsec.Token(`B`, "BRAZIL"),
		parsec.Token(`C`, "CANADA"),
		parsec.Token(`CH`, "CHINA"),
		parsec.Token(`D`, "NETHERLANDS"),
		parsec.Token(`E`, "EUROPE"),
		parsec.Token(`F`, "FRANCE"),
		parsec.Token(`G`, "GERMANY"),
		parsec.Token(`GR`, "GREECE"),
		parsec.Token(`HK`, "HONG KONG"),
		parsec.Token(`I`, "ITALY"),
		parsec.Token(`J`, "JAPAN"),
		parsec.Token(`K`, "KOREA"),
		parsec.Token(`NL`, "NETHERLANDS"),
		parsec.Token(`NO`, "NORWAY"),
		parsec.Token(`R`, "RUSSIA"),
		parsec.Token(`S`, "SPAIN"),
		parsec.Token(`SW`, "SWEDEN"),
		parsec.Token(`U`, "UNITED STATES"),
		parsec.Token(`UK`, "UNITED KINGDOM"),
		parsec.Token(`W`, "WORLD"),
		parsec.Token(`UNL`, "UNLICENSED"),
		parsec.Token(`PD`, "PUBLIC DOMAIN"),
		parsec.Token(`UNK`, "UNKNOWN COUNTRY"),
	)

	countryNode := func(ns []parsec.ParsecNode) parsec.ParsecNode {
		return ns[1]
	}

	return parsec.And(countryNode, openParens, countries, closeParens)
}

// Recognize a single version flag
func versionFlag() parsec.Parser {
	openParens := parsec.Token(`\(`, "PO")
	closeParens := parsec.Token(`\)`, "PC")

	flag := parsec.Token(`V`, "VERSION PREFIX")
	version := parsec.Token(`[0-9]\.[0-9]`, "VERSION")

	versionNode := func(ns []parsec.ParsecNode) parsec.ParsecNode {
		return ns[2]
	}

	return parsec.And(versionNode, openParens, flag, version, closeParens)
}

// Recognize a single standard flag
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

	// [ "[", "flag", "]" ] => "flag"
	standardFlagNode := func(ns []parsec.ParsecNode) parsec.ParsecNode {
		return ns[1]
	}

	return parsec.And(standardFlagNode, openSquare, flags, closeSquare)
}

func zeroOrMoreFlags() parsec.Parser {
	return parsec.Many(nil,
		parsec.OrdChoice(one2one,
			countryFlag(),
			versionFlag(),
			multiLanguageFlag(),
			standardFlag(),
			translationFlag(),
		),
	)
}

func title() parsec.Parser {
	return parsec.TokenExact(`[^\[\]\(\)]+`, "TITLE")
}

func extension() parsec.Parser {
	return parsec.TokenExact(`\.[0-9a-zA-Z]+`, "EXT")
}

func goodNameParser() parsec.Parser {
	return parsec.And(nil,
		title(),
		zeroOrMoreFlags(),
		extension(),
	)
}