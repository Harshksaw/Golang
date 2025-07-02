package app






type Complex struct{
	real float64
	imag float64

}
func NewCOmplex(real float64 , imag float64) *Complex{
	c := Complex{
		real:real,
		imag:imag,
	}
	return &c

}

