package main

import (
	"testing"
)

func Test_processFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "示例-1",
			args: args{
				filePath: "../../test/demo1.xlsx",
			},
		},
		{
			name: "示例-2",
			args: args{
				filePath: "../../test/demo2.xlsx",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processFile(tt.args.filePath)
		})
	}
}
