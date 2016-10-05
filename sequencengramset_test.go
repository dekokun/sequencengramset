package sequencengramset

import (
	"reflect"
	"sort"
	"testing"
)

func Test空文字が入ると空のSetが返る(t *testing.T) {
	expectSet := []string{}

	if set, _ := Sequencengramset("", 2); !reflect.DeepEqual(set, expectSet) {
		t.Errorf("expected %d to eq %d", set, expectSet)
	}
}

func TestABはAB(t *testing.T) {
	expectSet := []string{"AB"}

	if set, _ := Sequencengramset("AB", 2); !reflect.DeepEqual(set, expectSet) {
		t.Errorf("expected %d to eq %d", set, expectSet)
	}
}

func Test少し長い文字列(t *testing.T) {
	expectSet := []string{"AB", "BC", "CD"}
	sort.Strings(expectSet)
	set, _ := Sequencengramset("ABCD", 2)
	sort.Strings(set)
	if !reflect.DeepEqual(set, expectSet) {
		t.Errorf("expected %d to eq %d", set, expectSet)
	}
}

func Test日本語(t *testing.T) {
	expectSet := []string{"あい", "いう", "うえ"}
	sort.Strings(expectSet)
	set, _ := Sequencengramset("あいうえ", 2)
	sort.Strings(set)
	if !reflect.DeepEqual(set, expectSet) {
		t.Errorf("expected %d to eq %d", set, expectSet)
	}
}

func Test0は受け付けない(t *testing.T) {
	if _, err := Sequencengramset("AB", 0); err == nil {
		t.Errorf("expect err is not nil")
	}
}

func Test重複した文字列があっても大丈夫(t *testing.T) {
	expectSet := []string{"あい", "いう", "うえ", "えあ"}
	sort.Strings(expectSet)
	set, _ := Sequencengramset("あいうえあいうえ", 2)
	sort.Strings(set)
	if !reflect.DeepEqual(set, expectSet) {
		t.Errorf("expected %d to eq %d", set, expectSet)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sequencengramset("ABCDEFGFALKJLAKJDLKJFLKSDJFLKSJFLSKJDFLKSJFLKKDSJFLSJDFLSKDJFKLSJDFKLSJFLKDKJFLKSDJLFKSDJFLDSKJFLDSJLKSJDFLKJDSFD", 2)
	}
}
