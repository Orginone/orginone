echo "移除之前的进程并重新发布"
propath=$orginone/release
if [ -d $propath ]; then
    cd $propath/production
    cd ../../
fi
rm -rf $propath/production/*

echo "编译新版本"
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-s -w --extldflags "-static -fpic"' -o $propath/production/system_api.run $orginone/app/system/cmd/api/system.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-s -w --extldflags "-static -fpic"' -o $propath/production/system_rpc.run $orginone/app/system/cmd/rpc/system.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-s -w --extldflags "-static -fpic"' -o $propath/production/user_api.run $orginone/app/user/cmd/api/user.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-s -w --extldflags "-static -fpic"' -o $propath/production/user_rpc.run $orginone/app/user/cmd/rpc/user.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-s -w --extldflags "-static -fpic"' -o $propath/production/market_api.run $orginone/app/market/cmd/api/market.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags '-s -w --extldflags "-static -fpic"' -o $propath/production/market_rpc.run $orginone/app/market/cmd/rpc/market.go