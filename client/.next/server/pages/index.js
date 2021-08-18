"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
(() => {
var exports = {};
exports.id = "pages/index";
exports.ids = ["pages/index"];
exports.modules = {

/***/ "./src/app/hooks.ts":
/*!**************************!*\
  !*** ./src/app/hooks.ts ***!
  \**************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"useAppDispatch\": () => (/* binding */ useAppDispatch),\n/* harmony export */   \"useAppSelector\": () => (/* binding */ useAppSelector)\n/* harmony export */ });\n/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react-redux */ \"react-redux\");\n/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_redux__WEBPACK_IMPORTED_MODULE_0__);\n\nconst useAppDispatch = () => (0,react_redux__WEBPACK_IMPORTED_MODULE_0__.useDispatch)();\nconst useAppSelector = react_redux__WEBPACK_IMPORTED_MODULE_0__.useSelector;//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvYXBwL2hvb2tzLnRzLmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7QUFBQTtBQUdPLE1BQU1FLGNBQWMsR0FBRyxNQUFNRix3REFBVyxFQUF4QztBQUNBLE1BQU1HLGNBQStDLEdBQUdGLG9EQUF4RCIsInNvdXJjZXMiOlsid2VicGFjazovLy8uL3NyYy9hcHAvaG9va3MudHM/NTU1ZiJdLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBUeXBlZFVzZVNlbGVjdG9ySG9vaywgdXNlRGlzcGF0Y2gsIHVzZVNlbGVjdG9yIH0gZnJvbSBcInJlYWN0LXJlZHV4XCI7XG5pbXBvcnQgdHlwZSB7IEFwcERpc3BhdGNoLCBSb290U3RhdGUgfSBmcm9tIFwiLi9zdG9yZVwiO1xuXG5leHBvcnQgY29uc3QgdXNlQXBwRGlzcGF0Y2ggPSAoKSA9PiB1c2VEaXNwYXRjaDxBcHBEaXNwYXRjaD4oKTtcbmV4cG9ydCBjb25zdCB1c2VBcHBTZWxlY3RvcjogVHlwZWRVc2VTZWxlY3Rvckhvb2s8Um9vdFN0YXRlPiA9IHVzZVNlbGVjdG9yO1xuIl0sIm5hbWVzIjpbInVzZURpc3BhdGNoIiwidXNlU2VsZWN0b3IiLCJ1c2VBcHBEaXNwYXRjaCIsInVzZUFwcFNlbGVjdG9yIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./src/app/hooks.ts\n");

/***/ }),

/***/ "./src/features/counter/actions.ts":
/*!*****************************************!*\
  !*** ./src/features/counter/actions.ts ***!
  \*****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"increment\": () => (/* binding */ increment),\n/* harmony export */   \"decrement\": () => (/* binding */ decrement),\n/* harmony export */   \"incrementByAmount\": () => (/* binding */ incrementByAmount)\n/* harmony export */ });\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @reduxjs/toolkit */ \"@reduxjs/toolkit\");\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__);\n\nconst increment = (0,_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__.createAction)(\"counter/increment\");\nconst decrement = (0,_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__.createAction)(\"counter/decrement\");\nconst incrementByAmount = (0,_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__.createAction)(\"counter/incrementByAmount\");//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvZmVhdHVyZXMvY291bnRlci9hY3Rpb25zLnRzLmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7O0FBQUE7QUFFTyxNQUFNQyxTQUFTLEdBQUdELDhEQUFZLENBQUMsbUJBQUQsQ0FBOUI7QUFDQSxNQUFNRSxTQUFTLEdBQUdGLDhEQUFZLENBQUMsbUJBQUQsQ0FBOUI7QUFDQSxNQUFNRyxpQkFBaUIsR0FBR0gsOERBQVksQ0FDM0MsMkJBRDJDLENBQXRDIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vLy4vc3JjL2ZlYXR1cmVzL2NvdW50ZXIvYWN0aW9ucy50cz9mNzU5Il0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IGNyZWF0ZUFjdGlvbiB9IGZyb20gXCJAcmVkdXhqcy90b29sa2l0XCI7XG5cbmV4cG9ydCBjb25zdCBpbmNyZW1lbnQgPSBjcmVhdGVBY3Rpb24oXCJjb3VudGVyL2luY3JlbWVudFwiKTtcbmV4cG9ydCBjb25zdCBkZWNyZW1lbnQgPSBjcmVhdGVBY3Rpb24oXCJjb3VudGVyL2RlY3JlbWVudFwiKTtcbmV4cG9ydCBjb25zdCBpbmNyZW1lbnRCeUFtb3VudCA9IGNyZWF0ZUFjdGlvbjxudW1iZXI+KFxuICBcImNvdW50ZXIvaW5jcmVtZW50QnlBbW91bnRcIlxuKTtcbiJdLCJuYW1lcyI6WyJjcmVhdGVBY3Rpb24iLCJpbmNyZW1lbnQiLCJkZWNyZW1lbnQiLCJpbmNyZW1lbnRCeUFtb3VudCJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./src/features/counter/actions.ts\n");

/***/ }),

/***/ "./src/features/counter/index.ts":
/*!***************************************!*\
  !*** ./src/features/counter/index.ts ***!
  \***************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var _actions__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./actions */ \"./src/features/counter/actions.ts\");\n/* harmony reexport (unknown) */ var __WEBPACK_REEXPORT_OBJECT__ = {};\n/* harmony reexport (unknown) */ for(const __WEBPACK_IMPORT_KEY__ in _actions__WEBPACK_IMPORTED_MODULE_0__) if(__WEBPACK_IMPORT_KEY__ !== \"default\") __WEBPACK_REEXPORT_OBJECT__[__WEBPACK_IMPORT_KEY__] = () => _actions__WEBPACK_IMPORTED_MODULE_0__[__WEBPACK_IMPORT_KEY__]\n/* harmony reexport (unknown) */ __webpack_require__.d(__webpack_exports__, __WEBPACK_REEXPORT_OBJECT__);\n/* harmony import */ var _reducer__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./reducer */ \"./src/features/counter/reducer.ts\");\n/* harmony reexport (unknown) */ var __WEBPACK_REEXPORT_OBJECT__ = {};\n/* harmony reexport (unknown) */ for(const __WEBPACK_IMPORT_KEY__ in _reducer__WEBPACK_IMPORTED_MODULE_1__) if(__WEBPACK_IMPORT_KEY__ !== \"default\") __WEBPACK_REEXPORT_OBJECT__[__WEBPACK_IMPORT_KEY__] = () => _reducer__WEBPACK_IMPORTED_MODULE_1__[__WEBPACK_IMPORT_KEY__]\n/* harmony reexport (unknown) */ __webpack_require__.d(__webpack_exports__, __WEBPACK_REEXPORT_OBJECT__);\n/* harmony import */ var _selectors__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./selectors */ \"./src/features/counter/selectors.ts\");\n/* harmony reexport (unknown) */ var __WEBPACK_REEXPORT_OBJECT__ = {};\n/* harmony reexport (unknown) */ for(const __WEBPACK_IMPORT_KEY__ in _selectors__WEBPACK_IMPORTED_MODULE_2__) if(__WEBPACK_IMPORT_KEY__ !== \"default\") __WEBPACK_REEXPORT_OBJECT__[__WEBPACK_IMPORT_KEY__] = () => _selectors__WEBPACK_IMPORTED_MODULE_2__[__WEBPACK_IMPORT_KEY__]\n/* harmony reexport (unknown) */ __webpack_require__.d(__webpack_exports__, __WEBPACK_REEXPORT_OBJECT__);\n\n\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvZmVhdHVyZXMvY291bnRlci9pbmRleC50cy5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7Ozs7Ozs7O0FBQUE7QUFDQSIsInNvdXJjZXMiOlsid2VicGFjazovLy8uL3NyYy9mZWF0dXJlcy9jb3VudGVyL2luZGV4LnRzPzhlNmMiXSwic291cmNlc0NvbnRlbnQiOlsiZXhwb3J0ICogZnJvbSBcIi4vYWN0aW9uc1wiO1xuZXhwb3J0ICogZnJvbSBcIi4vcmVkdWNlclwiO1xuZXhwb3J0ICogZnJvbSBcIi4vc2VsZWN0b3JzXCI7XG4iXSwibmFtZXMiOltdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./src/features/counter/index.ts\n");

/***/ }),

/***/ "./src/features/counter/reducer.ts":
/*!*****************************************!*\
  !*** ./src/features/counter/reducer.ts ***!
  \*****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"counterReducer\": () => (/* binding */ counterReducer)\n/* harmony export */ });\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @reduxjs/toolkit */ \"@reduxjs/toolkit\");\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var _actions__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./actions */ \"./src/features/counter/actions.ts\");\n\n\nconst initialState = {\n  value: 0\n};\nconst counterReducer = (0,_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__.createReducer)(initialState, builder => {\n  builder.addCase(_actions__WEBPACK_IMPORTED_MODULE_1__.increment, state => {\n    state.value++;\n  }).addCase(_actions__WEBPACK_IMPORTED_MODULE_1__.decrement, state => {\n    state.value--;\n  }).addCase(_actions__WEBPACK_IMPORTED_MODULE_1__.incrementByAmount, (state, action) => {\n    state.value += action.payload;\n  });\n});//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvZmVhdHVyZXMvY291bnRlci9yZWR1Y2VyLnRzLmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7QUFBQTtBQUNBO0FBTUEsTUFBTUksWUFBMEIsR0FBRztBQUNqQ0MsRUFBQUEsS0FBSyxFQUFFO0FBRDBCLENBQW5DO0FBSU8sTUFBTUMsY0FBYyxHQUFHTiwrREFBYSxDQUFDSSxZQUFELEVBQWdCRyxPQUFELElBQWE7QUFDckVBLEVBQUFBLE9BQU8sQ0FDSkMsT0FESCxDQUNXTiwrQ0FEWCxFQUN1Qk8sS0FBRCxJQUFXO0FBQzdCQSxJQUFBQSxLQUFLLENBQUNKLEtBQU47QUFDRCxHQUhILEVBSUdHLE9BSkgsQ0FJV1AsK0NBSlgsRUFJdUJRLEtBQUQsSUFBVztBQUM3QkEsSUFBQUEsS0FBSyxDQUFDSixLQUFOO0FBQ0QsR0FOSCxFQU9HRyxPQVBILENBT1dMLHVEQVBYLEVBTzhCLENBQUNNLEtBQUQsRUFBUUMsTUFBUixLQUFtQjtBQUM3Q0QsSUFBQUEsS0FBSyxDQUFDSixLQUFOLElBQWVLLE1BQU0sQ0FBQ0MsT0FBdEI7QUFDRCxHQVRIO0FBVUQsQ0FYMEMsQ0FBcEMiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9zcmMvZmVhdHVyZXMvY291bnRlci9yZWR1Y2VyLnRzP2U0NDkiXSwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgY3JlYXRlUmVkdWNlciB9IGZyb20gXCJAcmVkdXhqcy90b29sa2l0XCI7XG5pbXBvcnQgeyBkZWNyZW1lbnQsIGluY3JlbWVudCwgaW5jcmVtZW50QnlBbW91bnQgfSBmcm9tIFwiLi9hY3Rpb25zXCI7XG5cbnR5cGUgQ291bnRlclN0YXRlID0ge1xuICB2YWx1ZTogbnVtYmVyO1xufTtcblxuY29uc3QgaW5pdGlhbFN0YXRlOiBDb3VudGVyU3RhdGUgPSB7XG4gIHZhbHVlOiAwLFxufTtcblxuZXhwb3J0IGNvbnN0IGNvdW50ZXJSZWR1Y2VyID0gY3JlYXRlUmVkdWNlcihpbml0aWFsU3RhdGUsIChidWlsZGVyKSA9PiB7XG4gIGJ1aWxkZXJcbiAgICAuYWRkQ2FzZShpbmNyZW1lbnQsIChzdGF0ZSkgPT4ge1xuICAgICAgc3RhdGUudmFsdWUrKztcbiAgICB9KVxuICAgIC5hZGRDYXNlKGRlY3JlbWVudCwgKHN0YXRlKSA9PiB7XG4gICAgICBzdGF0ZS52YWx1ZS0tO1xuICAgIH0pXG4gICAgLmFkZENhc2UoaW5jcmVtZW50QnlBbW91bnQsIChzdGF0ZSwgYWN0aW9uKSA9PiB7XG4gICAgICBzdGF0ZS52YWx1ZSArPSBhY3Rpb24ucGF5bG9hZDtcbiAgICB9KTtcbn0pO1xuIl0sIm5hbWVzIjpbImNyZWF0ZVJlZHVjZXIiLCJkZWNyZW1lbnQiLCJpbmNyZW1lbnQiLCJpbmNyZW1lbnRCeUFtb3VudCIsImluaXRpYWxTdGF0ZSIsInZhbHVlIiwiY291bnRlclJlZHVjZXIiLCJidWlsZGVyIiwiYWRkQ2FzZSIsInN0YXRlIiwiYWN0aW9uIiwicGF5bG9hZCJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./src/features/counter/reducer.ts\n");

/***/ }),

/***/ "./src/features/counter/selectors.ts":
/*!*******************************************!*\
  !*** ./src/features/counter/selectors.ts ***!
  \*******************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"selectCount\": () => (/* binding */ selectCount),\n/* harmony export */   \"countSelector\": () => (/* binding */ countSelector)\n/* harmony export */ });\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @reduxjs/toolkit */ \"@reduxjs/toolkit\");\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__);\n\nconst selectCount = state => state.counter.value;\nconst countSelector = (0,_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__.createSelector)(selectCount, state => state);//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvZmVhdHVyZXMvY291bnRlci9zZWxlY3RvcnMudHMuanMiLCJtYXBwaW5ncyI6Ijs7Ozs7OztBQUFBO0FBR08sTUFBTUMsV0FBVyxHQUFJQyxLQUFELElBQXNCQSxLQUFLLENBQUNDLE9BQU4sQ0FBY0MsS0FBeEQ7QUFFQSxNQUFNQyxhQUFhLEdBQUdMLGdFQUFjLENBQUNDLFdBQUQsRUFBZUMsS0FBRCxJQUFXQSxLQUF6QixDQUFwQyIsInNvdXJjZXMiOlsid2VicGFjazovLy8uL3NyYy9mZWF0dXJlcy9jb3VudGVyL3NlbGVjdG9ycy50cz8yOGNhIl0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IGNyZWF0ZVNlbGVjdG9yIH0gZnJvbSBcIkByZWR1eGpzL3Rvb2xraXRcIjtcbmltcG9ydCB7IFJvb3RTdGF0ZSB9IGZyb20gXCIuLi8uLi9hcHAvc3RvcmVcIjtcblxuZXhwb3J0IGNvbnN0IHNlbGVjdENvdW50ID0gKHN0YXRlOiBSb290U3RhdGUpID0+IHN0YXRlLmNvdW50ZXIudmFsdWU7XG5cbmV4cG9ydCBjb25zdCBjb3VudFNlbGVjdG9yID0gY3JlYXRlU2VsZWN0b3Ioc2VsZWN0Q291bnQsIChzdGF0ZSkgPT4gc3RhdGUpO1xuIl0sIm5hbWVzIjpbImNyZWF0ZVNlbGVjdG9yIiwic2VsZWN0Q291bnQiLCJzdGF0ZSIsImNvdW50ZXIiLCJ2YWx1ZSIsImNvdW50U2VsZWN0b3IiXSwic291cmNlUm9vdCI6IiJ9\n//# sourceURL=webpack-internal:///./src/features/counter/selectors.ts\n");

/***/ }),

/***/ "./src/pages/index.tsx":
/*!*****************************!*\
  !*** ./src/pages/index.tsx ***!
  \*****************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"default\": () => (__WEBPACK_DEFAULT_EXPORT__)\n/* harmony export */ });\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"react/jsx-dev-runtime\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"react\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var _app_hooks__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../app/hooks */ \"./src/app/hooks.ts\");\n/* harmony import */ var _features_counter__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../features/counter */ \"./src/features/counter/index.ts\");\n\n\nvar _jsxFileName = \"/Users/leo/workspace/notion_app/client/src/pages/index.tsx\";\n\n\n\n\nconst IndexPage = () => {\n  const dispatch = (0,_app_hooks__WEBPACK_IMPORTED_MODULE_2__.useAppDispatch)();\n  const count = (0,_app_hooks__WEBPACK_IMPORTED_MODULE_2__.useAppSelector)(_features_counter__WEBPACK_IMPORTED_MODULE_3__.selectCount);\n  const {\n    0: incrementAmount,\n    1: setIncrementAmount\n  } = (0,react__WEBPACK_IMPORTED_MODULE_1__.useState)(0);\n  return /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.Fragment, {\n    children: [/*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h1\", {\n      children: \"Welcome to the greatest app in the world!\"\n    }, void 0, false, {\n      fileName: _jsxFileName,\n      lineNumber: 17,\n      columnNumber: 7\n    }, undefined), /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h2\", {\n      children: [\"The current number is\", count]\n    }, void 0, true, {\n      fileName: _jsxFileName,\n      lineNumber: 18,\n      columnNumber: 7\n    }, undefined), /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n      children: [/*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"input\", {\n        value: incrementAmount,\n        onChange: e => setIncrementAmount(Number(e.target.value)),\n        type: \"number\"\n      }, void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 23,\n        columnNumber: 9\n      }, undefined), /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"button\", {\n        onClick: () => dispatch((0,_features_counter__WEBPACK_IMPORTED_MODULE_3__.incrementByAmount)(Number(incrementAmount))),\n        children: \"Increment by amount\"\n      }, void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 28,\n        columnNumber: 9\n      }, undefined)]\n    }, void 0, true, {\n      fileName: _jsxFileName,\n      lineNumber: 22,\n      columnNumber: 7\n    }, undefined), /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n      children: [/*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"button\", {\n        onClick: () => dispatch((0,_features_counter__WEBPACK_IMPORTED_MODULE_3__.decrement)()),\n        children: \"Decrement by 1\"\n      }, void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 35,\n        columnNumber: 9\n      }, undefined), /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"button\", {\n        onClick: () => dispatch((0,_features_counter__WEBPACK_IMPORTED_MODULE_3__.increment)()),\n        children: \"Increment by 1\"\n      }, void 0, false, {\n        fileName: _jsxFileName,\n        lineNumber: 36,\n        columnNumber: 9\n      }, undefined)]\n    }, void 0, true, {\n      fileName: _jsxFileName,\n      lineNumber: 34,\n      columnNumber: 7\n    }, undefined)]\n  }, void 0, true);\n};\n\n/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (IndexPage);//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvcGFnZXMvaW5kZXgudHN4LmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7Ozs7QUFBQTtBQUNBO0FBQ0E7O0FBT0EsTUFBTU8sU0FBbUIsR0FBRyxNQUFNO0FBQ2hDLFFBQU1DLFFBQVEsR0FBR1AsMERBQWMsRUFBL0I7QUFDQSxRQUFNUSxLQUFLLEdBQUdQLDBEQUFjLENBQUNJLDBEQUFELENBQTVCO0FBQ0EsUUFBTTtBQUFBLE9BQUNJLGVBQUQ7QUFBQSxPQUFrQkM7QUFBbEIsTUFBd0NYLCtDQUFRLENBQVMsQ0FBVCxDQUF0RDtBQUVBLHNCQUNFO0FBQUEsNEJBQ0U7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsaUJBREYsZUFFRTtBQUFBLDBDQUVHUyxLQUZIO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxpQkFGRixlQU1FO0FBQUEsOEJBQ0U7QUFDRSxhQUFLLEVBQUVDLGVBRFQ7QUFFRSxnQkFBUSxFQUFHRSxDQUFELElBQU9ELGtCQUFrQixDQUFDRSxNQUFNLENBQUNELENBQUMsQ0FBQ0UsTUFBRixDQUFTQyxLQUFWLENBQVAsQ0FGckM7QUFHRSxZQUFJLEVBQUM7QUFIUDtBQUFBO0FBQUE7QUFBQTtBQUFBLG1CQURGLGVBTUU7QUFDRSxlQUFPLEVBQUUsTUFBTVAsUUFBUSxDQUFDSCxvRUFBaUIsQ0FBQ1EsTUFBTSxDQUFDSCxlQUFELENBQVAsQ0FBbEIsQ0FEekI7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsbUJBTkY7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLGlCQU5GLGVBa0JFO0FBQUEsOEJBQ0U7QUFBUSxlQUFPLEVBQUUsTUFBTUYsUUFBUSxDQUFDTCw0REFBUyxFQUFWLENBQS9CO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLG1CQURGLGVBRUU7QUFBUSxlQUFPLEVBQUUsTUFBTUssUUFBUSxDQUFDSiw0REFBUyxFQUFWLENBQS9CO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLG1CQUZGO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxpQkFsQkY7QUFBQSxrQkFERjtBQXlCRCxDQTlCRDs7QUFnQ0EsaUVBQWVHLFNBQWYiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9zcmMvcGFnZXMvaW5kZXgudHN4PzQxZTAiXSwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgdXNlU3RhdGUgfSBmcm9tIFwicmVhY3RcIjtcbmltcG9ydCB7IHVzZUFwcERpc3BhdGNoLCB1c2VBcHBTZWxlY3RvciB9IGZyb20gXCIuLi9hcHAvaG9va3NcIjtcbmltcG9ydCB7XG4gIGRlY3JlbWVudCxcbiAgaW5jcmVtZW50LFxuICBpbmNyZW1lbnRCeUFtb3VudCxcbiAgc2VsZWN0Q291bnQsXG59IGZyb20gXCIuLi9mZWF0dXJlcy9jb3VudGVyXCI7XG5cbmNvbnN0IEluZGV4UGFnZTogUmVhY3QuRkMgPSAoKSA9PiB7XG4gIGNvbnN0IGRpc3BhdGNoID0gdXNlQXBwRGlzcGF0Y2goKTtcbiAgY29uc3QgY291bnQgPSB1c2VBcHBTZWxlY3RvcihzZWxlY3RDb3VudCk7XG4gIGNvbnN0IFtpbmNyZW1lbnRBbW91bnQsIHNldEluY3JlbWVudEFtb3VudF0gPSB1c2VTdGF0ZTxudW1iZXI+KDApO1xuXG4gIHJldHVybiAoXG4gICAgPD5cbiAgICAgIDxoMT5XZWxjb21lIHRvIHRoZSBncmVhdGVzdCBhcHAgaW4gdGhlIHdvcmxkITwvaDE+XG4gICAgICA8aDI+XG4gICAgICAgIFRoZSBjdXJyZW50IG51bWJlciBpc1xuICAgICAgICB7Y291bnR9XG4gICAgICA8L2gyPlxuICAgICAgPGRpdj5cbiAgICAgICAgPGlucHV0XG4gICAgICAgICAgdmFsdWU9e2luY3JlbWVudEFtb3VudH1cbiAgICAgICAgICBvbkNoYW5nZT17KGUpID0+IHNldEluY3JlbWVudEFtb3VudChOdW1iZXIoZS50YXJnZXQudmFsdWUpKX1cbiAgICAgICAgICB0eXBlPVwibnVtYmVyXCJcbiAgICAgICAgLz5cbiAgICAgICAgPGJ1dHRvblxuICAgICAgICAgIG9uQ2xpY2s9eygpID0+IGRpc3BhdGNoKGluY3JlbWVudEJ5QW1vdW50KE51bWJlcihpbmNyZW1lbnRBbW91bnQpKSl9XG4gICAgICAgID5cbiAgICAgICAgICBJbmNyZW1lbnQgYnkgYW1vdW50XG4gICAgICAgIDwvYnV0dG9uPlxuICAgICAgPC9kaXY+XG4gICAgICA8ZGl2PlxuICAgICAgICA8YnV0dG9uIG9uQ2xpY2s9eygpID0+IGRpc3BhdGNoKGRlY3JlbWVudCgpKX0+RGVjcmVtZW50IGJ5IDE8L2J1dHRvbj5cbiAgICAgICAgPGJ1dHRvbiBvbkNsaWNrPXsoKSA9PiBkaXNwYXRjaChpbmNyZW1lbnQoKSl9PkluY3JlbWVudCBieSAxPC9idXR0b24+XG4gICAgICA8L2Rpdj5cbiAgICA8Lz5cbiAgKTtcbn07XG5cbmV4cG9ydCBkZWZhdWx0IEluZGV4UGFnZTtcbiJdLCJuYW1lcyI6WyJ1c2VTdGF0ZSIsInVzZUFwcERpc3BhdGNoIiwidXNlQXBwU2VsZWN0b3IiLCJkZWNyZW1lbnQiLCJpbmNyZW1lbnQiLCJpbmNyZW1lbnRCeUFtb3VudCIsInNlbGVjdENvdW50IiwiSW5kZXhQYWdlIiwiZGlzcGF0Y2giLCJjb3VudCIsImluY3JlbWVudEFtb3VudCIsInNldEluY3JlbWVudEFtb3VudCIsImUiLCJOdW1iZXIiLCJ0YXJnZXQiLCJ2YWx1ZSJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./src/pages/index.tsx\n");

/***/ }),

/***/ "@reduxjs/toolkit":
/*!***********************************!*\
  !*** external "@reduxjs/toolkit" ***!
  \***********************************/
/***/ ((module) => {

module.exports = require("@reduxjs/toolkit");

/***/ }),

/***/ "react":
/*!************************!*\
  !*** external "react" ***!
  \************************/
/***/ ((module) => {

module.exports = require("react");

/***/ }),

/***/ "react-redux":
/*!******************************!*\
  !*** external "react-redux" ***!
  \******************************/
/***/ ((module) => {

module.exports = require("react-redux");

/***/ }),

/***/ "react/jsx-dev-runtime":
/*!****************************************!*\
  !*** external "react/jsx-dev-runtime" ***!
  \****************************************/
/***/ ((module) => {

module.exports = require("react/jsx-dev-runtime");

/***/ })

};
;

// load runtime
var __webpack_require__ = require("../webpack-runtime.js");
__webpack_require__.C(exports);
var __webpack_exec__ = (moduleId) => (__webpack_require__(__webpack_require__.s = moduleId))
var __webpack_exports__ = (__webpack_exec__("./src/pages/index.tsx"));
module.exports = __webpack_exports__;

})();