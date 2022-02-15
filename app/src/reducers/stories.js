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
var stories_1 = require("../actions/stories");
var initialState = {
    stories: {},
    error: null,
    loading: false,
    selected: null,
};
function default_1(state, action) {
    var _a;
    if (state === void 0) { state = initialState; }
    switch (action.type) {
        case stories_1.STORIES_INIT:
            return __assign(__assign({}, state), { loading: true });
        case stories_1.ADD_STORY:
            var story = action.payload.toObject();
            var selected = state.selected !== null ? state.selected : story;
            if (story.id) {
                return __assign(__assign({}, state), { loading: false, stories: __assign(__assign({}, state.stories), (_a = {}, _a[story.id] = story, _a)), selected: selected });
            }
            return state;
        case stories_1.SELECT_STORY:
            return __assign(__assign({}, state), { selected: state.stories[action.payload] });
        default:
            return state;
    }
}
exports.default = default_1;
