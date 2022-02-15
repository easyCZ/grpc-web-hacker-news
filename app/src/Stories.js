"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
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
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var jsx_runtime_1 = require("react/jsx-runtime");
var React = __importStar(require("react"));
var react_redux_1 = require("react-redux");
var semantic_ui_react_1 = require("semantic-ui-react");
var StoryList_1 = __importDefault(require("./StoryList"));
var StoryView_1 = __importDefault(require("./StoryView"));
var stories_1 = require("./actions/stories");
var Stories = /** @class */ (function (_super) {
    __extends(Stories, _super);
    function Stories() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Stories.prototype.componentDidMount = function () {
        this.props.fetchStories();
    };
    Stories.prototype.render = function () {
        return ((0, jsx_runtime_1.jsxs)(semantic_ui_react_1.Container, __assign({ style: { padding: '1em' }, fluid: true }, { children: [(0, jsx_runtime_1.jsx)(semantic_ui_react_1.Header, __assign({ as: "h1", dividing: true }, { children: "Hacker News with gRPC-Web" }), void 0), (0, jsx_runtime_1.jsxs)(semantic_ui_react_1.Grid, __assign({ columns: 2, stackable: true, divided: 'vertically' }, { children: [(0, jsx_runtime_1.jsx)(semantic_ui_react_1.Grid.Column, __assign({ width: 4 }, { children: (0, jsx_runtime_1.jsx)(StoryList_1.default, { selected: this.props.selected, stories: this.props.stories, onStorySelect: this.props.selectStory }, void 0) }), void 0), (0, jsx_runtime_1.jsx)(semantic_ui_react_1.Grid.Column, __assign({ width: 12, stretched: true }, { children: this.props.selected
                                ? (0, jsx_runtime_1.jsx)(StoryView_1.default, { story: this.props.selected }, void 0)
                                : null }), void 0)] }), void 0)] }), void 0));
    };
    return Stories;
}(React.Component));
function mapStateToProps(state) {
    return {
        stories: Object.keys(state.stories.stories).map(function (key) { return state.stories.stories[key]; }),
        loading: state.stories.loading,
        error: state.stories.error,
        selected: state.stories.selected,
    };
}
function mapDispatchToProps(dispatch) {
    return {
        fetchStories: function () {
            dispatch((0, stories_1.listStories)());
        },
        selectStory: function (storyId) {
            dispatch((0, stories_1.selectStory)(storyId));
        },
    };
}
exports.default = (0, react_redux_1.connect)(mapStateToProps, mapDispatchToProps)(Stories);
