package application

type ProductInterface interface {
	isValid() (boo, error)
	Enabled() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLE = "disabled"
	ENABLED = "enabled"
)

type Product struct {
	id string
	name string
	price float64
	status string
}


// func (p *Product) isValid() (boolean, error) {

// }

func (p *Product) Enabled() (erro) {
	if p.price > 0 {
		p.status = ENABLED
		return nil
	}

	return errors.New("o preço deve ser maior que zero");
r

// func (p *Product) Disable() (error) {

	
// }

func (p *Product) GetId() (string) {
	return p.id
}

func (p *Product) GetName() (string) {
	return p.name
}

func (p *Product) GetStatus() (string) {
	return p.status
}

func (p *Product) GetPrice() (float64) {
	return p.price
}