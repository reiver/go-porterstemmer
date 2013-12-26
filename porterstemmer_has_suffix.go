package porterstemmer



import (
    "testing"
)



func TestHasSuffix(t *testing.T) {

	tests := make([]struct {
		S []byte
		Suffix []byte
		Expected bool
	}, 82)



	i := 0


	tests[i].S         = []byte("ran")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("runner")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("runnar")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("runned")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("runnre")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("er")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("re")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = false
	i++



	tests[i].S         = []byte("ran")
	tests[i].Suffix    = []byte("ER")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("runner")
	tests[i].Suffix    = []byte("ER")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("runnar")
	tests[i].Suffix    = []byte("ER")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("runned")
	tests[i].Suffix    = []byte("ER")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("runnre")
	tests[i].Suffix    = []byte("ER")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("er")
	tests[i].Suffix    = []byte("ER")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("re")
	tests[i].Suffix    = []byte("ER")
	tests[i].Expected  = false
	i++



	tests[i].S         = []byte("")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = false
	i++

	tests[i].S         = []byte("e")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = false
	i++



	tests[i].S         = []byte("caresses")
	tests[i].Suffix    = []byte("sses")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("ponies")
	tests[i].Suffix    = []byte("ies")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("caress")
	tests[i].Suffix    = []byte("ss")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("cats")
	tests[i].Suffix    = []byte("s")
	tests[i].Expected  = true
	i++



	tests[i].S         = []byte("feed")
	tests[i].Suffix    = []byte("eed")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("agreed")
	tests[i].Suffix    = []byte("eed")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("plastered")
	tests[i].Suffix    = []byte("ed")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("bled")
	tests[i].Suffix    = []byte("ed")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("motoring")
	tests[i].Suffix    = []byte("ing")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("sing")
	tests[i].Suffix    = []byte("ing")
	tests[i].Expected  = true
	i++



	tests[i].S         = []byte("conflat")
	tests[i].Suffix    = []byte("at")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("troubl")
	tests[i].Suffix    = []byte("bl")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("siz")
	tests[i].Suffix    = []byte("iz")
	tests[i].Expected  = true
	i++



	tests[i].S         = []byte("happy")
	tests[i].Suffix    = []byte("y")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("sky")
	tests[i].Suffix    = []byte("y")
	tests[i].Expected  = true
	i++



	tests[i].S         = []byte("relational")
	tests[i].Suffix    = []byte("ational")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("conditional")
	tests[i].Suffix    = []byte("tional")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("rational")
	tests[i].Suffix    = []byte("tional")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("valenci")
	tests[i].Suffix    = []byte("enci")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("hesitanci")
	tests[i].Suffix    = []byte("anci")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("digitizer")
	tests[i].Suffix    = []byte("izer")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("conformabli")
	tests[i].Suffix    = []byte("abli")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("radicalli")
	tests[i].Suffix    = []byte("alli")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("differentli")
	tests[i].Suffix    = []byte("entli")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("vileli")
	tests[i].Suffix    = []byte("eli")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("analogousli")
	tests[i].Suffix    = []byte("ousli")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("vietnamization")
	tests[i].Suffix    = []byte("ization")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("predication")
	tests[i].Suffix    = []byte("ation")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("operator")
	tests[i].Suffix    = []byte("ator")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("feudalism")
	tests[i].Suffix    = []byte("alism")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("decisiveness")
	tests[i].Suffix    = []byte("iveness")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("hopefulness")
	tests[i].Suffix    = []byte("fulness")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("callousness")
	tests[i].Suffix    = []byte("ousness")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("formaliti")
	tests[i].Suffix    = []byte("aliti")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("sensitiviti")
	tests[i].Suffix    = []byte("iviti")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("sensibiliti")
	tests[i].Suffix    = []byte("biliti")
	tests[i].Expected  = true
	i++



	tests[i].S         = []byte("triplicate")
	tests[i].Suffix    = []byte("icate")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("formative")
	tests[i].Suffix    = []byte("ative")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("formalize")
	tests[i].Suffix    = []byte("alize")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("electriciti")
	tests[i].Suffix    = []byte("iciti")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("electrical")
	tests[i].Suffix    = []byte("ical")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("hopeful")
	tests[i].Suffix    = []byte("ful")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("goodness")
	tests[i].Suffix    = []byte("ness")
	tests[i].Expected  = true
	i++



	tests[i].S         = []byte("revival")
	tests[i].Suffix    = []byte("al")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("allowance")
	tests[i].Suffix    = []byte("ance")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("inference")
	tests[i].Suffix    = []byte("ence")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("airliner")
	tests[i].Suffix    = []byte("er")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("gyroscopic")
	tests[i].Suffix    = []byte("ic")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("adjustable")
	tests[i].Suffix    = []byte("able")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("defensible")
	tests[i].Suffix    = []byte("ible")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("irritant")
	tests[i].Suffix    = []byte("ant")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("replacement")
	tests[i].Suffix    = []byte("ement")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("adjustment")
	tests[i].Suffix    = []byte("ment")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("dependent")
	tests[i].Suffix    = []byte("ent")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("adoption")
	tests[i].Suffix    = []byte("ion")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("homologou")
	tests[i].Suffix    = []byte("ou")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("communism")
	tests[i].Suffix    = []byte("ism")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("activate")
	tests[i].Suffix    = []byte("ate")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("angulariti")
	tests[i].Suffix    = []byte("iti")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("homologous")
	tests[i].Suffix    = []byte("ous")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("effective")
	tests[i].Suffix    = []byte("ive")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("bowdlerize")
	tests[i].Suffix    = []byte("ize")
	tests[i].Expected  = true
	i++



	tests[i].S         = []byte("probate")
	tests[i].Suffix    = []byte("e")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("rate")
	tests[i].Suffix    = []byte("e")
	tests[i].Expected  = true
	i++

	tests[i].S         = []byte("cease")
	tests[i].Suffix    = []byte("e")
	tests[i].Expected  = true
	i++

	for _,datum := range tests {
		if actual := hasSuffix(datum.S, datum.Suffix) ; actual != datum.Expected {
			t.Errorf("Did NOT get what was expected for calling hasSuffix() on [%s] with suffix [%s]. Expect [%d] but got [%d]", string(datum.S), string(datum.Suffix), datum.Expected, actual)
		}
	}
}

