package errors

import "testing"

func TestError_Error(t *testing.T) {
	type fields struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
	}{{
		name: "0",
		fields: fields{
			format: "unkown command",
			args:   []interface{}{},
		},
	},
		{
			name: "1",
			fields: fields{
				format: "unkown command '%s'",
				args:   []interface{}{"add"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				format: tt.fields.format,
				args:   tt.fields.args,
			}
			t.Log(e.Error())
		})
	}
}
