// shuiXianhua
package algorithm

import (
	"math"
	"strconv"
)

func GetShuiXianHua(start, end int) (daffodil []int, err error) {
	for i := start; i <= end; i++ {
		i_str := strconv.Itoa(i)
		var s float64
		for _, v := range i_str {

			v_int, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, err
			}
			s = s + math.Pow(float64(v_int), 3)
		}

		if i == int(s) {
			daffodil = append(daffodil, i)
		}
	}
	return daffodil, nil
}
