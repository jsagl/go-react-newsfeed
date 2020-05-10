import './index.css';

import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux';
import { createStore, combineReducers, applyMiddleware } from 'redux';
import { logger } from 'redux-logger';
import reduxPromise from 'redux-promise';

import * as serviceWorker from './serviceWorker';

import articlesReducer from "./reducers/articlesReducer";
import selectedCategoryReducer from "./reducers/selectedCategoryReducer";
import currentUserReducer from "./reducers/currentUserReducer";
import toastReducer from "./reducers/toastReducer";
import sessionReducer from "./reducers/sessionReducer";

import App from "./App";
import {LOADING} from "./constants/constants";
import drawerReducer from "./reducers/drawerReducer";

let middleware
if (process.env.NODE_ENV == 'production') {
    middleware = applyMiddleware(reduxPromise)
} else {
    middleware = applyMiddleware(reduxPromise, logger)
}

const reducers = combineReducers({
    articles: articlesReducer,
    selectedCategory: selectedCategoryReducer,
    currentUser: currentUserReducer,
    session: sessionReducer,
    toast: toastReducer,
    drawer: drawerReducer,
});

const initialState = {
    articles: [],
    selectedCategory: 'all',
    currentUser: '',
    session: LOADING,
    toast: '',
    // drawer: window.innerWidth < 768 ? MOBILE_DRAWER_WIDTH : DESKTOP_DRAWER_WIDTH
    drawer: 0,
};

ReactDOM.render(
    <Provider store={createStore(reducers, initialState, middleware)}>
        <App />
    </Provider>,
    document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
