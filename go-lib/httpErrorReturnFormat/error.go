package httpErrorReturnFormat

type body struct {
	Path    string
	Message string
	Error   string
	Status  int
}
