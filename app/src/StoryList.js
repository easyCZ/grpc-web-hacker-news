"use strict";
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
Object.defineProperty(exports, "__esModule", { value: true });
var jsx_runtime_1 = require("react/jsx-runtime");
var semantic_ui_react_1 = require("semantic-ui-react");
var StoryList = function (props) {
    return ((0, jsx_runtime_1.jsx)(semantic_ui_react_1.Item.Group, __assign({ divided: true }, { children: props.stories.map(function (story, i) {
            return (0, jsx_runtime_1.jsx)(semantic_ui_react_1.Item, __assign({ style: props.selected && story.id === props.selected.id
                    ? { 'backgroundColor': 'rgba(0, 0, 0, 0.08)' }
                    : {}, onClick: function () {
                    if (story.id) {
                        props.onStorySelect(story.id);
                    }
                } }, { children: (0, jsx_runtime_1.jsxs)(semantic_ui_react_1.Item.Content, { children: [(0, jsx_runtime_1.jsx)(semantic_ui_react_1.Item.Header, __assign({ as: "a" }, { children: story.title }), void 0), (0, jsx_runtime_1.jsxs)(semantic_ui_react_1.Item.Extra, { children: [(0, jsx_runtime_1.jsx)(semantic_ui_react_1.Icon, { name: "star" }, void 0), story.score, " | ", (0, jsx_runtime_1.jsx)(semantic_ui_react_1.Icon, { name: "user" }, void 0), story.by] }, void 0)] }, void 0) }), i);
        }) }), void 0));
};
exports.default = StoryList;
