package main

import (
	"nfa"
	"nfa/gen_code"
	"os"
)

func main() {
	var DigitalTransCondict = nfa.TransitCond{}
	for i := '0'; i <= '9'; i++ {
		DigitalTransCondict = DigitalTransCondict.Allow(byte(i))
	}
	nfa.AddExtend('d', "{0-9}", DigitalTransCondict)

	var AlphabetTransCondict = nfa.TransitCond{}
	for i := 'a'; i <= 'z'; i++ {
		AlphabetTransCondict = AlphabetTransCondict.Allow(byte(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		AlphabetTransCondict = AlphabetTransCondict.Allow(byte(i))
	}
	nfa.AddExtend('a', "{a-z,A-Z}", AlphabetTransCondict)

	var AnyTransCondict = nfa.TransitCond{}
	for i := 0; i <= 255; i++ {
		AnyTransCondict = AnyTransCondict.Allow(byte(i))
	}
	nfa.AddExtend('*', "{any}", AnyTransCondict)

	code := gen_code.GenCodeFromStr("number", `((/ |-#)(((/d/d*)#(/ |((b|B)#)|((s|S)#)|((l|L)#)))|((/d/d*./d*)|(/d*./d/d*))#)(/ |((e|E)(/ |-#)(/d/d*)#))(/ |((f|F)#)|((d|D)#)))|(true#)|(false#)`)
	os.WriteFile("number_core.go", []byte(code), 0755)

	var WhiteSpaceTransCondict = nfa.TransitCond{}
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\t')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\n')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\v')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\f')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\r')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow(' ')
	nfa.AddExtend('b', "{whiteSpace}", WhiteSpaceTransCondict)

	code = gen_code.GenCodeFromStr("whiteSpace", `/b*`)
	os.WriteFile("white_space_core.go", []byte(code), 0755)

	code = gen_code.GenCodeFromStr("leftContainer", `("#)|('#)|({#)|([#(/ |(B;#)|(I;#)|(L;#)))`)
	os.WriteFile("leftContainer.go", []byte(code), 0755)

	var StringNCondict = nfa.TransitCond{}
	for i := 128; i <= 255; i++ {
		StringNCondict = StringNCondict.Allow(byte(i))
	}
	for i := 'a'; i <= 'z'; i++ {
		StringNCondict = StringNCondict.Allow(byte(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		StringNCondict = StringNCondict.Allow(byte(i))
	}
	for i := '0'; i <= '9'; i++ {
		StringNCondict = StringNCondict.Allow(byte(i))
	}
	StringNCondict = StringNCondict.Allow('+')
	StringNCondict = StringNCondict.Allow('-')
	StringNCondict = StringNCondict.Allow('.')
	StringNCondict = StringNCondict.Allow('_')
	nfa.AddExtend('n', "{stringN}", StringNCondict)
	code = gen_code.GenCodeFromStr("unwrapString", `/n*`)
	os.WriteFile("unwrapped_string.go", []byte(code), 0755)
}
