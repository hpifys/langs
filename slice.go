package langs

import (
    "errors"
    "fmt"
)

// GroupSlice 对切片slice按size进行分组
func GroupSlice[T any](slice []T, size int) ([][]T, error) {
    if slice == nil {
        return nil, nil
    }

    if size <= 0 {
        return nil, errors.New(fmt.Sprintf("size must great than zero, current is %d", size))
    }

    length := len(slice)

    result := make([][]T, 0, length)

    div := length / size
    trim := length % size

    for i := 0; i < div; i++ {
        sub := slice[i*size : (i+1)*size]
        result = append(result, sub)
    }

    if trim != 0 {
        sub := slice[div*size : div*size+trim]
        result = append(result, sub)
    }

    return result, nil
}

func Contains[T comparable](slice []T, ele T) bool {
    if slice == nil {
        return false
    }

    for _, t := range slice {
        if t == ele {
            return true
        }
    }

    return false
}
