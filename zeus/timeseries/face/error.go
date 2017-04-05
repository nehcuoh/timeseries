package face


type Error struct {
	error string
}

func (f *Error) Error() string {
	return "face Err:%s"+ f.error
}
