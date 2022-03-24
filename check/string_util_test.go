package check

import "testing"
import "fmt"

type GetHeaderKeyValueTest struct {
	arg1     string
	expected bool
	key      string
	value    string
}

var getHeaderKeyValueTests = []GetHeaderKeyValueTest{
	GetHeaderKeyValueTest{"Set-Cookie: 1P_JAR=2022-03-09-10; expires=Fri, 08-Apr-2022 10:30:18 GMT; path=/; domain=.google.com; Secure",
		true, "Set-Cookie", "1P_JAR=2022-03-09-10; expires=Fri, 08-Apr-2022 10:30:18 GMT; path=/; domain=.google.com; Secure"},
	GetHeaderKeyValueTest{"Date: Wed, 09 Mar 2022 06:08:37 GMT", true, "Date", "Wed, 09 Mar 2022 06:08:37 GMT"},
	GetHeaderKeyValueTest{"Date:Wed, 09 Mar 2022 06:08:37 GMT", false, "", ""},
	GetHeaderKeyValueTest{"Date Wed, 09 Mar 2022 06:08:37 GMT", false, "", ""},
}

func TestGetHeaderKeyValue(t *testing.T) {
	for index, test := range getHeaderKeyValueTests {
		fmt.Printf("checking %v index.\n", index)
		if status, outKey, outValue := GetHeaderKeyValue(test.arg1); status != test.expected || outKey != test.key || outValue != test.value {
			if status != test.expected {
				t.Errorf("actual_status(%t)  != expected_status(%t).", status, test.expected)
			} else if outKey != test.key {
				t.Errorf("actual_key(%q)  != expected_key(%q).", outKey, test.key)
			} else {
				t.Errorf("actual_value(%q)  != expected_value(%q).", outValue, test.value)
			}
		}
	}
}
