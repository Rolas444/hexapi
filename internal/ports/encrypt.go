package ports

type EncryptService interface {
	EncryptPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}
