package reflection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWal(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 20},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with nested fields",
			Input: Person{
				Name: "Chris",
				Profile: Profile{
					Age:  20,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with fields pointing to things",
			Input: &Person{
				Name: "Chris",
				Profile: Profile{
					Age:  20,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with slices",
			Input: []Profile{
				{Age: 20,
					City: "London"},
				{Age: 20,
					City: "São Paulo"},
			},
			ExpectedCalls: []string{"London", "São Paulo"},
		},
		{
			Name: "struct with arrays",
			Input: [2]Profile{
				{Age: 20,
					City: "London"},
				{Age: 20,
					City: "São Paulo"},
			},
			ExpectedCalls: []string{"London", "São Paulo"},
		},
	}

	for i, test := range cases {
		t.Run(fmt.Sprintf("[%d] %q", i, test.Name), func(t *testing.T) {

			var got []string = []string{}
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if len(got) != len(test.ExpectedCalls) {
				t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
			}

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	t.Run("struct with maps", func(t *testing.T) {
		var got []string = []string{}
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")

	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "São Paulo"}
			close(aChannel)
		}()

		var got []string = []string{}
		want := []string{"Berlin", "São Paulo"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, got []string, expected string) {
	t.Helper()

	verify := func(array []string, value string) bool {
		for _, value := range got {
			if value == expected {
				return true
			}
		}

		return false
	}

	contains := verify(got, expected)

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", got, expected)
	}
}
