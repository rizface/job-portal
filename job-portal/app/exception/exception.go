package exception

type Exception interface {
	Code() int
	Error() string
}
