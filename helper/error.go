package helper

func ErrorT(err error) {
	if err != nil {
		panic(err)
	}
}
