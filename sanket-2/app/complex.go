package app





type Complex struct{
	Real float64
	Imag float64

}
func NewCOmplex(real float64 , imag float64) *Complex{
	c := Complex{
		Real:real,
		Imag:imag,
	}
	return &c

}

