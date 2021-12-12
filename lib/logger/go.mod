module logger

go 1.17

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/tr2vil/its_manager/lib/config v0.0.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	golang.org/x/sys v0.0.0-20211210111614-af8b64212486 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/tr2vil/its_manager/lib/config => ../config
