package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age      int
	Location string
}

// golang challenge: write a function walk(x interface{}, fn func(string)) which takes a struct x
// and calls fn for all strings fields found inside. difficulty level: recursively.
func TestWalk(t *testing.T) {
	expectedCases := []struct {
		Name string
		// any is an alias for interface{}
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Wondini"},
			[]string{"Wondini"},
		},
		{
			"struct with two string fields",
			struct {
				Name     string
				Location string
			}{"Wondini", "Oxford"},
			[]string{"Wondini", "Oxford"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Wondini", 42},
			[]string{"Wondini"},
		},
		// witness the mess of nested anonymous structs
		// {
		// 	"struct with nested fields",
		// 	struct {
		// 		Name    string
		// 		Profile struct {
		// 			Age  int
		// 			City string
		// 		}
		// 	}{"Wondini", struct {
		// 		Age  int
		// 		City string
		// 	}{42, "Oxford"}},
		// 	[]string{"Wondini", "Oxford"},
		// },
		{
			"struct with nested fields",
			Person{"Wondini", Profile{42, "Oxford"}},
			[]string{"Wondini", "Oxford"},
		},
		{
			"we got a pointer",
			&Person{
				"Wondini",
				Profile{42, "Oxford"},
			},
			[]string{"Wondini", "Oxford"},
		},
		{
			"thing is a slice of things",
			[]Profile{
				{27, "Reading"},
				{42, "Oxford"},
			},
			[]string{"Reading", "Oxford"},
		},
		{
			"thing is an array of things",
			[2]Profile{
				{27, "Reading"},
				{42, "Oxford"},
			},
			[]string{"Reading", "Oxford"},
		},
	}

	for _, test := range expectedCases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	t.Run("test with maps separately as the ignore field order", func(t *testing.T) {
		mapped := map[string]string{
			"Horse":  "Neigh",
			"Donkey": "Eeyore",
		}

		var got []string

		walk(mapped, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Neigh")
		assertContains(t, got, "Eeyore")
	})

	t.Run("test with channel", func(t *testing.T) {
		channel := make(chan Profile)

		go func() {
			channel <- Profile{27, "Reading"}
			channel <- Profile{42, "Oxford"}
			close(channel)
		}()

		var got []string
		want := []string{"Reading", "Oxford"}

		walk(channel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("test function", func(t *testing.T) {
		function := func() (Profile, Profile) {
			return Profile{27, "Reading"}, Profile{42, "Oxford"}
		}

		var got []string
		want := []string{"Reading", "Oxford"}

		walk(function, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
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
		t.Errorf("expected %v to contain %q", haystack, needle)
	}
}
