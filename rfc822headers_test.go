package test

import (
	"bufio"
	"bytes"
	"net/textproto"
	"strings"
	"testing"
)

func TestParseHeaders(t *testing.T) {

	// Source: https://cran.r-project.org/src/contrib/PACKAGES
	in := []byte(`Package: ACNE
Version: 0.8.1
Depends: R (>= 3.0.0), aroma.affymetrix (>= 2.14.0)
Imports: MASS, R.methodsS3 (>= 1.7.0), R.oo (>= 1.19.0), R.utils (>=
        2.1.0), matrixStats (>= 0.14.2), R.filesets (>= 2.9.0),
        aroma.core (>= 2.14.0)
Suggests: DNAcopy
License: LGPL (>= 2.1)
NeedsCompilation: no
License_restricts_use: yes
`)

	rdr := bytes.NewReader(in)
	txtRdr := textproto.NewReader(bufio.NewReader(rdr))

	hdrMap := make(map[string]string, 1)
	for {
		line, err := txtRdr.ReadContinuedLine()
		// t.Logf("line: %q -- err: %q", line, err)
		if err != nil {
			t.Logf("err: %q", err)
			break
		}
		s := strings.SplitN(line, ":", 2)
		// t.Logf("s: %q", s)
		hdrMap[s[0]] = strings.TrimSpace(s[1])
	}
	t.Logf("hdrMap: %q", hdrMap)

}
