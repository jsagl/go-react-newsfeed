import React, {useEffect, useRef, useState} from "react";
import {useSelector} from "react-redux";

import Article from "./article";

import {makeStyles} from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import CircularProgress from '@material-ui/core/CircularProgress';
import Fab from "@material-ui/core/Fab";
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp';

const ArticlesList = (props) => {
    const useStyles = makeStyles((theme) => ({
        articlesContainer: {
            margin: '95px 5px 5px 5px',
            padding: '5px',
            height: 'calc(100vh - 110px)',
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'flex-start',
            alignItems: 'center',
            overflowY: 'auto',
            scrollbarWidth: 'thin',
        },
        spacing: {
            width: '100%',
            marginLeft: "0px",
            marginRight: "0px"
        },
        bottomIcon: {
            marginTop: '16px',
            marginBottom: '16px',
            width: '100%',
            alignItems: 'center',
            justifyContent: 'center',
        },
        visible: {
            display: 'flex',
        },
        hidden: {
            display: 'none',
        }
    }));
    const classes = useStyles();

    const topOfListRef = useRef(null);
    const scrollRef = useRef(null);

    const category = useSelector(state => state.selectedCategory);
    const articles = props.articles.filter(article => category === 'all' ? true : article.category === category);

    const initialArticlesLimit = 30;

    const [loaderDisplay, setLoaderDisplay] = useState(false)
    const [limitDisplay, setLimitDisplay] = useState(false)
    const [articlesLimit, setArticlesLimit] = useState(initialArticlesLimit);

    const location = window.location.pathname

    useEffect(()=> {
        scrollBackTop();
        setLimitDisplay(false);
        setArticlesLimit(initialArticlesLimit);
    }, [category, location])

    const handleScroll = (e) => {
        const bottomOfDiv = e.target.scrollHeight - e.target.scrollTop - 1 < e.target.clientHeight;
        const numOfArticles = articles.length

        if (bottomOfDiv && articlesLimit < numOfArticles) {
            setLoaderDisplay(true)
            setTimeout(() => {
                setLoaderDisplay(false)
                setArticlesLimit(articlesLimit + 10)
            }, 200);
        }

        if (bottomOfDiv && articlesLimit >= numOfArticles) {
            setLimitDisplay(true)
        }
    }

    const scrollBackTop = () => {
        topOfListRef.current.scrollIntoView({
            behavior: 'smooth',
            block: 'start',
        });
    }

    return(
        <div className={classes.articlesContainer} onScroll={handleScroll} ref={scrollRef}>
            <div ref={topOfListRef} id="articlesListTop"></div>
            <Grid
                container
                spacing={1}
                justify="center"
                classes={{container: classes.spacing}}
            >
                <Grid item xs={12} style={{width: '100%',}}/>
                {
                    articles.slice(0, articlesLimit).map((article, index) => <Article key={`${article.date}&${index}`} article={article} type={props.resourcesType} />)
                }
                <Grid item className={`${loaderDisplay ? classes.visible : classes.hidden} ${classes.bottomIcon}` }>
                    <CircularProgress />
                </Grid>
                <Grid item className={`${limitDisplay ? classes.visible : classes.hidden} ${classes.bottomIcon}`}>
                    <Fab color="secondary" size="small" aria-label="scroll back to top" onClick={scrollBackTop}>
                        <KeyboardArrowUpIcon />
                    </Fab>
                </Grid>
            </Grid>
        </div>

    )
}

export default ArticlesList;