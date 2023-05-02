package util

func GetStrPtr(s string) *string {
	return &s
}

func GetPtr[T any](t T) *T {
	return &t
}
