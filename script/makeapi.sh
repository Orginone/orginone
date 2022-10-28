#!/bin/bash

goctl api plugin -p goctl-go-compact -api $orginone/source/api/user.api -dir $orginone/app/user/cmd/api -style "goZero"
goctl api plugin -p goctl-go-compact -api $orginone/source/api/system.api -dir $orginone/app/system/cmd/api -style "goZero"
goctl api plugin -p goctl-go-compact -api $orginone/source/api/market.api -dir $orginone/app/market/cmd/api -style "goZero"
go mod tidy

