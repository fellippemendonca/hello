package main

import (
	"fmt";
	"net/http";
	"io/ioutil";
	"github.com/fellippemendonca/hello/lib";
)

const (
	pi float64 = 3.14159265358979323846264338327950;
	testOne int = 123;
)

func main() {

	//testOne++;
	fmt.Println(pi);

	var number float64 = 10;
	//number2 := 10;
	fmt.Println(number+pi);
	

	res, _ := http.Get("http://www.fellippe.com/api/stocks");
	content, _ := ioutil.ReadAll(res.Body);
	res.Body.Close();

	fmt.Println(content);


	fmt.Println("Distance Meter:",lib.DistanceMeter(2));
	fmt.Println(lib.LibVar);

	
	b := true;
	i := 10;
	f := 10.12;
	s := "golang";
	fx := lib.DistanceMeter;
	//pkgVar := lib.LibVar;
	anonimousFx := func() string {
		return "hi there";
	};
	

	var bb bool;
	var ii int;
	var ff float64;
	var ss string;
	//var oo struct;
	

	fmt.Printf("Value: %v, Type: %T, Value: %v, Type:%T\n", b, b, bb, bb); 
	fmt.Printf("Value: %v, Type: %T, Value: %v, Type:%T\n", i, i, ii, ii);
	fmt.Printf("Value: %v, Type: %T, Value: %v, Type:%T\n", f, f, ff, ff);
	fmt.Printf("Value: %v, Type: %T, Value: %v, Type:%T\n", s, s, ss, ss);
	fmt.Printf("Value: %v, Type: %T\n", fx, fx);
	fmt.Printf("Value: %v, Type: %T\n", anonimousFx, anonimousFx);
	

	//fmt.Println(fx(1.2));
	
	fmt.Println(Test);
	fmt.Println(Test(-23,-46));
	fmt.Println(Test(-23,-46)());

	var a int;
	var pb *int;

	fmt.Println("Enter any number for a:");
	fmt.Scan(&a);

	pb = &a;
	
	fmt.Println("Memory Address of a:", &a, "and the value is", a);
	
	fmt.Println("Memory Address of b:", &pb, "and the value is", pb, "and the value stored is: ", *pb);
	
	fmt.Println("Enter any number for pointer b:");
	fmt.Scan(pb);

	
	fmt.Println("Memory Address of a:", &a, "and the value is", a);
	
	fmt.Println("Memory Address of b:", &pb, "and the value is", pb, "and the value stored is: ", *pb);
	

	pointerTest(&a);

	fmt.Println("Memory Address of a:", &a, "and the value is", a);
	
	fmt.Println("Memory Address of b:", &pb, "and the value is", pb, "and the value stored is: ", *pb);
	

	fmt.Println("\n\n");
};

func pointerTest (input *int) {
	*input = *input % 2;
}

func Test(latitude float64, longitude float64) func() float64 {
	return func () float64 {
			return latitude - longitude;
		};
}