package main

import (
	"fmt"
	"testing"
)

func Test_regSentence(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1",args{"这个老板不怎么骗人"}},
		{"test2",args{"但是一次上门体验的经历还是让我非常后悔。"}},
		{"test3",args{"这块蛋糕没有异味。"}},
		{"test4",args{"这个电饭煲感觉不怎么有惊喜，买来以后感觉很后悔。"}},
		{"test5",args{"这台车的发动机有很明显的且让人厌烦的异响，我本来高高兴兴的，结果让我喜爱不起来"}},
		{"test6",args{"我非常地放心这一个没有出过问题的电吹风"}},
		{"test7",args{"我真是后悔不早点买下来，不是我不爱你，是因为一直没赚到钱"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			regSentence(tt.args.sentence)
		})
	}
}
