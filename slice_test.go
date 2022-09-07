package langs

import (
    "fmt"
    "testing"
)

func TestGroupSlice(t *testing.T) {
    type args struct {
        slice []int
        size  int
    }

    generateTestSlice := func() []int {
        slicelength := RandIntRange(0, 10)
        intslice := make([]int, 0, slicelength)
        for i := 0; i < slicelength; i++ {
            intslice = append(intslice, RandIntRange(0, 100))
        }
        return intslice
    }

    tests := []struct {
        name    string
        args    args
        want    [][]int
        wantErr bool
    }{
        {name: "整型slice， size=0", args: struct {
            slice []int
            size  int
        }{slice: generateTestSlice(), size: 0}, want: nil, wantErr: true},
    }

    // 随机生成100条测试数据
    for i, _ := range [100]struct{}{} {
        size := RandIntRange(1, 12)
        test := struct {
            name    string
            args    args
            want    [][]int
            wantErr bool
        }{
            name: fmt.Sprintf("测试用用例 %d", i+1),
            args: args{
                slice: generateTestSlice(),
                size:  size,
            },
            want:    nil,
            wantErr: false,
        }

        tests = append(tests, test)
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := GroupSlice(tt.args.slice, tt.args.size)
            if (err != nil) != tt.wantErr {
                t.Errorf("GroupSlice() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            t.Logf("Slice: %v, Size: %d, Got: %v\n", tt.args.slice, tt.args.size, got)

            if len(got) > 2 {
                for _, part := range got[:len(got)-1] {
                    if tt.args.size != len(part) {
                        t.Errorf("GroupSlice size error: want = %d, act = %d", tt.args.size, len(part))
                    }
                }
            }
        })
    }
}
