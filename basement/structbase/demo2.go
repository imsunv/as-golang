package structbase

import "fmt"

type Dog struct {
	name           string // 名字。
	scientificName string // 学名。
	category       string // 动物学基本分类。
}

func New(name, scientificName, category string) Dog {
	return Dog{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) SetNameOfCopy(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) ScientificName() string {
	return dog.scientificName
}

func (dog Dog) Category() string {
	return dog.category
}

func (dog Dog) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		dog.scientificName, dog.category, dog.name)
}
