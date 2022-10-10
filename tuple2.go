package tuple

type Tuple2[V1 any, V2 any] struct {
	V1 V1
	V2 V2
}

// ------------------------------------------------ ---------------------------------------------------------------------

func New2[V1 any, V2 any](v1, v2 any) *Tuple2[V1, V2] {
	return &Tuple2[V1, V2]{
		V1: v1,
		V2: v2,
	}
}

//
//func New2FromSlice() {
//
//}
//
//func New2FromArray() {
//
//}
//
//
//
//// ------------------------------------------------ ---------------------------------------------------------------------
//
//func (x *Tuple2[V1, V2]) Len() int {
//	return 2
//}
//
//func (x *Tuple2[V1, V2]) Equals() bool {
//
//}
