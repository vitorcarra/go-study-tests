package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name           string
		Input          interface{}
		ExcpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{
				"Chris", "London",
			},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{
				"Chris", 33,
			},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointer to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "Chris"},
				{34, "John"},
			},
			[]string{"Chris", "John"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Chris"},
				{34, "John"},
			},
			[]string{"Chris", "John"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExcpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExcpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{33, "Chris"}
			ch <- Profile{34, "John"}
			close(ch)
		}()

		var got []string
		want := []string{"Chris", "John"}

		Walk(ch, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		fn := func() (Profile, Profile) {
			return Profile{33, "A"}, Profile{34, "B"}
		}

		var got []string
		want := []string{"A", "B"}

		Walk(fn, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v but got %v", want, got)
		}
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
