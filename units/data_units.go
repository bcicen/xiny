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

	Bit     = Data.NewUnit("bit", "b")
	ExaBit  = Exa(Bit)
	PetaBit = Peta(Bit)
	TeraBit = Tera(Bit)
	GigaBit = Giga(Bit)
	MegaBit = Mega(Bit)
	KiloBit = Kilo(Bit)

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
