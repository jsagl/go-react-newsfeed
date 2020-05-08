import React from 'react';
import { useHistory } from 'react-router-dom';

import LoginControl from "./loginControl";

import {makeStyles, withStyles} from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';

import MenuIcon from '@material-ui/icons/Menu';
import MenuOpenIcon from '@material-ui/icons/MenuOpen';

import Button from "@material-ui/core/Button";
import {useDispatch, useSelector} from "react-redux";
import {emptyArticles, fetchArticles, setDrawerWidth} from "../../actions";
import {DESKTOP_DRAWER_WIDTH, MOBILE_DRAWER_WIDTH} from "../../constants/constants";

const NavBar = () => {
    const Home = withStyles((theme) => ({
        root: {
            color: 'white',
            textTransform: 'none',
            fontSize: '1em',
        },
    }))(Button);

    const useStyles = makeStyles((theme) => ({
        menuButton: {
            marginRight: theme.spacing(2),
        },
    }));
    const classes = useStyles();
    const dispatch = useDispatch();
    const history = useHistory();
    const regex = /\/react-newsfeed/
    const location = window.location.pathname.replace(regex, '')

    const handleClickHome = () => {
        if (location !== '/') {
            dispatch(emptyArticles());
            dispatch(fetchArticles());
            history.push('/');
        }
    }

    const drawerWidth = useSelector(state => state.drawer)

    const handleClickMenu = () => {
        if (drawerWidth === 0) {
            dispatch(setDrawerWidth(window.innerWidth < 768 ? MOBILE_DRAWER_WIDTH : DESKTOP_DRAWER_WIDTH));
        } else {
            dispatch(setDrawerWidth(0))
        }
    }

    const icon = drawerWidth === 0 ? <MenuIcon /> : <MenuOpenIcon/>

    return (
        <AppBar position="fixed" square elevation={1}>
            <Toolbar variant="dense">
                <IconButton edge="start" onClick={handleClickMenu} className={classes.menuButton} color="inherit" aria-label="menu">
                    {icon}
                </IconButton>
                <Home size="small" onClick={handleClickHome}>Home</Home>
                <div style={{flexGrow: 1}}></div>
                <LoginControl/>
            </Toolbar>
        </AppBar>
    );
}

export default  NavBar