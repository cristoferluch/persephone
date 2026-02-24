package entity

type Column struct {
	Name          string
	Type          string
	Length        int
	Precision     int
	IsNullable    bool
	HasPrimaryKey bool
}
