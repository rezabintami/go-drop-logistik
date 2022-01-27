package convertpointer

func ConvertPointerString(x *string) *string {
	if *x == "" {
		x = nil
	}
	return x
}

func ConvertNilPointerString(x *string) string {
	if x == nil {
		return ""
	}
	return *x
}
