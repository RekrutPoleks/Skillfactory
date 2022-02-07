package autom

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type Automobil struct {
	brand       string
	model       string
	maxSpeed    int
	enginePower int
	size
}

func NewAutoEuro(brnd, mdl string, maxS, pwr int, lng, wdth, hgth float64) Automobil {
	return Automobil{brand: brnd, model: mdl, maxSpeed: maxS, enginePower: pwr, size: NewSizeCM(lng, wdth, hgth)}
}

func NewAutoAmerica(brnd, mdl string, maxS, pwr int, lng, wdth, hgth float64) Automobil {
	return Automobil{brand: brnd, model: mdl, maxSpeed: maxS, enginePower: pwr, size: NewSizeInch(lng, wdth, hgth)}
}

func (auto *Automobil) Brand() string {
	return auto.brand
}

func (auto *Automobil) Model() string {
	return auto.model
}

func (auto *Automobil) Dimensions() Dimensions {
	return &auto.size
}

func (auto *Automobil) MaxSpeed() int {
	return auto.maxSpeed
}

func (auto *Automobil) EnginePower() int {
	return auto.enginePower
}
