package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Salary > 0 || r.Link != "" {
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}

type CreateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Login == "" && r.Password == "" && r.Name == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Login == "" {
		return errParamIsRequired("login", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	return nil
}

type UpdateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.Login != "" || r.Password != "" || r.Name != "" {
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	if r.Login == "" && r.Password == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Login == "" {
		return errParamIsRequired("login", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	return nil
}
