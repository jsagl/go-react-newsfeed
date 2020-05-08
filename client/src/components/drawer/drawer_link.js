import React, {useEffect, useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import ListItem from '@material-ui/core/ListItem';
import {AUTHENTICATED} from "../../constants/constants";
import {useDispatch, useSelector} from "react-redux";
import { useHistory } from 'react-router-dom';
import {emptyArticles, fetchArticles, fetchBookmarks, selectCategory} from "../../actions";


export default function DrawerLink(props) {
    const useStyles = makeStyles((theme) => ({
        listItem: {
            color: theme.palette.primary.dark,
        },
        clickedListItem: {
            color: theme.palette.secondary.main,
            fontWeight: "bolder",
        },
        listText: {
            paddingLeft: '24px',
        }
    }));

    const history = useHistory()
    const dispatch = useDispatch()
    const classes = useStyles();
    const session = useSelector(state => state.session);
    const [active, setActive] = useState(false);
    const regex = /\/react-newsfeed/
    const location = window.location.pathname.replace(regex, '')
    const linkRoute = props.route

    useEffect(() => {
        if (linkRoute === location) {
            setActive(true)
        } else {
            setActive(false)
        }
    }, [location, linkRoute])

    const handleClick = (route) => {
        if (route !== '/' && session !== AUTHENTICATED) {
            history.push('/signin')
            return
        }

        dispatch(emptyArticles())

        if (route === '/' || route === '/bookmarks') {
            dispatch(selectCategory('all'))
            dispatch(fetchRessources(route));
        }

        setActive(true);
        return history.push(route)
    }

    const fetchRessources = (route) => {
        return route === '/bookmarks' ? fetchBookmarks() : fetchArticles()
    }

    return (
        <ListItem button className={active ? classes.clickedListItem : classes.listItem} key={props.text} onClick={() => handleClick(props.route)}>
            {props.icon}
            <Typography className={classes.listText}>{props.text}</Typography>
        </ListItem>
    );
}