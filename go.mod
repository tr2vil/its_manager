module github.com/tr2vil/its_manager

go 1.17

require github.com/tr2vil/its_manager/lib v0.0.0

require (
	github.com/takama/daemon v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20200722175500-76b94024e4b6 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/tr2vil/its_manager/lib v0.0.0 => ./lib
