module prophetstor.com/alameda

go 1.15

require (
	cloud.google.com/go v0.46.3
	github.com/Shopify/sarama v1.27.0
	github.com/golang/mock v1.4.0
	github.com/golang/protobuf v1.4.2
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.2
	github.com/influxdata/influxdb v1.7.9
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/manifoldco/promptui v0.8.0
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.1
	github.com/openshift/api v0.0.0-20200720083901-0c4b3ae5f5df
	github.com/openshift/client-go v0.0.0-20200623090625-83993cebb5ae
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/robfig/cron v1.2.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.6.1
	go.uber.org/zap v1.16.0
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/tools v0.0.0-20200616133436-c1934b75d054 // indirect
	google.golang.org/genproto v0.0.0-20200923140941-5646d36feee1
	google.golang.org/grpc v1.32.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/mail.v2 v2.3.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	k8s.io/api v0.19.0
	k8s.io/apiextensions-apiserver v0.18.9 // indirect
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/klog/v2 v2.1.0 // indirect
	prophetstor.com/api v0.0.0-00010101000000-000000000000
	sigs.k8s.io/controller-runtime v0.6.3
)

replace (
	k8s.io/api => k8s.io/api v0.18.9
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.9
	k8s.io/client-go => k8s.io/client-go v0.18.9
	prophetstor.com/api => gitlab.prophetservice.com/ProphetStor/api v0.0.0-20200925042446-63bb4a5defaf
)
