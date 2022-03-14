package git

type Client interface {
	Clone() error
	Pull() error
	Add(filePath string) error
	AddAll() error
	IsLocalUpToDate() bool
	Commit(message string) error
	Reset(previousHead int) error
	Push() error
	GetRepositoryName() string
}

type GitAuth interface {
	SetGlobalCredentials() error
}
