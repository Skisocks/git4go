package git

type Auth struct {
	Username string
	Email    string
	Token    string
}

func NewAuth(userName string, email string, token string) *Auth {
	return &Auth{
		Username: userName,
		Email:    email,
		Token:    token,
	}
}

func (a *Auth) SetGlobalCredentials() error {
	_, err := RunCommand("", "config", "--global", "--add", "user.name", a.Username)
	_, err = RunCommand("", "config", "--global", "--add", "user.email", a.Email)
	_, err = RunCommand("", "config", "--global", "credential.helper", "store")
	return err
}
