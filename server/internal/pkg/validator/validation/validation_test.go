package validation

import "testing"

func TestValidNameChar(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "underline",
			args: args{name: "Test_anyName"},
			want: false,
		},
		{
			name: "short-line",
			args: args{name: "any-name"},
			want: true,
		},
		{
			name: "short-line start",
			args: args{name: "-any-name"},
			want: false,
		},
		{
			name: "start number",
			args: args{name: "0any-name"},
			want: false,
		},
		{
			name: "has number",
			args: args{name: "any0name"},
			want: true,
		},
		{
			name: "Chinese",
			args: args{name: "小明"},
			want: true,
		},
		{
			name: "Chinese",
			args: args{name: "龘"},
			want: true,
		},
		{
			name: "empty",
			args: args{name: ""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidNameChar(tt.args.name); got != tt.want {
				t.Errorf("ValidNameChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
