from golang

RUN apt-get update && apt-get install -y unzip git \
    && rm -rf /var/lib/apt/lists/*

RUN go get github.com/fjukstad/luftkvalitet
RUN go get github.com/paulmach/go.geojson

RUN mkdir -p $GOPATH/src/github.com/fjukstad/luft
ADD . $GOPATH/src/github.com/fjukstad/luft
WORKDIR  $GOPATH/src/github.com/fjukstad/luft
RUN go install 

CMD PORT=80 luft 
