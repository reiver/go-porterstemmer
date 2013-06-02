package porterstemmer



import (
	"bufio"
	"os"
	"strings"
    "testing"
)



func TestStemString(t *testing.T) {

	vocFileName := "testdata/voc.txt"
	vocFd, err := os.Open(vocFileName)
	if nil != err {
		t.Errorf("Could NOT open testdata file: [%s]. Received error: [%v]", vocFileName, err)
/////// RETURN
		return
	}
	defer vocFd.Close()

	voc := bufio.NewReaderSize(vocFd, 1024)



	outFileName := "testdata/output.txt"
	outFd, err := os.Open(outFileName)
	if nil != err {
		t.Errorf("Could NOT open testdata file: [%s]. Received error: [%v]", outFileName, err)
/////// RETURN
		return
	}
	defer outFd.Close()

	out := bufio.NewReaderSize(outFd, 1024)



	for {

		vocS, err := voc.ReadString('\n')
		if nil != err {
	/////// BREAK
			break
		}

		vocS = strings.Trim(vocS, "\n\r\t ")



		expected, err := out.ReadString('\n')
		if nil != err {
			t.Errorf("Received unexpected error when trying to read a line from [%s]. Received error: [%v]", outFileName, err)
	/////// BREAK
			break

		}

		expected = strings.Trim(expected, "\n\r\t ")



		actual := StemString(vocS)
		if expected != actual {
			t.Errorf("Input: [%s] -> Actual: [%s]. Expected: [%s]", vocS, actual, expected)
		}
	}
}
