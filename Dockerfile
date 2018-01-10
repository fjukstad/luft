FROM node AS wiki-gen
ADD wiki /wiki
WORKDIR /wiki
# Install node modules for wiki-gen
RUN ["npm", "install"]
# Run wiki auto-generation
RUN ["node", "airbit.wiki.js"]

FROM golang

RUN go get github.com/fjukstad/luftkvalitet
RUN go get github.com/fjukstad/met
RUN go get github.com/paulmach/go.geojson

RUN mkdir -p $GOPATH/src/github.com/fjukstad/luft
ADD . $GOPATH/src/github.com/fjukstad/luft
WORKDIR  $GOPATH/src/github.com/fjukstad/luft
# wiki submodule directory does not need
# to be part of docker image
RUN rm -rf wiki
# Copy generated wiki HTML assets from wiki-gen stage
COPY --from=wiki-gen /wiki/_site/ $GOPATH/src/github.com/fjukstad/luft/public/wiki/
RUN go install 

CMD PORT=80 luft 
