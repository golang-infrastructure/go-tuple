package tuple

//import (
//	"reflect"
//	"sort"
//)
//
//type Comparable interface {
//	CompareTo(target any) int
//}
//
//// Equals 比较值是否相同，如果实现了Comparable接口的话就使用接口的方法来比较，否则只简单的做等号比较
//func Equals(a, b any) bool {
//	ac, ok := a.(Comparable)
//	if !ok {
//		return a == b
//	}
//	bc, ok := b.(Comparable)
//	if !ok {
//		return a == b
//	}
//
//	return ac.CompareTo(bc) == 0
//}
//
//func Compare(a, b any) (int, bool) {
//	reflect.DeepEqual(a, b)
//}
//
//func SliceEquals(sliceA, sliceB []any) bool {
//	if len(sliceA) != len(sliceB) {
//		return false
//	}
//	for index := 0; index < len(sliceA); index++ {
//		if !Equals(sliceA[index], sliceB[index]) {
//			return false
//		}
//	}
//	return true
//}
//
//func SliceContains(slice []any, v any) bool {
//	for _, a := range slice {
//		if Equals(a, v) {
//			return true
//		}
//	}
//	return false
//}
//
//// SliceMax 获取数组中的最大值，数组中的值的类型必须是相同的，否则无法比较
//func SliceMax(slice []any) any {
//	isSame, reflectType := sliceElemType(slice)
//	if !isSame {
//		return nil
//	}
//	switch reflectType.Kind() {
//	case reflect.Int8:
//
//	case reflect.Int:
//		sort.SliceStable(slice, func(i, j int) bool {
//			return slice[i].(int) < slice[j].(int)
//		})
//	default:
//		for {
//
//		}
//	}
//}
//
//// SliceMin 获取数组中的最小值，数组中的值的类型必须是相同的，否则无法比较
//func SliceMin(slice []any) any {
//
//}
//
//// SliceCompare 比较两个数组的大小，两个数组对应下表位置的元素的类型应该是一致的，否则无法比较
//func SliceCompare(sliceA []any, sliceB []any) int {
//
//}
//
//// 查看数组中的元素的值的类型
//func sliceElemType(slice []any) (isSame bool, reflectType reflect.Type) {
//	if len(slice) == 0 {
//		return
//	}
//	reflectType = reflect.ValueOf(slice[0]).Type()
//	for i := 1; i < len(slice); i++ {
//		r := reflect.ValueOf(slice[0]).Type()
//		if r.Kind() != reflectType.Kind() {
//			return
//		}
//	}
//	isSame = true
//	return
//}
