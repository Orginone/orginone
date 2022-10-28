if [ ! -d $orginone/doc ]; then
    mkdir $orginone/doc
fi

if [ ! -d $orginone/doc/swagger ]; then
    mkdir $orginone/doc/swagger
fi

echo "generate swagger file"
goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api $orginone/source/api/user.api -dir $orginone/doc/swagger/
goctl api plugin -plugin goctl-swagger="swagger -filename system.json" -api $orginone/source/api/system.api -dir $orginone/doc/swagger/
goctl api plugin -plugin goctl-swagger="swagger -filename market.json" -api $orginone/source/api/market.api -dir $orginone/doc/swagger/
echo "generate schema file"
go run entgo.io/ent/cmd/ent describe $orginone/source/schema > $orginone/doc/schema.txt