import React, {useEffect, useState} from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import {useDispatch, useSelector} from "react-redux";
import {selectCategory} from "../../actions";
import withStyles from "@material-ui/core/styles/withStyles";
import uniq from "lodash/uniq";

const CategoriesList = (props) => {
    const useStyles = makeStyles((theme) => ({
        root: {
            zIndex: '1098',
            paddingLeft: drawerWidth,
            transition: 'all 0.2s linear',
            top: '48px',
            backgroundColor: theme.palette.background.paper,
        },
    }));
    const drawerWidth = useSelector(state => state.drawer)
    const classes = useStyles();


    const StyledTabs = withStyles((theme) => ({
        root: {
            minHeight: "40px",
        },
        indicator: {
            backgroundColor: 'transparent',
        },
    }))((props) => <Tabs {...props} TabIndicatorProps={{ children: <div /> }} />);
    const StyledTab = withStyles((theme) => ({
        root: {
            margin: "auto",
            padding: "2px 12px",
            minHeight: "40px",
            '&:hover': {
                color: theme.palette.primary.main,
                opacity: 1,
            },
            '&$selected': {
                color: theme.palette.primary.main,
            },
            '&:focus': {
                color: theme.palette.primary.main,
            },
        },
        selected: {},
    }))((props) => <Tab disableRipple {...props} />);

    const categories = uniq(props.articles.map(article => article.category))
    if (categories.length !== 0) {
        categories.unshift('all')
    }

    const [value, setValue] = useState(0)
    const dispatch = useDispatch()
    const regex = /\/react-newsfeed/
    const location = window.location.pathname.replace(regex, '')

    useEffect(() => {
        setValue(0)
    }, [location])

    const handleChange = (event, newValue) => {
        setValue(newValue);
    };

    const determineValue = () => {
        if (value > categories.length - 1) {
            return 0
        }

        return value
    }

    return (
        <AppBar position="fixed" color="default" elevation={1} square className={classes.root}>
            <StyledTabs
                value={determineValue()}
                onChange={handleChange}
                variant="scrollable"
                scrollButtons="on"
                aria-label="scrollable auto tabs example"
            >
                {
                    categories.map((category, index) => {
                        return <StyledTab
                            key={`${category}${index}`}
                            label={category}
                            onClick={() => { dispatch(selectCategory(category)) }}
                        />
                    })
                }
            </StyledTabs>
        </AppBar>
    );
                            // {...tabProps(index)}
}

// function tabProps(index) {
//     return {
//         id: `scrollable-auto-tab-${index}`,
//         'aria-controls': `scrollable-auto-tabpanel-${index}`,
//     };
// }

export default CategoriesList;