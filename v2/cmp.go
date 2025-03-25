package ternary

func Cmp[E any](comparison bool, truthy, falsy E) E {
	if comparison {
		return truthy
	}
	return falsy
}

func CmpLazy[E any](comparison bool, truthy, falsy func () E) E {
	if comparison {
		return truthy()
	}
	return falsy()
}
