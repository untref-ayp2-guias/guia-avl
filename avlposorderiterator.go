package guiaAvl

import (
	"github.com/untref-ayp2/data-structures/types"
)

type AVLPosOrderIterator[T types.Ordered] struct {
}

func NewAVLPosOrderIterator[T types.Ordered](root *AVLNode[T]) *AVLPosOrderIterator[T] {
	return nil
}

func (it *AVLPosOrderIterator[T]) HasNext() bool {
	return false
}

func (it *AVLPosOrderIterator[T]) Next() (T, error) {
	var data T
	return data, nil
}
