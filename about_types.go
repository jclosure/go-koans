package go_koans

type coolNumber int

func (cn coolNumber) multiplyByTwo() int {
	return int(cn) * 2
}
func (cn coolNumber) multiplyByThree() int {
	return int(cn) * 3
}

func aboutTypes() {
	i := coolNumber(4)
	assert(i == coolNumber(4))     // values can be converted between compatible types
	assert(i.multiplyByTwo() == 8) // you can add methods on any type you define
	assert(i.multiplyByThree() == 12) // you can add methods on any type you define
}
