package atomicx

import "sync/atomic"

type Value[T any] struct {
	val atomic.Value
}

// NewValue 创建Value对象，默认零值
func NewValue[T any]() *Value[T] {
	var t T
	return NewValueOf(t)
}

// NewValueOf  根据传的值创建Value对象
func NewValueOf[T any](t T) *Value[T] {
	val := atomic.Value{}
	val.Store(t)
	return &Value[T]{
		val: val,
	}
}

// Load 获取值
func (v *Value[T]) Load() (val T) {
	val = v.val.Load().(T)
	return
}

// Store 设置值
func (v *Value[T]) Store(val T) {
	v.val.Store(val)
}

// Swap 设置值并返回旧值
func (v *Value[T]) Swap(new T) (old T) {
	old = v.val.Swap(new).(T)
	return
}

// CompareAndSwap 比较并设置值
func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.val.CompareAndSwap(old, new)
}
