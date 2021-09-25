package example1

import "errors"

func Deduct(total, deduction int) (int, error) {
	if deduction > total {
		return 0, errors.New("deduction must be less than or equal to total")
	}
	return total - deduction, nil
}
