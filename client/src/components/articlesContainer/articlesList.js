import React, {useEffect, useState} from "react";
import {useSelector} from "react-redux";

import Article from "./article";

import {makeStyles} from "@material-ui/core/styles";
import CircularProgress from '@material-ui/core/CircularProgress';
import Fab from "@material-ui/core/Fab";
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp';

const ArticlesList = (props) => {
    const useStyles = makeStyles((theme) => ({
        articlesContainer: {
            scrollbarWidth: 'thin',
            maxWidth: '1000px',
            margin: '0 auto',
            marginTop: '96px',
            transition: 'all 0.2s linear',
            paddingLeft: drawerWidth,
        },
        spacing: {
            width: '100%',
            marginLeft: "0px",
            marginRight: "0px"
        },
        bottomIcon: {
            margin: '0 auto',
            marginTop: '16px',
            marginBottom: '16px',
            alignItems: 'center',
            justifyContent: 'center',
        },
        visible: {
            display: 'block',
        },
        hidden: {
            display: 'none',
        }
    }));
    const drawerWidth = useSelector(state => state.drawer)
    const classes = useStyles();

    const category = useSelector(state => state.selectedCategory);
    const articles = props.articles.filter(article => category === 'all' ? true : article.category === category);
    const numOfArticles = articles.length;

    const initialArticlesLimit = 30;
    const [loaderDisplay, setLoaderDisplay] = useState(false)
    const [limitDisplay, setLimitDisplay] = useState(false)

    const [articlesLimit, setArticlesLimit] = useState(initialArticlesLimit);

    const location = window.location.pathname

    useEffect(()=> {
        const handleScroll = () => {
            const bottomReached = window.innerHeight + document.documentElement.scrollTop + 5 >= document.documentElement.offsetHeight;
            if (bottomReached && articlesLimit < numOfArticles) {
                setLoaderDisplay(true)
                setTimeout(() => {
                    setLoaderDisplay(false)
                    setArticlesLimit(articlesLimit + 10)
                }, 300);
            }
            if (bottomReached && articlesLimit >= numOfArticles) {
                setLimitDisplay(true)
            }

            if (document.documentElement.scrollTop <= 100) {
                setLimitDisplay(false)
            }
        }

        window.addEventListener('scroll', handleScroll);

        return () => window.removeEventListener('scroll', handleScroll);
    }, [category, location, articles, articlesLimit, numOfArticles])

    const scrollBackTop = () => {
        window.scrollTo({
            top: 0,
            behavior: 'smooth',
        });
    }

    return(
        <div className={classes.articlesContainer}>
            {
                articles.slice(0, articlesLimit).map((article, index) => <Article key={`${article.date}&${index}&${article.bookmarked}`} article={article} type={props.resourcesType} />)
            }
            <CircularProgress
                className={`${loaderDisplay ? classes.visible : classes.hidden} ${classes.bottomIcon}` }
            />
            <Fab
                color="secondary" size="small" aria-label="scroll back to top"
                onClick={scrollBackTop}
                className={`${limitDisplay ? classes.visible : classes.hidden} ${classes.bottomIcon}`}
            >
                <KeyboardArrowUpIcon />
            </Fab>
        </div>

    )
}

export default ArticlesList;