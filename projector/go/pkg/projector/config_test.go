package projector_test

import (
	"github.com/mubashirjamali101/aoc/pkg/projector"
	"reflect"
	"testing"
)

func getOpts(args []string) *projector.Opts {
	opts := &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}

	return opts
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation projector.Operation) {
	opts := getOpts(args)

	config, err := projector.NewConfig(opts)

	if err != nil {
		t.Errorf("Expected to get no error %v", err)
	}

	if !reflect.DeepEqual(expectedArgs, config.Args) {
		t.Errorf("Expected Args to be %+v but got %+v",args, config.Args)
	}
	
	if config.Operation != operation {
		t.Errorf("operation expected was %v but got %v", operation, config.Operation)
	}
}

func TestConfigPrintALL(t *testing.T) {
	testConfig(t, []string{}, []string{}, projector.Print)
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
	testConfig(t, []string{"add","foo","bar"}, []string{"foo","bar"}, projector.Add)
}

func TestConfigRemoveKey(t *testing.T) {
	testConfig(t, []string{"rm","foo"}, []string{"foo"}, projector.Remove)
}
