package main

import "testing"

func Test_searchInString(t *testing.T) {
	type args struct {
		total      string
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Basic test",
			want: "UA-43128098-1",
			args: args{total: "_gaq.push(['_setAccount', 'UA-43128098-1']);", expression: `UA-\d{5,8}-\d{1,2}`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInString(tt.args.total, tt.args.expression); got != tt.want {
				t.Errorf("searchInString() = %v, want %v", got, tt.want)
			}
		})
	}
}
