package main

//Написать реализацию InMemory кэша
type Cache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}
