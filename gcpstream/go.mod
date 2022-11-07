module github.com/aurc/loggo/gcpstream

go 1.18

require (
	cloud.google.com/go/logging v1.5.0
	github.com/aurc/loggo v0.3.16
	github.com/aurc/loggo/cli v0.0.0
	github.com/rivo/tview v0.0.0-20221029100920-c4a7e501810d
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.1
	google.golang.org/api v0.84.0
	google.golang.org/genproto v0.0.0-20220719170305-83ca9fad585f
)

require (
	cloud.google.com/go v0.102.1 // indirect
	cloud.google.com/go/compute v1.7.0 // indirect
	github.com/alecthomas/participle/v2 v2.0.0-beta.5 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.5.3 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.0.0-20220520183353-fd19c99a87aa // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/net v0.0.0-20220607020251-c690dde0001d // indirect
	golang.org/x/oauth2 v0.0.0-20220608161450-d0670ef3b1eb // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/term v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/aurc/loggo/cli => ../cli
