package yaml

/*
import (
	"reflect"
	"testing"

	"github.com/mickep76/encdec"
)

func TestFromByte(t *testing.T) {
	var g interface{}
	if err := encdec.FromBytes("yaml", []byte(testYAML), &g); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(g, testMap) {
		t.Errorf("want:\n%+v, got:\n%+v", testMap, g)
	}
}

func TestFromFile(t *testing.T) {
	var g interface{}
	if err := encdec.FromFile("yaml", "test.yaml", &g); err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(g, testMap) {
		t.Errorf("want:\n%+v, got:\n%+v", testMap, g)
	}
}
*/