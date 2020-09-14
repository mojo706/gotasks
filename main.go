package main
import (
	"fmt"
	"sort"
)
// Person is a named struct that can be used to create other structs of the Person type
type Person struct {
	age int
}
var isExactlyTwiceAsOld = func(p []Person) bool {
	sort.Slice(p, func(i, j int) bool {
		return p[i].age < p[j].age
	})
	oldest := p[len(p)-1].age
	for i := 0; i < len(p); i++ {
		if p[i].age*2 > oldest {
			return false
		}
		for j := i+1; j < len(p); j++ {
			if p[i].age * 2 == p[j].age {
				return true
			}
		}
	}
	return false
}
var atLeastTwiceAsOld = func(p []Person) bool {
	sort.Slice(p, func(i, j int) bool {
		return p[i].age < p[j].age
	})
	youngest := p[0]
	oldest := p[len(p)-1]
	return youngest.age * 2 <= oldest.age
}
func main() {
	p := []Person{{age: 31}, {age: 67}, {age: 43}}
	fmt.Println(isExactlyTwiceAsOld(p))
	fmt.Println(atLeastTwiceAsOld(p))
}


