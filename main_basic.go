package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
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

func main() {
	//basic();
	//arrays();
	//slices();
	//maps();
	//structs();
	//fmtUsage()
	//logUsage()
	//logUsage2()
	//traceUsage()
	timeUsage()
}

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

func fmtUsage() {
	file, err := os.Open("family.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var age int
		var name string
		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)
		if err != nil {
			panic(err)
		}
		if n == 2 {
			fmt.Printf("%s,%d\n", name, age)
		}
	}
}

func logUsage() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("This is a log message !")

}

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func logUsage2() {
	writeLog(INFO, "This is an info message !")
	writeLog(WARNING, "This is a warning message !")
	writeLog(ERROR, "This is an error message !")
	writeLog(FATAL, "This is a fatal message !")
}

func writeLog(messagetype messageType, message string) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	switch messagetype {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message)
	}
}

// Oluşan dosyayı okumak için go tool trace trace.out
func traceUsage() {
	file, err := os.OpenFile("trace.out", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open trace file %v\n", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Failed to close trace file %v\n", err)
		}
	}()

	if err := trace.Start(file); err != nil {
		log.Fatalf("Failed to start trace %d\n", err)
	}
	defer trace.Stop()

	n1 := rand.Intn(100)
	n2 := rand.Intn(100)
	time.Sleep(2 * time.Second)
	n3 := n1 * n2
	fmt.Printf("Result is %d\n", n3)

}
func timeUsage() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.RFC1123))
	//Mon, 02 Jan 2006 15:04:05 MST
	fmt.Println(t.Format("02.01.2006 15:04 Mon"))
	startDate := time.Date(2021, 10, 7, 11, 20, 00, 00, time.UTC)
	fmt.Println(startDate.Format("02.01.2006 15:04 Mon"))
}
