module github.com/manojkva/go-redfish-plugin

go 1.13

require (
	github.com/bm-metamorph/MetaMorph v0.0.0
	github.com/go-resty/resty/v2 v2.0.0
	github.com/golang/protobuf v1.4.2
	github.com/hashicorp/go-plugin v1.3.0
	github.com/hashicorp/go-version v1.2.1
	github.com/manojkva/go-redfish-api-wrapper v1.0.6
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.4.0
	go.uber.org/zap v1.15.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	golang.org/x/tools v0.0.0-20200708183856-df98bc6d456c // indirect
	google.golang.org/grpc v1.28.0
	google.golang.org/protobuf v1.23.0
	opendev.org/airship/go-redfish/client v0.0.0 // indirect
)

replace opendev.org/airship/go-redfish/client => /root/go/src/opendev.org/airship/go-redfish/client // Use opendev/org/airship/go-redfish refs/changes/77/737177/3

replace github.com/manojkva/go-redfish-api-wrapper => /root/go/src/github.com/manojkva/go-redfish-api-wrapper //Replace the above redfish PS in the local dir of api-wrapper too

replace github.com/bm-metamorph/MetaMorph => /root/go/src/github.com/manojkva/MetaMorph
