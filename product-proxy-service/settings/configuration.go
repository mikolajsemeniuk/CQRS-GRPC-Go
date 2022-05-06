package settings

type Configuration interface {
	Take() configuration
}

type configuration struct {
	Port string
}

func (c configuration) Take() configuration {
	return c
}

func NewConfiguration() Configuration {
	return &configuration{
		Port: ":8080",
	}
}
