package main

import (
	"fmt"
)

func GetPrefix(name string) (prefix string) {
	var prefixMap map[string]string
	prefixMap = make(map[string]string)

	prefixMap["Bob"] = "Mr "
	prefixMap["Joe"] = "Dr "
	prefixMap["Mary"] = "Dr "
	prefixMap["Amy"] = "Mrs "

	return prefixMap[name]
}

func GetPrefixShorthand(name string) (prefix string) {
	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Mary": "Dr ",
		"Amy":  "Mrs ",
	}
	prefixMap["Amy"] = "Jr "
	delete(prefixMap, "Amy")
	delete(prefixMap, "Bob")
	delete(prefixMap, "Bob") //no error
	if value, isExists := prefixMap[name]; isExists {
		return value
	}
	return "NONONO"
}

func main() {
	fmt.Println(GetPrefix("Amy"), "Amy")
	fmt.Println(GetPrefixShorthand("Amy"), "Amy")
}
