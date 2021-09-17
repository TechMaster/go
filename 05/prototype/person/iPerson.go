package person

type iPerson interface {
	Clone() iPerson
	ChangeAddress(string)
	Print()
}