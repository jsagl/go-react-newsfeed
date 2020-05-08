import React, {useEffect, useState} from "react";
import { useHistory } from 'react-router-dom';
import {useDispatch, useSelector} from "react-redux";

import ArticlesList from "./articlesList";
import CategoriesList from "./categoriesList";
import Snackbar from "@material-ui/core/Snackbar";

import {hideToast} from "../../actions";
import {DENIED} from "../../constants/constants";

const ArticlesAndBookmarksPages = (props) => {
    const toast = useSelector(state => state.toast);
    const dispatch = useDispatch()
    
    const articles = useSelector(state => state.articles);
    const [shouldFetch, setShouldFetch] = useState(true);

    const session = useSelector(state => state.session)
    const history = useHistory()

    const regex = /\/react-newsfeed/
    const location = window.location.pathname.replace(regex, '')

    if (location === '/bookmarks' && session === DENIED) {
        history.push('/signin');
    }

    const fetchResources = props.fetchResources

    useEffect(()=> {
        if (shouldFetch) {
            dispatch(fetchResources());
            setShouldFetch(false)
        }

        const interval = setInterval(
            () => {
                dispatch(fetchResources())
            },
            15 * 60 * 1000
        );

        return () => clearInterval(interval);
    }, [dispatch, shouldFetch, location, fetchResources])
    
    const handleClose = () => {
        dispatch(hideToast());
    };

    return (
        <div>
            <CategoriesList articles={articles} />
            <ArticlesList articles={articles} resourcesType={props.resourcesType}/>
            <Snackbar open={toast !== ''} autoHideDuration={5000} onClose={handleClose} message={toast}/>
        </div>
    )
}

export default ArticlesAndBookmarksPages;