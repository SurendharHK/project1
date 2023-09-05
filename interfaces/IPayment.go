package interfaces


type Ipayment interface{
	CreatePayment( float64, int,float64)(string, error)
	
}