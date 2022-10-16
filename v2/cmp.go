package ternary

func Cmp[E any](comparison bool, truthy, falsy E) E {
	if comparison {
		return truthy
	}
	return falsy
}
