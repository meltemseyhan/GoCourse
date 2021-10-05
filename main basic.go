package main

import (
	"fmt"
)

const (
	first  = 1
	second = "second"
	third  = iota      //2
	fourth = iota      //3
	x      = iota + 6  //10
	y      = 2 << iota // 2 << 5 = 64
	z                  // 2 << 6 = 128

)

const (
	//iota is reset for each constant block
	alfa = iota
	beta
	gama
	_
	teta
)

/*
func main() {
	basic();
	arrays();
	slices();
	maps();
	structs();
}*/

func basic() {

	var i int = 42
	fmt.Println(i)

	firstname := "Meltem"
	fmt.Println(firstname)

	var f float32 = 3.14
	fmt.Println(f)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	var name *string
	fmt.Println(name)
	//Hata verir
	//fmt.Println(*name)
	name = new(string)
	fmt.Println(name)
	fmt.Println(*name)
	*name = "Meltem"
	fmt.Println(name)
	fmt.Println(*name)

	lastname := "Seyhan"
	fmt.Println(lastname)
	ptrToLastname := &lastname
	fmt.Println(ptrToLastname)
	fmt.Println(*ptrToLastname)
	lastname = "Seyhannnn"
	fmt.Println(ptrToLastname)
	fmt.Println(*ptrToLastname)

	const pi = 3.14
	fmt.Println(pi)

	//fmt.Println(first);
	//fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(alfa)
	fmt.Println(beta)
	fmt.Println(gama)
	fmt.Println(teta)
}

func arrays() {
	var myarr [3]int
	myarr[0] = 1
	myarr[1] = 2
	myarr[2] = 3
	fmt.Println(myarr)

	myarr2 := [3]int{4, 5, 6}
	fmt.Println(myarr2)
}

func slices() {
	myarr := [3]int{1, 2, 3}
	myslice := myarr[:]
	fmt.Println(myarr)
	fmt.Println(myslice)

	myarr[1] = 42
	myslice[2] = 27
	fmt.Println(myarr)
	fmt.Println(myslice)

	myslice2 := []int{1, 2, 3}
	fmt.Println(myslice2)
	myslice2 = append(myslice2, 4)
	fmt.Println(myslice2)

	s1 := myslice2[1:]
	s2 := myslice2[:3]
	s3 := myslice2[1:3]
	fmt.Println(s1, s2, s3)

}

func maps() {
	m := map[string]int{"foo": 42, "bar": 14}
	fmt.Println(m)

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	m2 := map[string]map[int]string{
		"key1": {
			0: "Meltem",
			1: "Ozan",
		},
		"key2": {
			26: "Yirmialtı",
			32: "Otuziki",
		},
	}
	fmt.Println(m2)
}

func structs() {
	type person struct {
		ID        int
		Firstname string
		Lastname  string
	}

	var meltem person
	fmt.Println(meltem)

	meltem.ID = 43
	meltem.Firstname = "Meltem"
	meltem.Lastname = "Seyhan"
	fmt.Println(meltem)

	ahmet := person{ID: 1, Firstname: "Ahmet", Lastname: "Yılmaz"}
	fmt.Println(ahmet)
}
