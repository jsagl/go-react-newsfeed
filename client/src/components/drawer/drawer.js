import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import BookmarksIcon from '@material-ui/icons/Bookmarks';
import RssFeedIcon from '@material-ui/icons/RssFeed';
import Drawer from "@material-ui/core/Drawer";
import {useSelector} from "react-redux";
import DrawerLink from "./drawer_link";


export default function NavDrawer() {
    const useStyles = makeStyles((theme) => ({
        root: {
            display: 'flex',
        },
        drawer: {
            zIndex: 1099,
            width: drawerWidth,
            transition: 'all 0.2s linear',
            flexShrink: 0,
        },
        drawerPaper: {
            width: drawerWidth,
            zIndex: 1099,
            transition: 'all 0.2s linear',
            backgroundColor: 'white',
        },
        drawerContainer: {
            overflowY: 'auto',
            overflowX: 'hidden',
            paddingTop: '48px',
        },
        listItem: {
            color: theme.palette.primary.dark,
        },
        clickedListItem: {
            color: theme.palette.primary.dark,
            fontWeight: "bold",
        },
        listText: {
            paddingLeft: '24px',
        }
    }));

    const drawerWidth = useSelector(state => state.drawer)
    const classes = useStyles();

    return (
        <Drawer
            className={classes.drawer}
            variant="permanent"
            classes={{
                paper: classes.drawerPaper,
            }}
        >
            <div className={classes.drawerContainer}>
                <List>
                    <DrawerLink text='Newsfeed' route='/' icon={<RssFeedIcon/>} />
                    <DrawerLink text='Bookmarks' route='/bookmarks' icon={<BookmarksIcon/>} />
                </List>
            </div>
        </Drawer>
    );
}