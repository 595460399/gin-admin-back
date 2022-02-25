package modelInterface

type CURD interface {
	Create() (error, CURD)
	Update() (error, CURD)
	Read() (error, CURD)
	Delete() (error, CURD)
}
