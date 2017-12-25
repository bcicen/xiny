package units

// Formatter for units
//type ValueFmt interface {
//String() string
//FString(bool, int) string
//}

//func (v *ValueFmt) String() string { return v.FString(false, 3) }

//func (v *ValueFmt) FString(long bool, precision int) string {
//label := v.unit.Symbol
//if long {
//label = v.unit.Name
//}

//if plural(v.val) {
//label = fmt.Sprintf("%ss", label)
//}

//vstr := strconv.FormatFloat(v.val, 'f', precision, 64)

//return fmt.Sprintf("%s %s", vstr, label)
//}
