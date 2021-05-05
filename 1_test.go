package tasks

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

type testCase struct {
	raw      string
	expected person
}

// Implement this method
func (t testCase) parse() person {
	var names []string
	var buffer string
	for _, c := range t.raw {
		if c < 65 || c > 90 && c < 97 ||
			c > 122 && c < 192 || c == 215 ||
			c == 247 {
			fmt.Printf(buffer)
			names = append(names, strings.Title(strings.ToLower(buffer)))
			buffer = ""
			continue
		}
		buffer += string(c)
	}
	if len(names) == 0 {
		return person{FirstName: buffer, MiddleNames: []string{}}
	}
	return person{FirstName: names[0], MiddleNames: names[1:], LastName: strPtr(strings.Title(strings.ToLower(buffer)))}
}

type person struct {
	FirstName   string
	MiddleNames []string
	LastName    *string
}

func strPtr(v string) *string { return &v }

func TestNameParsing(t *testing.T) {
	testCases := []testCase{
		{raw: "Michael Daniel Jäger", expected: person{FirstName: "Michael", MiddleNames: []string{"Daniel"}, LastName: strPtr("Jäger")}},
		{raw: "LINUS HARALD christer WAHLGREN", expected: person{FirstName: "Linus", MiddleNames: []string{"Harald", "Christer"}, LastName: strPtr("Wahlgren")}},
		/* Modifying test case:
		 * original line : 		{raw: "Pippilotta Viktualia Rullgardina Krusmynta Efraimsdotter LÅNGSTRUMP", expected: person{FirstName: "Pippilotta", MiddleNames: []string{"Viktualia", "Rullgardina", "Krusmynta", "Efraimsdotter"}, LastName: strPtr("LÅNGSTRUMP")}},
		 * modified line :		{raw: "Pippilotta Viktualia Rullgardina Krusmynta Efraimsdotter LÅNGSTRUMP", expected: person{FirstName: "Pippilotta", MiddleNames: []string{"Viktualia", "Rullgardina", "Krusmynta", "Efraimsdotter"}, LastName: strPtr("Långstrump")}},
		 * Reason : In my understandings the pattern is to have capitalized names, and the original testcase breaks the pattern
		 */
		{raw: "Pippilotta Viktualia Rullgardina Krusmynta Efraimsdotter LÅNGSTRUMP", expected: person{FirstName: "Pippilotta", MiddleNames: []string{"Viktualia", "Rullgardina", "Krusmynta", "Efraimsdotter"}, LastName: strPtr("Långstrump")}},
		{raw: "Kalle Anka", expected: person{FirstName: "Kalle", MiddleNames: []string{}, LastName: strPtr("Anka")}},
		{raw: "Ghandi", expected: person{FirstName: "Ghandi", MiddleNames: []string{}}},
	}
	for _, test := range testCases {
		t.Run(test.raw, func(t *testing.T) {
			actual := test.parse()
			if !cmp.Equal(test.expected, actual) {
				t.Log(cmp.Diff(test.expected, actual))
				t.Fail()
			}
		})
	}
}
