// go test -v -run="TestRemoveCIDRs"

package cidrman

import (
	"reflect"
	"testing"
)

func TestRremoveCIDRs(t *testing.T) {
	type TestCase struct {
		Input  []string
		Remove []string
		Output []string
		Error  bool
	}

	testCases := []TestCase{
		{
			Input:  nil,
			Remove: nil,
			Output: nil,
			Error:  false,
		},
		{
			Input:  []string{},
			Remove: nil,
			Output: []string{},
			Error:  false,
		},
		{
			Input:  nil,
			Remove: []string{},
			Output: nil,
			Error:  false,
		},
		{
			Input:  []string{},
			Remove: []string{},
			Output: []string{},
			Error:  false,
		},
		{
			Input:  []string{
				"10.0.0.0/8",
			},
			Remove: []string{},
			Output: []string{
				"10.0.0.0/8",
			},
			Error:  false,
		},
		{
			Input:  []string{
				"10.0.0.0/8",
			},
			Remove: nil,
			Output: []string{
				"10.0.0.0/8",
			},
			Error:  false,
		},
		{
			Input:  nil,
			Remove: []string{
				"10.0.0.0/8",
			},
			Output: nil,
			Error:  false,
		},
		{
			Input:  []string{
				"10.0.0.0/8",
			},
			Remove: []string{
				"10.0.0.0/8",
			},
			Output: []string{},
			Error:  false,
		},
		{
			Input:  []string{
				"10.0.0.0/8",
				"0.0.0.0/0",
			},
			Remove: []string{},
			// With nothing to remove, we get back what we sent
			Output: []string{
				"10.0.0.0/8",
				"0.0.0.0/0",
			},
			Error:  false,
		},
		{
			Input:  []string{
				"10.0.0.0/8",
				"10.0.0.0/8",
			},
			Remove: []string{},
			// With nothing to remove, we get back what we sent
			Output: []string{
				"10.0.0.0/8",
				"10.0.0.0/8",
			},
			Error:  false,
		},
		{
			Input:  []string{
				"192.0.128.0/24",
				"192.0.129.0/24",
			},
			Remove: []string{
				"10.0.0.0/8",
			},
			// RemoveIPNets will first do MergeIPNets() before processing the remove list
			Output: []string{
				"192.0.128.0/23",
			},
			Error:  false,
		},
		{
			Input:  []string{
				"192.0.128.0/24",
				"192.0.129.0/24",
			},
			Remove: []string{
				"192.0.128.0/23",
			},
			Output: []string{},
			Error:  false,
		},
		{
			Input:  []string{
				"192.0.128.0/24",
				"192.0.139.0/24",
			},
			Remove: []string{
				"192.0.128.0/23",
			},
			Output: []string{
				"192.0.139.0/24",
			},
			Error:  false,
		},
		{
			Input:  []string{
				"172.16.10.0/24",
				"172.16.11.0/24",
				"172.16.12.0/24",
				"172.16.13.0/24",
				"172.16.14.0/24",
			},
			Remove: []string{
				"172.16.8.0/22",
			},
			Output: []string{
				"172.16.12.0/23",
				"172.16.14.0/24",
			},
			Error:  false,
		},
		{
			Input:  []string{
				"172.16.10.0/24",
				"172.16.11.0/24",
				"172.16.12.0/24",
				"172.16.13.0/24",
				"172.16.14.0/24",
			},
			Remove: []string{
				"172.16.12.0/22",
			},
			Output: []string{
				"172.16.10.0/23",
			},
			Error:  false,
		},
		{
			Input:  []string{
				"172.16.8.0/21",
			},
			Remove: []string{
				"172.16.12.0/23",
			},
			Output: []string{
				"172.16.8.0/22",
				"172.16.14.0/23",
			},
			Error:  false,
		},
		{
			Input:  []string{
				"172.16.10.0/24",
				"172.16.11.0/24",
				"172.16.12.0/24",
				"172.16.13.0/24",
				"172.16.14.0/24",
			},
			Remove: []string{
				"172.16.8.0/21",
			},
			Output: []string{},
			Error:  false,
		},
	}

	for _, testCase := range testCases {
		output, err := RemoveCIDRs(testCase.Input, testCase.Remove)
		if err != nil {
			if !testCase.Error {
				t.Errorf("RemoveCIDRs(%#v, %#v) failed: %s", testCase.Input, testCase.Remove, err.Error())
			}
		}
		if !reflect.DeepEqual(testCase.Output, output) {
			t.Errorf("RemoveCIDRs(%#v, %#v) expected: %#v, got: %#v", testCase.Input, testCase.Remove, testCase.Output, output)
		}
	}
}
