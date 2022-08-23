module go.uber.org/cadence

go 1.16

require (
	github.com/apache/thrift v0.13.0 // indirect
	github.com/apache/thrift/thriftV0100 v0.0.0-20161221203622-b2a4d4ae21c7
	github.com/cristalhq/jwt/v3 v3.1.0
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.4.4
	github.com/kisielk/errcheck v1.5.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pborman/uuid v0.0.0-20160209185913-a97ce2ca70fa
	github.com/robfig/cron v1.2.0
	github.com/streadway/quantile v0.0.0-20220407130108-4246515d968d // indirect
	github.com/stretchr/testify v1.4.0
	github.com/twmb/murmur3 v1.1.6 // indirect
	github.com/uber-go/tally/v4 v4.1.1
	github.com/uber/cadence-idl v0.0.0-20220713235846-fda89e95df1e
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/tchannel-go v1.16.0
	go.uber.org/atomic v1.9.0
	go.uber.org/fx v1.13.1 // indirect
	go.uber.org/goleak v1.0.0
	go.uber.org/multierr v1.6.0
	go.uber.org/thriftrw v1.25.0
	go.uber.org/yarpc v1.55.0
	go.uber.org/zap v1.13.0
	golang.org/x/lint v0.0.0-20200130185559-910be7a94367
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
	golang.org/x/sys v0.0.0-20220403205710-6acee93ad0eb // indirect
	golang.org/x/time v0.0.0-20170927054726-6dc17368e09b
	honnef.co/go/tools v0.0.1-2019.2.3
)

replace github.com/apache/thrift/thriftV0100 => github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7
