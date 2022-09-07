package langs

import "testing"

func TestStringContainsParts(t *testing.T) {
    type args struct {
        str           string
        parts         []string
        caseSensitive bool
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        // TODO: Add test cases.
        {name: "用例1", args: struct {
            str           string
            parts         []string
            caseSensitive bool
        }{str: "abc", parts: nil, caseSensitive: false}, want: false},
        {name: "用例2", args: struct {
            str           string
            parts         []string
            caseSensitive bool
        }{str: "abc", parts: []string{}, caseSensitive: false}, want: false},
        {name: "用例3", args: struct {
            str           string
            parts         []string
            caseSensitive bool
        }{str: "abc", parts: []string{"A"}, caseSensitive: false}, want: true},
        {name: "用例4", args: struct {
            str           string
            parts         []string
            caseSensitive bool
        }{str: "abc", parts: []string{"A"}, caseSensitive: true}, want: false},
        {name: "用例5", args: struct {
            str           string
            parts         []string
            caseSensitive bool
        }{str: "abc", parts: []string{"d"}, caseSensitive: false}, want: false},
        {name: "用例6", args: struct {
            str           string
            parts         []string
            caseSensitive bool
        }{str: "中文adgc啦三的冷", parts: []string{"A"}, caseSensitive: false}, want: true},
        {name: "用例7", args: struct {
            str           string
            parts         []string
            caseSensitive bool
        }{str: "中文adgc啦三的冷", parts: []string{"A"}, caseSensitive: true}, want: false},
        {name: "用例8", args: struct {
            str           string
            parts         []string
            caseSensitive bool
        }{str: "中文adgc啦三的冷", parts: []string{"A", "三的"}, caseSensitive: true}, want: true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := StringContainsParts(tt.args.str, tt.args.parts, tt.args.caseSensitive); got != tt.want {
                t.Errorf("StringContainsParts() = %v, want %v", got, tt.want)
            }
        })
    }
}
