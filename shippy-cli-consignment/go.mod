module github.com/xoronedge/shippy/shippy-cli-consignment

go 1.15

//replace github.com/xoronedge/shippy/shippy-service-consignment => ../shippy-service-consignment

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/micro/micro/v3 v3.0.4
	github.com/xoronedge/shippy/shippy-service-consignment v0.0.0-20210110172803-e09e968b98dd
)
