package bytes

import (
	"testing"
)

var testBytes = []byte(`[
		{"Name": "Ed", "Text": "Knock knock.", "Friends": ["[]", "\[\]", "{}", "\[]"]},
		{"Name": "Sam", "Text": "Who's there?", "Friends": "[]"}
]`)

func testIndexQuoted(b []byte) (int, int, bool) {
	return IndexQuoted(b, '\\', '"')
}

func testIndexScoped(b []byte) (int, int, bool) {
	return IndexScoped(b, '\\', '"', '{', '}')
}

func testEachPrintln(b []byte) error {
	println(string(b))
	return nil
}

func testEachNop(b []byte) error {
	return nil
}

func TestIndexQuoted(t *testing.T) {
	p, err := IndexForEach(testBytes, testIndexQuoted, testEachPrintln)
	if err != nil {
		t.Fail()
	}
	t.Log("len(testBytes)", len(testBytes), "p", p)
}

func TestIndexScoped(t *testing.T) {
	p, err := IndexForEach(testBytes, testIndexScoped, testEachPrintln)
	if err != nil {
		t.Fail()
	}
	t.Log("len(testBytes)", len(testBytes), "p", p)
}

func BenchmarkIndexQuoted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := IndexForEach(testBytes, testIndexQuoted, testEachNop)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIndexScoped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := IndexForEach(testBytes, testIndexScoped, testEachNop)
		if err != nil {
			b.Fatal(err)
		}
	}
}
