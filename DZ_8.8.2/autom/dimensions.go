package autom

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u *Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		// сконвертировать value в заданный в параметре UnitType
		switch {
		case t == Inch:
			value *= 0.39
		case t == CM:
			value *= 2.54
		}
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
	Getsize() (float64, float64, float64)
}

type size struct {
	length Unit
	width  Unit
	heigth Unit
}

func NewSizeCM(lng, wdth, hgth float64) size {
	return size{length: Unit{Value: lng, T: CM}, width: Unit{Value: wdth, T: CM}, heigth: Unit{Value: hgth, T: CM}}
}

func NewSizeInch(lng, wdth, hgth float64) size {
	return size{length: Unit{Value: lng, T: Inch}, width: Unit{Value: wdth, T: Inch}, heigth: Unit{Value: hgth, T: Inch}}
}

func (sz *size) Length() Unit {
	return sz.length
}

func (sz *size) Width() Unit {
	return sz.width
}

func (sz *size) Height() Unit {
	return sz.heigth
}

func (sz *size) Getsize() (float64, float64, float64) {
	return sz.Length().Value, sz.Width().Value, sz.Height().Value
}
