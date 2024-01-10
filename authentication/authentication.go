package authentication

var (
	create *CreateToken
	verify *VerifyToken
)

func GetCreateToken() *CreateToken {
	if create == nil {
		create = &CreateToken{}
	}
	return create
}

func GetVerifyToken() *VerifyToken {
	if verify == nil {
		verify = &VerifyToken{}
	}
	return verify
}
