package iteration

import "testing"

func TestRepeat(t *testing.T) {

	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("  expected '%q' got '%q'   ", expected, repeated)

	}

	repeated = Repeat("b", 3)
	expected = "bbb"
	if repeated != expected {
		t.Errorf(" expected '%q' got '%q'", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {

	repeated := Repeat("abc", 2)
	println(repeated)
	//ouput "abcabc"
}
