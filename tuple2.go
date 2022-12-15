package tuple

import (
	"fmt"
	reflect_utils "github.com/golang-infrastructure/go-reflect-utils"
)

type Tuple2[V1, V2 any] struct {
	V1 V1
	V2 V2
}

// ------------------------------------------------ In ---------------------------------------------------------------------

func New2[V1, V2 any](v1 V1, v2 V2) *Tuple2[V1, V2] {
	return &Tuple2[V1, V2]{
		V1: v1,
		V2: v2,
	}
}

// New2FromSlice 从Slice创建元组，slice元素会从左到右解构赋值给元组，如果长度超过了的话超过的部分会忽略，如果长度不够不够的部分会保持对应类型的零值
func New2FromSlice[T any](slice []T) *Tuple2[T, T] {
	t := &Tuple2[T, T]{}
	if len(slice) >= 2 {
		t.V1 = slice[0]
		t.V2 = slice[1]
	} else if len(slice) >= 1 {
		t.V1 = slice[0]
	}
	return t
}

// ------------------------------------------------ Use ---------------------------------------------------------------------

// Len 返回元素的长度，固定为 2
func (x *Tuple2[V1, V2]) Len() int {
	return 2
}

// Equals 判断两个元素是否相等
func (x *Tuple2[V1, V2]) Equals(target *Tuple2[V1, V2]) bool {

	// avoid panic, allow nil tuple compare
	if x == nil || target == nil {
		return x == target
	}

	// 转换为数组比较
	return SliceEquals(x.Slice(), target.Slice())
}

// Contains 元组是否包含给定的值
func (x *Tuple2[V1, V2]) Contains(v any) bool {
	// 转换为切片
	return SliceContains(x.Slice(), v)
}

// Remove 如果当前元组与目标元组的对应字段相等的话，则移除
func (x *Tuple2[V1, V2]) Remove(target Tuple2[V1, V2]) {

	if Equals(x.V1, target.V1) {
		var zero V1
		x.V1 = zero
	}

	if Equals(x.V2, target.V2) {
		var zero V2
		x.V2 = zero
	}

}

// Merge 如果当前字段是空的，则使用给定的元组的对应下标来覆盖
func (x *Tuple2[V1, V2]) Merge(target Tuple2[V1, V2]) {

	if reflect_utils.IsZero(x.V1) {
		x.V1 = target.V1
	}

	if reflect_utils.IsZero(x.V2) {
		x.V2 = target.V2
	}

}

//// Max 仅当元组中的元素类型都相同时才可以被调用
//func (x *Tuple2[V1, V2]) Max() any {
//	return SliceMax(x.Slice())
//}
//
//// Min 返回元组中最小的元素，仅当元组中的元素类型都相同时才可以被调用
//func (x *Tuple2[V1, V2]) Min() any {
//	return SliceMin(x.Slice())
//}
//
//// CompareTo 比较两个元组的大小
//func (x *Tuple2[V1, V2]) CompareTo(target Tuple2[V1, V2]) int {
//
//}

// First 返回元组的第一个元素
func (x *Tuple2[V1, V2]) First() V1 {
	return x.V1
}

// Last 返回元素中的最后一个元素
func (x *Tuple2[V1, V2]) Last() V2 {
	return x.V2
}

// ------------------------------------------------ Out ---------------------------------------------------------------------

// Unpack 把元组解构为多个单个的变量返回
func (x *Tuple2[V1, V2]) Unpack() (V1, V2) {
	return x.V1, x.V2
}

// Slice 把元组转为any切片
func (x *Tuple2[V1, V2]) Slice() []any {
	return []any{
		x.V1,
		x.V2,
	}
}

// Copy 把当前的元组拷贝一份，注意这里只是一个浅拷贝
func (x *Tuple2[V1, V2]) Copy() *Tuple2[V1, V2] {
	return &Tuple2[V1, V2]{
		x.V1,
		x.V2,
	}
}

// Concat 用给定的分隔符把元组中的所有元素ToString连接起来
func (x *Tuple2[V1, V2]) Concat(delimiter ...string) string {
	localDelimiter := ""
	if len(delimiter) > 0 {
		localDelimiter = delimiter[0]
	}
	return fmt.Sprintf("%#v%s%#v", x.V1, localDelimiter, x.V2)
}

// 把当前的元组转换为字符串，一般是用于打印之类的
func (x *Tuple2[V1, V2]) String() string {
	return fmt.Sprintf("(%#v, %#v)", x.V1, x.V2)
}

// ---------------------------------------------------------------------------------------------------------------------
