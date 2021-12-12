module proc_stream

go 1.17

replace github.com/tr2vil/its_manager/lib/logger => ../lib/logger

replace github.com/tr2vil/its_manager/lib/config => ../lib/config

replace github.com/tr2vil/its_manager/lib/message => ../lib/message

require (
	github.com/tr2vil/its_manager/lib/config v0.0.0
	github.com/tr2vil/its_manager/lib/logger v0.0.0
	github.com/tr2vil/its_manager/lib/message v0.0.0
)

require (
	github.com/confluentinc/confluent-kafka-go v1.7.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20211210111614-af8b64212486 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
