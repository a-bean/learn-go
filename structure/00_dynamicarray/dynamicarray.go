package dynamicarray

import (
	"errors"
)

var defaultCapacity = 10

type DynamicArray struct {
	Size        int
	Capacity    int
	ElementData []any
}

func (da *DynamicArray) Put(index int, element any) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	da.ElementData[index] = element

	return nil
}

func (da *DynamicArray) Add(element any) {
	if da.Size == da.Capacity {
		da.NewCapacity()
	}

	da.ElementData[da.Size] = element
	da.Size++
}

func (da *DynamicArray) Remove(index int) error {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return err
	}

	copy(da.ElementData[index:], da.ElementData[index+1:])

	da.ElementData[da.Size-1] = nil

	da.Size--

	return nil
}

func (da *DynamicArray) Get(index int) (any, error) {
	err := da.CheckRangeFromIndex(index)

	if err != nil {
		return nil, err
	}

	return da.ElementData[index], nil
}

func (da *DynamicArray) IsEmpty() bool {
	return da.Size == 0
}

func (da *DynamicArray) GetData() []any {
	return da.ElementData[:da.Size]
}

func (da *DynamicArray) CheckRangeFromIndex(index int) error {
	if index >= da.Size || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

func (da *DynamicArray) NewCapacity() {
	if da.Capacity == 0 {
		da.Capacity = defaultCapacity
	} else {
		da.Capacity = da.Capacity << 1
	}

	newDataElement := make([]any, da.Capacity)

	copy(newDataElement, da.ElementData)

	da.ElementData = newDataElement
}
