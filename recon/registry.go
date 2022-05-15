package luvrecon

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

//RegKeyData is a struct for storing the data of a registry key
type RegKeyData struct {
	Name string
	Data []string
}

//GetRegKeyData gets the data from a registry key and returns a RegKeyData struct
func GetRegKeyData(key string) RegKeyData {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer func(k registry.Key) {
		err := k.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(k)

	regStat, err := k.Stat()
	if err != nil {
		log.Fatal(err)
	}

	regValues := regStat.ValueCount

	values, err := k.ReadValueNames(int(regValues))
	if err != nil {
		log.Fatal(err)
	}
	return RegKeyData{Name: key, Data: values}
}
