module main

go 1.17

require github.com/tr2vil/its_manager/lib v0.0.0-20211212021519-4cc9ddbadec4

require (
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/tr2vil/its_manager/lib => ../lib
