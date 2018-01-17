'use strict';

var fs = require('fs');
var path = require('path');
var marked = require('marked');
var escapeHtml = require('escape-html');
var validurl = require('valid-url');
let markdownExtensions = [".md", ".markdown"];

const NAME_TOKEN = "d792aafa37404406b4ea8f526b479e01";
const LINK_TOKEN = "d5210c06f33b4edeb6f777113049127a";
const SIDEBAR_TOKEN = "ca68a7c31184428f8a99ec8b9bc15dba";
const CONTENT_TOKEN = "c81addc8f61b4d8da7418500fb88b242";
const FOOTER_TOKEN = "ed3c8e5ccb9c488e90fb78b1cbd3c99d";
const TEMPLATE_FILE = path.join(__dirname, "airbit.uit.no.wiki.template.html");

fs.readFile(TEMPLATE_FILE, "utf8", function (templateError, HTML_TEMPLATE) {
    if (typeof templateError !== "undefined" && templateError) {
        console.error("Unabled to read HTML template file: " + TEMPLATE_FILE,
            templateError);
        return;
    }

    const targetdir = path.join(__dirname, "_site");
    if (!fs.existsSync(targetdir)) {
        fs.mkdir(targetdir, function (targetDirError) {
            if (typeof targetDirError !== "undefined" && targetDirError) {
                console.error("Unable to create directory: " + targetdir,
                    targetDirError);
            }
        });
    }

    fs.readdir(__dirname, function (readDirError, files) {
        if (typeof readDirError !== "undefined" && readDirError) {
            console.error(readDirError);
            return;
        }

        var wikiPages = {};
        var wikiLinks = {};
        var assetFiles = {};

        var pageList = files.filter(function (filepath) {
            return markdownExtensions.some(function (ext) {
                return path.extname(filepath).toUpperCase() == ext.toUpperCase();
            })
        }).map(function (filepath) {
            var filename = path.basename(filepath, path.extname(filepath));
            var wikiname = filename.replace(/\-/g, " ");
            var page = {
                filepath: filepath,
                wikilink: filename,
                name: wikiname,
                markdown: fs.readFileSync(filepath, "utf8"),
                targetfile: path.join(targetdir, filename + ".html")
            };
            wikiPages[page.name] = page;
            wikiLinks[page.wikilink] = page;
            return page;
        });
        pageList.forEach(function (page) {
            var markedOptions = {
                breaks: false,
                gfm: true,
                sanitize: false,
                tables: true,
                xhtml: true
            };

            var renderer = new marked.Renderer(markedOptions);
            var markedLinkRenderer = renderer.link;
            renderer.link = function (href, title, text) {
                if (!validurl.isUri(href)) {
                    // HREF is not a URL, asume path
                    var filepath = href;
                    var poundIdx = filepath.indexOf("#");
                    if (poundIdx >= 0) {
                        filepath = filepath.substr(0, poundIdx); // Strip URL-fragment from filepath
                        if (filepath.trim().length < 1) {
                            return markedLinkRenderer.call(renderer, href, title, text);
                        }
                    }
                    if (!fs.existsSync(filepath)) {
                        var isMarkdownLink = markdownExtensions.some(function (ext) {
                            var markdownPath = filepath + ext;
                            return fs.existsSync(markdownPath);
                        });
                        if (isMarkdownLink) {
                            var htmlPath = filepath + ".html";
                            return markedLinkRenderer.call(renderer, htmlPath, title, text);
                        } else if (!filepath.startsWith("#")) {
                            console.warn("[" + page.filepath + "]",
                                "Link points to unresolved path:",
                                filepath);
                        }
                    } else if (!path.isAbsolute(filepath)) {
                        assetFiles[filepath] = filepath;
                    }
                }
                return markedLinkRenderer.call(renderer, href, title, text);
            };
            var markedImgRenderer = renderer.image;
            renderer.image = function (href, title, text) {
                if (!validurl.isUri(href)) {
                    // HREF is not a URL, asume path
                    var filepath = href;
                    if (!fs.existsSync(filepath)) {
                        var isMarkdownLink = markdownExtensions.some(function (ext) {
                            var markdownPath = filepath + ext;
                            return fs.existsSync(markdownPath);
                        });
                        if (isMarkdownLink) {
                            var htmlPath = filepath + ".html";
                            return markedImgRenderer.call(renderer, htmlPath, title, text);
                        } else {
                            console.warn("Image sources to unresolved path: ",
                                filepath);
                        }
                    } else if (!path.isAbsolute(filepath)) {
                        assetFiles[filepath] = filepath;
                    }
                }
                return markedImgRenderer.call(renderer, href, title, text);
            };
            var markedHeadingRenderer = renderer.heading;
            renderer.heading = function (text, level, raw) {
                const HELPLINK_TOKEN = "HELPLINK";
                if (raw.toUpperCase().endsWith(HELPLINK_TOKEN)) {
                    var removeHelplinkToken = function (text) {
                        return text
                            .substr(0, text.length - HELPLINK_TOKEN.length)
                            .trimRight()
                            ;
                    }
                    raw = removeHelplinkToken(raw);
                    text = removeHelplinkToken(text);                    
                    text +=
                        `<a class="wiki-helplink" href="mailto:skolelaboratoriet@nt.uit.no?subject=%5Bairbit%5D%20-%20Hjelp%21%20%7C%20${encodeURIComponent(page.name)}%20-%20${encodeURIComponent(raw)}&body=Hjelp%20jeg%20sitter%20fast%21%0D%0AJeg%20leser%20f%C3%B8lgende%20side%3A%20http%3A%2F%2Fairbit.uit.no%2Fpublic%2Fwiki%2F${encodeURIComponent(page.wikilink)}.html${encodeURIComponent('#')}${encodeURIComponent(raw.toLowerCase().replace(/[^\w]+/g, '-'))}">
    Hjelp meg! skolelaboratoriet@nt.uit.no
</a>`;
                }
                return markedHeadingRenderer.call(renderer, text, level, raw);
            };
            markedOptions.highlight = function (code, lang) {
                var highlightjs = require('highlight.js');
                var highlightlang = highlightjs.getLanguage(lang);
                if (typeof highlightlang === "undefined" || !highlightlang) {
                    return code;
                }
                var highlighted = highlightjs.highlight(lang, code);
                return highlighted.value;
            };
            markedOptions.renderer = renderer;

            marked(page.markdown, markedOptions, function (markdownError, htmlPartial) {
                if (typeof markdownError !== "undefined" && markdownError) {
                    console.error(markdownError);
                }
                page.htmlPartial = htmlPartial;
            });
        });
        var sidebarPage = wikiPages._Sidebar;
        var footerPage = wikiPages._Footer;
        var sidebarContent = "";
        if (typeof sidebarPage !== "undefined" && sidebarPage) {
            sidebarContent = sidebarPage.htmlPartial;
        } else {
            sidebarPage = null;
        }
        var footerContent = "";
        if (typeof footerPage !== "undefined" && footerPage) {
            sidebarContent = footerPage.htmlPartial;
        } else {
            footerPage = null;
        }
        pageList.forEach(function (page) {
            if (page === sidebarPage || page === footerPage) {
                return;
            }
            var pagenamehtml = escapeHtml(page.name);
            var fullHtml = HTML_TEMPLATE
                .replace(new RegExp(NAME_TOKEN, "g"), pagenamehtml)
                .replace(new RegExp(SIDEBAR_TOKEN, "g"), sidebarContent)
                .replace(new RegExp(CONTENT_TOKEN, "g"), page.htmlPartial)
                .replace(new RegExp(FOOTER_TOKEN, "g"), footerContent)
                ;
            page.full = fullHtml;

            fs.writeFile(page.targetfile, page.full, "utf8", function (targetfileError) {
                if (typeof targetfileError !== "undefined" && targetfileError) {
                    console.error("Unable to write Wiki page '" + page.name + "' to target file: " + page.targetfile,
                        targetfileError);
                    return;
                }
            })
        });
        Object.keys(assetFiles).forEach(function (assetFile) {
            fs.readFile(assetFile, { encoding: null }, function (assetReadError, assetData) {
                if (typeof assetReadError !== "undefined" && assetReadError) {
                    console.error(assetReadError);
                    return;
                }
                var targetPath = path.join(targetdir, assetFile);
                fs.writeFile(targetPath, assetData, { encoding: null }, function (assetWriteError) {
                    if (typeof assetWriteError !== "undefined" && assetWriteError) {
                        console.error(assetWriteError);
                        return;
                    }
                });
            });
        });
    });
});