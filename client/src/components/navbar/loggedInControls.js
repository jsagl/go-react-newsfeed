import React, {useEffect, useState} from 'react';
import { useHistory } from 'react-router-dom';

import IconButton from '@material-ui/core/IconButton';
import Menu from '@material-ui/core/Menu';
import MenuItem from "@material-ui/core/MenuItem";

import {AccountCircle} from "@material-ui/icons";
import {useDispatch} from "react-redux";
import {fetchArticles, refreshJwtToken, signOut} from "../../actions";

const LoggedInControls = () => {
    const dispatch = useDispatch();

    useEffect(() => {
        const interval = setInterval(
            () => {
                dispatch(refreshJwtToken())
            },
            270 * 1000
        );

        return () => clearInterval(interval);
    }, [dispatch])

    const [anchorEl, setAnchorEl] = useState(null);
    const history = useHistory();

    const handleClick = (event) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    const signout = () => {
        dispatch(signOut());
        setAnchorEl(null);
        history.push('/');
        dispatch(fetchArticles());
    }


    return (
        <div>
            <IconButton
                edge="end"
                aria-label="account"
                aria-controls="session-menu"
                aria-haspopup="true"
                color="inherit"
                onClick={handleClick}
            >
                <AccountCircle />
            </IconButton>
            <Menu
                id="session-menu"
                anchorEl={anchorEl}
                keepMounted
                open={Boolean(anchorEl)}
                onClose={handleClose}
            >
                <MenuItem onClick={signout}>Sign out</MenuItem>
            </Menu>
        </div>
    );
}

export default  LoggedInControls