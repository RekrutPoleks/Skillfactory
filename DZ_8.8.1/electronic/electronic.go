package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}

type structPhone struct {
	model string
	brand string
	_type string
}

type Smart struct {
	_os string
	structPhone
}

func (phone *structPhone) Brand() string {
	return phone.brand
}

func (phone *structPhone) Model() string {
	return phone.model
}

func (phone *structPhone) Type() string {
	return phone._type
}

func (phone *Smart) OS() string {
	return phone._os
}

type ApplePhone struct {
	Smart
}

func NewApplePhone(mdl string) ApplePhone {
	return ApplePhone{Smart: Smart{_os: "IOS", structPhone: structPhone{brand: "Apple", model: mdl, _type: "smartfone"}}}
}

type AndroidPhone struct {
	Smart
}

func NewAndroidPhone(brnd, mdl string) AndroidPhone {
	return AndroidPhone{Smart: Smart{_os: "Andorid", structPhone: structPhone{brand: brnd, model: mdl, _type: "smartfone"}}}
}

type RadioPhone struct {
	structPhone
	buttonscount int
}

func NewRadioPhone(brnd, mdl string, btsc int) RadioPhone {
	return RadioPhone{buttonscount: btsc, structPhone: structPhone{brand: brnd, model: mdl, _type: "station"}}
}

func (rph *RadioPhone) ButtonsCount() int {
	return rph.buttonscount
}
