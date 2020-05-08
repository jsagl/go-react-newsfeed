import React from 'react';
import {useSelector} from "react-redux";
import NotLoggedInControls from "./notLoggedInControls";
import LoggedInControls from "./loggedInControls";
import {AUTHENTICATED, LOADING} from "../../constants/constants";

const LoginControl = () => {
    const session = useSelector(state => state.session)

    if (session === LOADING) {
        return (
            <div/>
        );
    } else if (session === AUTHENTICATED) {
        return (
            <LoggedInControls />
        );
    } else {
        return (
            <NotLoggedInControls />
        );
    }

}

export default  LoginControl