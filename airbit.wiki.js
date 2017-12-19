var fs = require('fs');
var path = require('path');
var Wiki = require('wiki-serve');
var wiki = new Wiki({
    directory: __dirname,
    bootstrap: true
});
Object.keys(wiki.pages).forEach(function (name) {
    console.log(Object.keys(wiki.pages[name]));
});