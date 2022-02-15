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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var jsx_runtime_1 = require("react/jsx-runtime");
var react_dom_1 = __importDefault(require("react-dom"));
var react_redux_1 = require("react-redux");
require("./src/index.css");
var store_1 = __importDefault(require("./src/store"));
var Stories_1 = __importDefault(require("./src/Stories"));
react_dom_1.default.render((0, jsx_runtime_1.jsx)(react_redux_1.Provider, __assign({ store: store_1.default }, { children: (0, jsx_runtime_1.jsx)("div", { children: (0, jsx_runtime_1.jsx)(Stories_1.default, {}, void 0) }, void 0) }), void 0), document.getElementById('root'));
