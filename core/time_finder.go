package core

type TimeFinder interface {
	Try(source string) []FinderResult	
	GetName() string
}
