"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var jsx_runtime_1 = require("react/jsx-runtime");
var StoryView = function (props) {
    var url = "http://localhost:8900/article-proxy?q=".concat(encodeURIComponent(props.story.url));
    return ((0, jsx_runtime_1.jsx)("iframe", { frameBorder: "0", style: {
            height: '100vh',
            width: '100%',
        }, src: url }, void 0));
};
exports.default = StoryView;
