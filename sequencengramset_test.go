package sequencengramset

import (
	"reflect"
	"testing"
)

func Test空文字が入ると空のSetが返る(t *testing.T) {
	expectSet := make(map[string]struct{})

	if set, _ := Sequencengramset("", 2); !reflect.DeepEqual(set, expectSet) {
		t.Errorf("expected %d to eq %d", set, expectSet)
	}
}

func TestABはAB(t *testing.T) {
	expectSet := map[string]struct{}{"AB": struct{}{}}

	if set, _ := Sequencengramset("AB", 2); !reflect.DeepEqual(set, expectSet) {
		t.Errorf("expected %d to eq %d", set, expectSet)
	}
}

func Test少し長い文字列(t *testing.T) {
	expectSet := map[string]struct{}{"AB": struct{}{}, "BC": struct{}{}, "CD": struct{}{}}

	if set, _ := Sequencengramset("ABCD", 2); !reflect.DeepEqual(set, expectSet) {
		t.Errorf("expected %d to eq %d", set, expectSet)
	}
}

func Test日本語(t *testing.T) {
	expectSet := map[string]struct{}{"あい": struct{}{}, "いう": struct{}{}, "うえ": struct{}{}}

	if set, _ := Sequencengramset("あいうえ", 2); !reflect.DeepEqual(set, expectSet) {
		t.Errorf("expected %d to eq %d", set, expectSet)
	}
}

func Test0は受け付けない(t *testing.T) {
	if _, err := Sequencengramset("AB", 0); err == nil {
		t.Errorf("expect err is not nil")
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sequencengramset("ABCDEFGFALKJLAKJDLKJFLKSDJFLKSJFLSKJDFLKSJFLKKDSJFLSJDFLSKDJFKLSJDFKLSJFLKDKJFLKSDJLFKSDJFLDSKJFLDSJLKSJDFLKJDSFD", 2)
	}
}
