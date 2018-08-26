package utils

import (
	"strconv"
	"math"
	"strings"
)


func Min(x, y int) int{
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func ToFloat64(v interface{}) float64 {
	if v == nil {
		return 0.0
	}

	switch v.(type) {
	case float64:
		return v.(float64)
	case string:
		vStr := v.(string)
		vF, _ := strconv.ParseFloat(vStr, 64)
		return vF
	case int:
		return float64(v.(int))
	default:
		panic("to float64 error.")
	}
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	//return fmt.Sprintf("%.f", input_num)
	return strconv.FormatFloat(input_num, 'f', -1, 64)
}


func Int64ToString(input_num int64) string {
	// to convert a float number to a string
	//return fmt.Sprintf("%.f", input_num)
	return strconv.FormatInt(input_num,10)
}

func IntToString(input int)string{
	return strconv.Itoa(input)
}

func AdjustFloat(input_num float64,precision int,floor bool) (float64,string) {
	if precision==-1{
		return input_num,strconv.FormatFloat(input_num, 'f', -1, 64)
	}
	pow10_n := math.Pow10(precision)
	var adjusted float64
	if floor{
		adjusted=math.Trunc(input_num*pow10_n) / pow10_n
	}else{
		adjusted=math.Ceil(input_num*pow10_n) / pow10_n

	}

	return adjusted,strconv.FormatFloat(adjusted, 'f', precision, 64)
}
func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}

	switch v.(type) {
	case string:
		vStr := v.(string)
		vInt, _ := strconv.Atoi(vStr)
		return vInt
	case int:
		return v.(int)
	case float64:
		vF := v.(float64)
		return int(vF)
	default:
		panic("to int error.")
	}
}


func GetPrecision(num float64) int {
	_,numStr:=AdjustFloat(num,-1,true)
	i:=strings.Index(numStr,".")
	if i==-1{
		return 0
	}else{
		return len(numStr)-i-1
	}
}