[[ ! -L /go/src/news ]] && ln -s /root/workspace/GoInAction/news /go/src

mkdir -p /root/workspace/GoInAction/news/bin

for GOOS in darwin linux; do
  for GOARCH in 386 amd64; do
    export GOOS GOARCH
    go build -o bin/news-$GOOS-$GOARCH
  done
done
