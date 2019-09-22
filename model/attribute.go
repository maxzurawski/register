package model

type Attribute struct {
	Symbol      *string `gorm:"column:symbol;primary_key;type:varchar(50)"` // NOTE: only CAPITAL LETTERS, no whitespaces, space = _
	Name        *string `gorm:"column:name;type:varchar(100)"`
	Description *string `gorm:"column:description;type:varchar(255)"`
	Inputtype   *string `gorm:"column:inputtype;type:varchar(20)"`
}

type AttributeArray []Attribute
type Filter func(attribute *Attribute, value string) bool

func (attributes AttributeArray) FilterBy(fn Filter, value string) *Attribute {
	for _, item := range attributes {
		if fn(&item, value) {
			return &item
		}
	}
	return nil
}

func FilterBySymbol(attribute *Attribute, value string) bool {
	if *attribute.Symbol == value {
		return true
	}
	return false
}

// NOTE: Future attributes
// 	ACTIVE
//	ACCEPTABLE_MAX
//	ACCEPTABLE_MIN
