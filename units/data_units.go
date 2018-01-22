package units

const (
	// byte constants
	_          = iota
	kb float64 = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	eb
	zb
	yb
)

var (
	Data = NewQuantity("bytes")

	Byte      = Data.NewUnit("byte", "B")
	KiloByte  = Data.NewUnit("kilobyte", "KB")
	MegaByte  = Data.NewUnit("megabyte", "MB")
	GigaByte  = Data.NewUnit("gigabyte", "GB")
	TeraByte  = Data.NewUnit("terabyte", "TB")
	PetaByte  = Data.NewUnit("petabyte", "PB")
	ExaByte   = Data.NewUnit("exabyte", "")
	ZettaByte = Data.NewUnit("zettabyte", "")
	YottaByte = Data.NewUnit("yottabyte", "")

	nosys   = UnitOptionSystem("") // don't assign bits to metric system
	Bit     = Data.NewUnit("bit", "b")
	ExaBit  = exa.makeUnit(Data, Bit)
	PetaBit = peta.makeUnit(Data, Bit)
	TeraBit = tera.makeUnit(Data, Bit)
	GigaBit = giga.makeUnit(Data, Bit)
	MegaBit = mega.makeUnit(Data, Bit)
	KiloBit = kilo.makeUnit(Data, Bit)

	Nibble = Data.NewUnit("nibble", "")
)

func init() {
	Data.NewRatioConv(Nibble, Bit, 4.0)
	Data.NewRatioConv(Byte, Bit, 8.0)
	Data.NewRatioConv(KiloByte, Byte, kb)
	Data.NewRatioConv(MegaByte, Byte, mb)
	Data.NewRatioConv(GigaByte, Byte, gb)
	Data.NewRatioConv(TeraByte, Byte, tb)
	Data.NewRatioConv(PetaByte, Byte, pb)
	Data.NewRatioConv(ExaByte, Byte, eb)
	Data.NewRatioConv(ZettaByte, Byte, zb)
	Data.NewRatioConv(YottaByte, Byte, yb)
}
