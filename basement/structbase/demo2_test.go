package structbase

import (
	"fmt"
	"testing"
)

func TestDog_SetName(t *testing.T) {

	// 值方法的接收者是该方法所属的那个类型值的一个副本。我们在该方法内对该副本的修改一般都不会体现在原值上，
	// 除非这个类型本身是某个引用类型（比如切片或字典）的别名类型。

	// 而指针方法的接收者，是该方法所属的那个基本类型值的指针值的一个副本。
	// 我们在这样的方法内对该副本指向的值进行修改，却一定会体现在原值上

	dog := New("wang", "xxx", "dog")

	dog.SetName("wang2")
	fmt.Printf("The dog: %s\n", dog)

	// 修改不会对原值生效，因为传递的是原值的副本
	dog.SetNameOfCopy("wangCopy")
	fmt.Printf("The dog: %s\n", dog)
}

func Test_Converter(t *testing.T) {
	dog := New("wang", "xxx", "dog")

	// 值方法和指针方法的不同之处，
	// 包括这两种方法各自能做什么、不能做什么以及会影响到其所属类型的哪些方面。
	// 这涉及值的修改、方法集合和接口实现。

	// 一个指针类型实现了某某接口类型，但它的基本类型却不一定能够作为该接口的实现类型。

	//一个自定义数据类型的方法集合中仅会包含它的所有值方法，而该类型的指针类型的方法集合却囊括了前者的所有方法，包括所有值方法和所有指针方法
	// Go 语言会适时地为我们进行自动地转译，使得我们在这样的值上也能调用到它的指针方法。比如，在Cat类型的变量cat之上，之所以我们可以通过cat.SetName("monster")修改猫的名字，是因为 Go 语言把它自动转译为了(&cat).SetName("monster")，即：先取cat的指针值，然后在该指针值上调用SetName方法。

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	// 为什么值 不能被转换为具体的类型？？

	// 指针类型方法集合 包含所有的方法
	// 值类型方法集合 只包含值相关方法。
	// 之所以调用指针类型的方法，Golang 在语言层会有一次转换，先取值的指针，再调用

	_, ok := interface{}(dog).(Pet) // false
	fmt.Printf("Dog implements interface Pet: %v\n", ok)

	_, ok = interface{}(&dog).(Pet) // true
	fmt.Printf("*Dog implements interface Pet: %v\n", ok)
}
