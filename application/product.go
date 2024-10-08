package application

import (
	"errors"

	uuid "github.com/satori/go.uuid"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enabled() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product Product) (ProductInterface, error)
	Disable(product Product) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}
type ProductPersistenseInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLE = "disabled"
	ENABLED = "enabled"
)

type Product struct {
	Id     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func NewProduct() *Product {
	product := Product{
		Id:     uuid.NewV4().String(),
		Status: DISABLE,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLE
	}

	if p.Status != ENABLED && p.Status != DISABLE {
		return false, errors.New("produto invalido")
	}

	if p.Price < 0 {
		return false, errors.New("produto invalido")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enabled() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("o preço deve ser maior que zero")
}

func (p *Product) Disable() error {
	if p.Price <= 0 {
		p.Status = DISABLE
		return nil
	}

	return errors.New("o produto não pode ser desabilitado")
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
