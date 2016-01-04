package main

type UserSlice []User

func (slice UserSlice) Len() int {
    return len(slice)
}

func (slice UserSlice) Less(i, j int) bool {
    return slice[i].Balance < slice[j].Balance;
}

func (slice UserSlice) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

func Abs(val float32) float32 {
	if (val < 0) {
		return -val
	} else {
		return val
	}
}


