package entity

type Category struct {
	ID       string `gorm:"primary_key" json:"category_id"`
	Category string `json:"category"`
}

type Categories []Category

func (c Categories) Len() int {
	return c.Len()
}

func (c Categories) Swap(i, j int) {
	c[i], c[j] = c[j],c[i]
}

func (c Categories) Less(i, j int) bool {
	return c[i].Category < c[j].Category
}