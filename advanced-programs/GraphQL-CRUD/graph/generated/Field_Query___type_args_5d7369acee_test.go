package main
import (
"context"
"reflect"
"testing"
"github.com/99designs/gqlgen/graphql"
)

type executionContext struct {
}

func (ec *executionContext) unmarshalNString2string(ctx context.Context, v interface{}) (string, error) {
return v.(string), nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
var err error
args := map[string]interface{}{}
var arg0 string
if tmp, ok := rawArgs["name"]; ok {
arg0, err = ec.unmarshalNString2string(ctx, tmp)
if err != nil {
return nil, err
}
}
args["name"] = arg0
return args, nil
}

func TestField_Query___type_args(t *testing.T) {
ctx := context.Background()
ec := &executionContext{}
tests := []struct {
name     string
rawArgs  map[string]interface{}
wantArgs map[string]interface{}
wantErr  bool
}{
{"No args", nil, map[string]interface{}{"name": ""}, false},
{"Args with non-string name", map[string]interface{}{"name": 123}, nil, true},
{"Name arg present", map[string]interface{}{"name": "test"}, map[string]interface{}{"name": "test"}, false},
}
for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
gotArgs, err := ec.field_Query___type_args(ctx, tt.rawArgs)
if (err != nil) != tt.wantErr {
t.Errorf("executionContext.field_Query___type_args() error = %v, wantErr %v", err, tt.wantErr)
return
}
if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
t.Errorf("executionContext.field_Query___type_args() = %v, want %v", gotArgs, tt.wantArgs)
}
})
}
}
