import axios from 'axios';
import {
    AUTHENTICATED, CHECK_SESSION, DENIED, EMPTY_ARTICLES,
    FETCH_ARTICLES,
    FETCH_BOOKMARKS,
    HIDE_TOAST,
    REFRESH_TOKEN, REMOVE_RESOURCE,
    SELECT_CATEGORY,
    SET_CURRENT_USER, SET_DRAWER_WIDTH,
    SET_SESSION,
    SHOW_TOAST, SIGN_OUT
} from "../constants/constants";

const axiosInstance = axios.create({
    baseURL: 'api/v1',
    headers: {
        "Accept": "application/json",
        "Content-Type": "application/json",
    },
})

const articlesError = {
    type: FETCH_ARTICLES,
    payload: [{
        title: "An unexpected error occurred, please try to reload the page",
        category: 'golang',
        target_url: '',
        source_name: 'Go back-end'
    }]
}

const checkSession = () => {
    return axiosInstance({
        url: `/check_session`,
        withCredentials: true,
        cookies: ['sessionToken', 'rememberMeToken'],
        method: 'get',
    }).then(response => {
        return {type: CHECK_SESSION, payload: AUTHENTICATED }
    }).catch(response => {
        return {type: CHECK_SESSION, payload: DENIED }
    })
}

const refreshJwtToken = () => {
    return axiosInstance({
        url: `/refresh`,
        withCredentials: true,
        cookie: 'sessionToken',
        method: 'get',
    }).then(response => {
        // setAuthenticationCookie(response.data)
        return {type: REFRESH_TOKEN, payload: AUTHENTICATED }
    }).catch(response => {
        return {type: REFRESH_TOKEN, payload: DENIED }
    })
}

const signOut = () => {
    return {type: SIGN_OUT, payload: DENIED }
}

const setSession = (status) => {
    return {type: SET_SESSION, payload: status }
}

const fetchArticles = () => {
    return axiosInstance({
        url: `/articles`,
        withCredentials: true,
        cookie: 'sessionToken',
        method: 'get',
    }).then(response => {
       return {type: FETCH_ARTICLES, payload: response.data }
    }).catch(response => articlesError)
};

const removeResourceFromList = (resource) => {
    return {type: REMOVE_RESOURCE, payload: resource}
}

const emptyArticles = () => {
    return {type: EMPTY_ARTICLES, payload: [] }
};

const fetchBookmarks = () => {
    return axiosInstance({
        url: `/favorites`,
        withCredentials: true,
        cookie: 'sessionToken',
        method: 'get',
    }).then(response => {
        return {type: FETCH_BOOKMARKS, payload: response.data }
    }).catch(response => articlesError)
};

const selectCategory = (category) => {
    return {
        type: SELECT_CATEGORY,
        payload: category
    }
}

const setCurrentUser = (user) => {
    return {
        type: SET_CURRENT_USER,
        payload: user
    }
}

const showToast = (message) => {
    return {
        type: SHOW_TOAST,
        payload: message
    }
}

const hideToast = () => {
    return {
        type: HIDE_TOAST,
        payload: ''
    }
}

const setDrawerWidth = (width) => {
    return {
        type: SET_DRAWER_WIDTH,
        payload: width,
    }
}

export {
    refreshJwtToken, REFRESH_TOKEN,
    fetchArticles, FETCH_ARTICLES,
    fetchBookmarks, FETCH_BOOKMARKS,
    selectCategory, SELECT_CATEGORY,
    setCurrentUser, SET_CURRENT_USER,
    showToast, SHOW_TOAST,
    hideToast, HIDE_TOAST,
    setSession, SET_SESSION,
    emptyArticles,
    setDrawerWidth, SET_DRAWER_WIDTH,
    removeResourceFromList, REMOVE_RESOURCE,
    signOut, SIGN_OUT,
    checkSession, CHECK_SESSION
}
