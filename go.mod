module asciiaas

go 1.17

require (
	github.com/stretchr/testify v1.7.0
	localmods/asciibuilder v0.0.0-00010101000000-000000000000
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
	localmods/asciigenerator v0.0.0 // indirect
	localmods/asciimodel v0.0.0 // indirect
)

replace (
	localmods/asciibuilder => ./asciibuilder
	localmods/asciigenerator => ./asciigenerator
	localmods/asciimodel => ./asciimodel
)
