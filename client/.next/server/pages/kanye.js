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
exports.id = "pages/kanye";
exports.ids = ["pages/kanye"];
exports.modules = {

/***/ "./src/app/hooks.ts":
/*!**************************!*\
  !*** ./src/app/hooks.ts ***!
  \**************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"useAppDispatch\": () => (/* binding */ useAppDispatch),\n/* harmony export */   \"useAppSelector\": () => (/* binding */ useAppSelector)\n/* harmony export */ });\n/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react-redux */ \"react-redux\");\n/* harmony import */ var react_redux__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_redux__WEBPACK_IMPORTED_MODULE_0__);\n\nconst useAppDispatch = () => (0,react_redux__WEBPACK_IMPORTED_MODULE_0__.useDispatch)();\nconst useAppSelector = react_redux__WEBPACK_IMPORTED_MODULE_0__.useSelector;//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvYXBwL2hvb2tzLnRzLmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7QUFBQTtBQUdPLE1BQU1FLGNBQWMsR0FBRyxNQUFNRix3REFBVyxFQUF4QztBQUNBLE1BQU1HLGNBQStDLEdBQUdGLG9EQUF4RCIsInNvdXJjZXMiOlsid2VicGFjazovLy8uL3NyYy9hcHAvaG9va3MudHM/NTU1ZiJdLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBUeXBlZFVzZVNlbGVjdG9ySG9vaywgdXNlRGlzcGF0Y2gsIHVzZVNlbGVjdG9yIH0gZnJvbSBcInJlYWN0LXJlZHV4XCI7XG5pbXBvcnQgdHlwZSB7IEFwcERpc3BhdGNoLCBSb290U3RhdGUgfSBmcm9tIFwiLi9zdG9yZVwiO1xuXG5leHBvcnQgY29uc3QgdXNlQXBwRGlzcGF0Y2ggPSAoKSA9PiB1c2VEaXNwYXRjaDxBcHBEaXNwYXRjaD4oKTtcbmV4cG9ydCBjb25zdCB1c2VBcHBTZWxlY3RvcjogVHlwZWRVc2VTZWxlY3Rvckhvb2s8Um9vdFN0YXRlPiA9IHVzZVNlbGVjdG9yO1xuIl0sIm5hbWVzIjpbInVzZURpc3BhdGNoIiwidXNlU2VsZWN0b3IiLCJ1c2VBcHBEaXNwYXRjaCIsInVzZUFwcFNlbGVjdG9yIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./src/app/hooks.ts\n");

/***/ }),

/***/ "./src/features/kanye/actions.ts":
/*!***************************************!*\
  !*** ./src/features/kanye/actions.ts ***!
  \***************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"getKanyeQuote\": () => (/* binding */ getKanyeQuote)\n/* harmony export */ });\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @reduxjs/toolkit */ \"@reduxjs/toolkit\");\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! axios */ \"axios\");\n/* harmony import */ var axios__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(axios__WEBPACK_IMPORTED_MODULE_1__);\n\n\nconst getKanyeQuote = (0,_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__.createAsyncThunk)(\"kanye/kanyeQuote\", async () => {\n  const response = await axios__WEBPACK_IMPORTED_MODULE_1___default().get(\"https://api.kanye.rest/\");\n  return response.data;\n});//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvZmVhdHVyZXMva2FueWUvYWN0aW9ucy50cy5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7OztBQUFBO0FBQ0E7QUFFTyxNQUFNRSxhQUFhLEdBQUdGLGtFQUFnQixDQUFDLGtCQUFELEVBQXFCLFlBQVk7QUFDNUUsUUFBTUcsUUFBUSxHQUFHLE1BQU1GLGdEQUFBLENBQVUseUJBQVYsQ0FBdkI7QUFFQSxTQUFPRSxRQUFRLENBQUNFLElBQWhCO0FBQ0QsQ0FKNEMsQ0FBdEMiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9zcmMvZmVhdHVyZXMva2FueWUvYWN0aW9ucy50cz9mOTM1Il0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IGNyZWF0ZUFzeW5jVGh1bmsgfSBmcm9tIFwiQHJlZHV4anMvdG9vbGtpdFwiO1xuaW1wb3J0IGF4aW9zIGZyb20gXCJheGlvc1wiO1xuXG5leHBvcnQgY29uc3QgZ2V0S2FueWVRdW90ZSA9IGNyZWF0ZUFzeW5jVGh1bmsoXCJrYW55ZS9rYW55ZVF1b3RlXCIsIGFzeW5jICgpID0+IHtcbiAgY29uc3QgcmVzcG9uc2UgPSBhd2FpdCBheGlvcy5nZXQoXCJodHRwczovL2FwaS5rYW55ZS5yZXN0L1wiKTtcblxuICByZXR1cm4gcmVzcG9uc2UuZGF0YTtcbn0pO1xuIl0sIm5hbWVzIjpbImNyZWF0ZUFzeW5jVGh1bmsiLCJheGlvcyIsImdldEthbnllUXVvdGUiLCJyZXNwb25zZSIsImdldCIsImRhdGEiXSwic291cmNlUm9vdCI6IiJ9\n//# sourceURL=webpack-internal:///./src/features/kanye/actions.ts\n");

/***/ }),

/***/ "./src/features/kanye/index.ts":
/*!*************************************!*\
  !*** ./src/features/kanye/index.ts ***!
  \*************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var _actions__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./actions */ \"./src/features/kanye/actions.ts\");\n/* harmony reexport (unknown) */ var __WEBPACK_REEXPORT_OBJECT__ = {};\n/* harmony reexport (unknown) */ for(const __WEBPACK_IMPORT_KEY__ in _actions__WEBPACK_IMPORTED_MODULE_0__) if(__WEBPACK_IMPORT_KEY__ !== \"default\") __WEBPACK_REEXPORT_OBJECT__[__WEBPACK_IMPORT_KEY__] = () => _actions__WEBPACK_IMPORTED_MODULE_0__[__WEBPACK_IMPORT_KEY__]\n/* harmony reexport (unknown) */ __webpack_require__.d(__webpack_exports__, __WEBPACK_REEXPORT_OBJECT__);\n/* harmony import */ var _reducer__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./reducer */ \"./src/features/kanye/reducer.ts\");\n/* harmony reexport (unknown) */ var __WEBPACK_REEXPORT_OBJECT__ = {};\n/* harmony reexport (unknown) */ for(const __WEBPACK_IMPORT_KEY__ in _reducer__WEBPACK_IMPORTED_MODULE_1__) if(__WEBPACK_IMPORT_KEY__ !== \"default\") __WEBPACK_REEXPORT_OBJECT__[__WEBPACK_IMPORT_KEY__] = () => _reducer__WEBPACK_IMPORTED_MODULE_1__[__WEBPACK_IMPORT_KEY__]\n/* harmony reexport (unknown) */ __webpack_require__.d(__webpack_exports__, __WEBPACK_REEXPORT_OBJECT__);\n/* harmony import */ var _selectors__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./selectors */ \"./src/features/kanye/selectors.ts\");\n/* harmony reexport (unknown) */ var __WEBPACK_REEXPORT_OBJECT__ = {};\n/* harmony reexport (unknown) */ for(const __WEBPACK_IMPORT_KEY__ in _selectors__WEBPACK_IMPORTED_MODULE_2__) if(__WEBPACK_IMPORT_KEY__ !== \"default\") __WEBPACK_REEXPORT_OBJECT__[__WEBPACK_IMPORT_KEY__] = () => _selectors__WEBPACK_IMPORTED_MODULE_2__[__WEBPACK_IMPORT_KEY__]\n/* harmony reexport (unknown) */ __webpack_require__.d(__webpack_exports__, __WEBPACK_REEXPORT_OBJECT__);\n\n\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvZmVhdHVyZXMva2FueWUvaW5kZXgudHMuanMiLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7OztBQUFBO0FBQ0EiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9zcmMvZmVhdHVyZXMva2FueWUvaW5kZXgudHM/NWJkNCJdLCJzb3VyY2VzQ29udGVudCI6WyJleHBvcnQgKiBmcm9tIFwiLi9hY3Rpb25zXCI7XG5leHBvcnQgKiBmcm9tIFwiLi9yZWR1Y2VyXCI7XG5leHBvcnQgKiBmcm9tIFwiLi9zZWxlY3RvcnNcIjtcbiJdLCJuYW1lcyI6W10sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./src/features/kanye/index.ts\n");

/***/ }),

/***/ "./src/features/kanye/reducer.ts":
/*!***************************************!*\
  !*** ./src/features/kanye/reducer.ts ***!
  \***************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"kanyeReducer\": () => (/* binding */ kanyeReducer),\n/* harmony export */   \"default\": () => (__WEBPACK_DEFAULT_EXPORT__)\n/* harmony export */ });\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @reduxjs/toolkit */ \"@reduxjs/toolkit\");\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var _actions__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./actions */ \"./src/features/kanye/actions.ts\");\n\n\nconst initialState = {\n  data: {\n    quote: \"click that button\"\n  },\n  pending: false,\n  error: false\n};\nconst kanyeReducer = (0,_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__.createReducer)(initialState, builder => {\n  builder.addCase(_actions__WEBPACK_IMPORTED_MODULE_1__.getKanyeQuote.pending, state => {\n    state.pending = true;\n  }).addCase(_actions__WEBPACK_IMPORTED_MODULE_1__.getKanyeQuote.fulfilled, (state, {\n    payload\n  }) => {\n    state.pending = false;\n    state.data = payload;\n  }).addCase(_actions__WEBPACK_IMPORTED_MODULE_1__.getKanyeQuote.rejected, state => {\n    state.pending = false;\n    state.error = true;\n  });\n});\n/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (kanyeReducer);//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvZmVhdHVyZXMva2FueWUvcmVkdWNlci50cy5qcyIsIm1hcHBpbmdzIjoiOzs7Ozs7OztBQUFBO0FBQ0E7QUFRQSxNQUFNRSxZQUF3QixHQUFHO0FBQy9CQyxFQUFBQSxJQUFJLEVBQUU7QUFBRUMsSUFBQUEsS0FBSyxFQUFFO0FBQVQsR0FEeUI7QUFFL0JDLEVBQUFBLE9BQU8sRUFBRSxLQUZzQjtBQUcvQkMsRUFBQUEsS0FBSyxFQUFFO0FBSHdCLENBQWpDO0FBTU8sTUFBTUMsWUFBWSxHQUFHUCwrREFBYSxDQUFDRSxZQUFELEVBQWdCTSxPQUFELElBQWE7QUFDbkVBLEVBQUFBLE9BQU8sQ0FDSkMsT0FESCxDQUNXUiwyREFEWCxFQUNtQ1MsS0FBRCxJQUFXO0FBQ3pDQSxJQUFBQSxLQUFLLENBQUNMLE9BQU4sR0FBZ0IsSUFBaEI7QUFDRCxHQUhILEVBSUdJLE9BSkgsQ0FJV1IsNkRBSlgsRUFJb0MsQ0FBQ1MsS0FBRCxFQUFRO0FBQUVFLElBQUFBO0FBQUYsR0FBUixLQUF3QjtBQUN4REYsSUFBQUEsS0FBSyxDQUFDTCxPQUFOLEdBQWdCLEtBQWhCO0FBQ0FLLElBQUFBLEtBQUssQ0FBQ1AsSUFBTixHQUFhUyxPQUFiO0FBQ0QsR0FQSCxFQVFHSCxPQVJILENBUVdSLDREQVJYLEVBUW9DUyxLQUFELElBQVc7QUFDMUNBLElBQUFBLEtBQUssQ0FBQ0wsT0FBTixHQUFnQixLQUFoQjtBQUNBSyxJQUFBQSxLQUFLLENBQUNKLEtBQU4sR0FBYyxJQUFkO0FBQ0QsR0FYSDtBQVlELENBYndDLENBQWxDO0FBZVAsaUVBQWVDLFlBQWYiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9zcmMvZmVhdHVyZXMva2FueWUvcmVkdWNlci50cz82OTVhIl0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IGNyZWF0ZVJlZHVjZXIgfSBmcm9tIFwiQHJlZHV4anMvdG9vbGtpdFwiO1xuaW1wb3J0IHsgZ2V0S2FueWVRdW90ZSB9IGZyb20gXCIuL2FjdGlvbnNcIjtcblxuZXhwb3J0IHR5cGUgS2FueWVTdGF0ZSA9IHtcbiAgZGF0YTogeyBxdW90ZTogc3RyaW5nIH07XG4gIHBlbmRpbmc6IGJvb2xlYW47XG4gIGVycm9yOiBib29sZWFuO1xufTtcblxuY29uc3QgaW5pdGlhbFN0YXRlOiBLYW55ZVN0YXRlID0ge1xuICBkYXRhOiB7IHF1b3RlOiBcImNsaWNrIHRoYXQgYnV0dG9uXCIgfSxcbiAgcGVuZGluZzogZmFsc2UsXG4gIGVycm9yOiBmYWxzZSxcbn07XG5cbmV4cG9ydCBjb25zdCBrYW55ZVJlZHVjZXIgPSBjcmVhdGVSZWR1Y2VyKGluaXRpYWxTdGF0ZSwgKGJ1aWxkZXIpID0+IHtcbiAgYnVpbGRlclxuICAgIC5hZGRDYXNlKGdldEthbnllUXVvdGUucGVuZGluZywgKHN0YXRlKSA9PiB7XG4gICAgICBzdGF0ZS5wZW5kaW5nID0gdHJ1ZTtcbiAgICB9KVxuICAgIC5hZGRDYXNlKGdldEthbnllUXVvdGUuZnVsZmlsbGVkLCAoc3RhdGUsIHsgcGF5bG9hZCB9KSA9PiB7XG4gICAgICBzdGF0ZS5wZW5kaW5nID0gZmFsc2U7XG4gICAgICBzdGF0ZS5kYXRhID0gcGF5bG9hZDtcbiAgICB9KVxuICAgIC5hZGRDYXNlKGdldEthbnllUXVvdGUucmVqZWN0ZWQsIChzdGF0ZSkgPT4ge1xuICAgICAgc3RhdGUucGVuZGluZyA9IGZhbHNlO1xuICAgICAgc3RhdGUuZXJyb3IgPSB0cnVlO1xuICAgIH0pO1xufSk7XG5cbmV4cG9ydCBkZWZhdWx0IGthbnllUmVkdWNlcjtcbiJdLCJuYW1lcyI6WyJjcmVhdGVSZWR1Y2VyIiwiZ2V0S2FueWVRdW90ZSIsImluaXRpYWxTdGF0ZSIsImRhdGEiLCJxdW90ZSIsInBlbmRpbmciLCJlcnJvciIsImthbnllUmVkdWNlciIsImJ1aWxkZXIiLCJhZGRDYXNlIiwic3RhdGUiLCJmdWxmaWxsZWQiLCJwYXlsb2FkIiwicmVqZWN0ZWQiXSwic291cmNlUm9vdCI6IiJ9\n//# sourceURL=webpack-internal:///./src/features/kanye/reducer.ts\n");

/***/ }),

/***/ "./src/features/kanye/selectors.ts":
/*!*****************************************!*\
  !*** ./src/features/kanye/selectors.ts ***!
  \*****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"selectQuote\": () => (/* binding */ selectQuote),\n/* harmony export */   \"kanyeQuoteSelector\": () => (/* binding */ kanyeQuoteSelector)\n/* harmony export */ });\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @reduxjs/toolkit */ \"@reduxjs/toolkit\");\n/* harmony import */ var _reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__);\n\nconst selectQuote = state => state.kanyeQuote;\nconst kanyeQuoteSelector = (0,_reduxjs_toolkit__WEBPACK_IMPORTED_MODULE_0__.createSelector)(selectQuote, state => state);//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvZmVhdHVyZXMva2FueWUvc2VsZWN0b3JzLnRzLmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7QUFDQTtBQUVPLE1BQU1DLFdBQVcsR0FBSUMsS0FBRCxJQUFzQkEsS0FBSyxDQUFDQyxVQUFoRDtBQUVBLE1BQU1DLGtCQUFrQixHQUFHSixnRUFBYyxDQUFDQyxXQUFELEVBQWVDLEtBQUQsSUFBV0EsS0FBekIsQ0FBekMiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vLi9zcmMvZmVhdHVyZXMva2FueWUvc2VsZWN0b3JzLnRzPzM3YzciXSwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHsgUm9vdFN0YXRlIH0gZnJvbSBcIi4uLy4uL2FwcC9zdG9yZVwiO1xuaW1wb3J0IHsgY3JlYXRlU2VsZWN0b3IgfSBmcm9tIFwiQHJlZHV4anMvdG9vbGtpdFwiO1xuXG5leHBvcnQgY29uc3Qgc2VsZWN0UXVvdGUgPSAoc3RhdGU6IFJvb3RTdGF0ZSkgPT4gc3RhdGUua2FueWVRdW90ZTtcblxuZXhwb3J0IGNvbnN0IGthbnllUXVvdGVTZWxlY3RvciA9IGNyZWF0ZVNlbGVjdG9yKHNlbGVjdFF1b3RlLCAoc3RhdGUpID0+IHN0YXRlKTtcbiJdLCJuYW1lcyI6WyJjcmVhdGVTZWxlY3RvciIsInNlbGVjdFF1b3RlIiwic3RhdGUiLCJrYW55ZVF1b3RlIiwia2FueWVRdW90ZVNlbGVjdG9yIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///./src/features/kanye/selectors.ts\n");

/***/ }),

/***/ "./src/pages/kanye/index.tsx":
/*!***********************************!*\
  !*** ./src/pages/kanye/index.tsx ***!
  \***********************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   \"default\": () => (__WEBPACK_DEFAULT_EXPORT__)\n/* harmony export */ });\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"react/jsx-dev-runtime\");\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"react\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var _app_hooks__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../../app/hooks */ \"./src/app/hooks.ts\");\n/* harmony import */ var _features_kanye__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ../../features/kanye */ \"./src/features/kanye/index.ts\");\n\nvar _jsxFileName = \"/Users/leo/workspace/notion_app/client/src/pages/kanye/index.tsx\";\n\n\n\n\nconst kanye = () => {\n  const dispatch = (0,_app_hooks__WEBPACK_IMPORTED_MODULE_2__.useAppDispatch)();\n  const {\n    data,\n    pending,\n    error\n  } = (0,_app_hooks__WEBPACK_IMPORTED_MODULE_2__.useAppSelector)(_features_kanye__WEBPACK_IMPORTED_MODULE_3__.kanyeQuoteSelector);\n  return /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n    children: [/*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h2\", {\n      children: \"Generate random Kanye West quote\"\n    }, void 0, false, {\n      fileName: _jsxFileName,\n      lineNumber: 11,\n      columnNumber: 7\n    }, undefined), pending && /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"p\", {\n      children: \"Loading...\"\n    }, void 0, false, {\n      fileName: _jsxFileName,\n      lineNumber: 12,\n      columnNumber: 19\n    }, undefined), data && /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"p\", {\n      children: data.quote\n    }, void 0, false, {\n      fileName: _jsxFileName,\n      lineNumber: 13,\n      columnNumber: 16\n    }, undefined), error && /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"p\", {\n      children: \"Oops, something went wrong\"\n    }, void 0, false, {\n      fileName: _jsxFileName,\n      lineNumber: 14,\n      columnNumber: 17\n    }, undefined), /*#__PURE__*/(0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"button\", {\n      onClick: () => dispatch((0,_features_kanye__WEBPACK_IMPORTED_MODULE_3__.getKanyeQuote)()),\n      disabled: pending,\n      children: \"Generate Kanye Quote\"\n    }, void 0, false, {\n      fileName: _jsxFileName,\n      lineNumber: 15,\n      columnNumber: 7\n    }, undefined)]\n  }, void 0, true, {\n    fileName: _jsxFileName,\n    lineNumber: 10,\n    columnNumber: 5\n  }, undefined);\n};\n\n/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = (kanye);//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvcGFnZXMva2FueWUvaW5kZXgudHN4LmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7OztBQUFBO0FBQ0E7QUFDQTs7QUFFQSxNQUFNSyxLQUFlLEdBQUcsTUFBTTtBQUM1QixRQUFNQyxRQUFRLEdBQUdMLDBEQUFjLEVBQS9CO0FBQ0EsUUFBTTtBQUFFTSxJQUFBQSxJQUFGO0FBQVFDLElBQUFBLE9BQVI7QUFBaUJDLElBQUFBO0FBQWpCLE1BQTJCUCwwREFBYyxDQUFDRSwrREFBRCxDQUEvQztBQUVBLHNCQUNFO0FBQUEsNEJBQ0U7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsaUJBREYsRUFFR0ksT0FBTyxpQkFBSTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxpQkFGZCxFQUdHRCxJQUFJLGlCQUFJO0FBQUEsZ0JBQUlBLElBQUksQ0FBQ0c7QUFBVDtBQUFBO0FBQUE7QUFBQTtBQUFBLGlCQUhYLEVBSUdELEtBQUssaUJBQUk7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsaUJBSlosZUFLRTtBQUFRLGFBQU8sRUFBRSxNQUFNSCxRQUFRLENBQUNILDhEQUFhLEVBQWQsQ0FBL0I7QUFBa0QsY0FBUSxFQUFFSyxPQUE1RDtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxpQkFMRjtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsZUFERjtBQVdELENBZkQ7O0FBaUJBLGlFQUFlSCxLQUFmIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vLy4vc3JjL3BhZ2VzL2thbnllL2luZGV4LnRzeD9kN2JmIl0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCBSZWFjdCBmcm9tIFwicmVhY3RcIjtcbmltcG9ydCB7IHVzZUFwcERpc3BhdGNoLCB1c2VBcHBTZWxlY3RvciB9IGZyb20gXCIuLi8uLi9hcHAvaG9va3NcIjtcbmltcG9ydCB7IGdldEthbnllUXVvdGUsIGthbnllUXVvdGVTZWxlY3RvciB9IGZyb20gXCIuLi8uLi9mZWF0dXJlcy9rYW55ZVwiO1xuXG5jb25zdCBrYW55ZTogUmVhY3QuRkMgPSAoKSA9PiB7XG4gIGNvbnN0IGRpc3BhdGNoID0gdXNlQXBwRGlzcGF0Y2goKTtcbiAgY29uc3QgeyBkYXRhLCBwZW5kaW5nLCBlcnJvciB9ID0gdXNlQXBwU2VsZWN0b3Ioa2FueWVRdW90ZVNlbGVjdG9yKTtcblxuICByZXR1cm4gKFxuICAgIDxkaXY+XG4gICAgICA8aDI+R2VuZXJhdGUgcmFuZG9tIEthbnllIFdlc3QgcXVvdGU8L2gyPlxuICAgICAge3BlbmRpbmcgJiYgPHA+TG9hZGluZy4uLjwvcD59XG4gICAgICB7ZGF0YSAmJiA8cD57ZGF0YS5xdW90ZX08L3A+fVxuICAgICAge2Vycm9yICYmIDxwPk9vcHMsIHNvbWV0aGluZyB3ZW50IHdyb25nPC9wPn1cbiAgICAgIDxidXR0b24gb25DbGljaz17KCkgPT4gZGlzcGF0Y2goZ2V0S2FueWVRdW90ZSgpKX0gZGlzYWJsZWQ9e3BlbmRpbmd9PlxuICAgICAgICBHZW5lcmF0ZSBLYW55ZSBRdW90ZVxuICAgICAgPC9idXR0b24+XG4gICAgPC9kaXY+XG4gICk7XG59O1xuXG5leHBvcnQgZGVmYXVsdCBrYW55ZTtcbiJdLCJuYW1lcyI6WyJSZWFjdCIsInVzZUFwcERpc3BhdGNoIiwidXNlQXBwU2VsZWN0b3IiLCJnZXRLYW55ZVF1b3RlIiwia2FueWVRdW90ZVNlbGVjdG9yIiwia2FueWUiLCJkaXNwYXRjaCIsImRhdGEiLCJwZW5kaW5nIiwiZXJyb3IiLCJxdW90ZSJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./src/pages/kanye/index.tsx\n");

/***/ }),

/***/ "@reduxjs/toolkit":
/*!***********************************!*\
  !*** external "@reduxjs/toolkit" ***!
  \***********************************/
/***/ ((module) => {

module.exports = require("@reduxjs/toolkit");

/***/ }),

/***/ "axios":
/*!************************!*\
  !*** external "axios" ***!
  \************************/
/***/ ((module) => {

module.exports = require("axios");

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
var __webpack_exports__ = (__webpack_exec__("./src/pages/kanye/index.tsx"));
module.exports = __webpack_exports__;

})();