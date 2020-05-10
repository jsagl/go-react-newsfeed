import React, {useEffect, useState} from 'react';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import ArticlesAndBookmarksPages from "./components/articlesContainer/articlesAndBookmarksPages";

import NavBar from "./components/navbar/navbar";
import SignUpForm from "./components/signup/signUpForm";
import SignInForm from "./components/signin/signinForm";
import {useDispatch, useSelector} from "react-redux";
import {checkSession, fetchArticles, fetchBookmarks} from "./actions";
import NavDrawer from "./components/drawer/drawer";
import {ARTICLE, BOOKMARK} from "./constants/constants";

const App = () => {
    const dispatch = useDispatch();
    const [shouldCheckSession, setShouldCheckSession] = useState(true);

    useEffect(() => {
        if (shouldCheckSession) {
            dispatch(checkSession());
            setShouldCheckSession(false);
        }
    }, [dispatch, shouldCheckSession])

    const resource = useSelector(state => state.articles)

    return (
        <Router basename={process.env.PUBLIC_URL}>
            <NavBar />
            <NavDrawer/>
            <div style={{flexGrow: 1}}>
            <Switch>
                    <Route exact path="/">
                        <ArticlesAndBookmarksPages resource={resource} fetchResources={fetchArticles} resourcesType={ARTICLE} />
                    </Route>
                    <Route exact path="/bookmarks">
                        <ArticlesAndBookmarksPages resource={resource} fetchResources={fetchBookmarks} resourcesType={BOOKMARK} />
                    </Route>
                    <Route exact path="/signup">
                        <SignUpForm/>
                    </Route>
                    <Route exact path="/signin">
                        <SignInForm/>
                    </Route>
            </Switch>
            </div>
        </Router>
    );
}

export default App;
