import React, {useState} from 'react';
import { useHistory } from 'react-router-dom';
import {useDispatch, useSelector} from "react-redux";
import moment from "moment/moment";
import axios from 'axios';

import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import Avatar from '@material-ui/core/Avatar';
import IconButton from '@material-ui/core/IconButton';
import Grid from "@material-ui/core/Grid";
import BookmarkIcon from '@material-ui/icons/Bookmark';
import BookmarkBorderIcon from '@material-ui/icons/BookmarkBorder';
import DeleteIcon from '@material-ui/icons/Delete';
import {Link} from "@material-ui/core";

import {ARTICLE, AUTHENTICATED} from "../../constants/constants";
import {removeResourceFromList, showToast} from "../../actions";


const Article = (props) => {
    const useStyles = makeStyles((theme) => ({
        item: {
            marginRight: '12px',
            marginLeft: '12px',
            width: '100%',
            maxWidth: '1000px',
        },
        avatarImg:{
            objectFit: 'contain',
        },
        link: {
            textDecoration: 'none',
            '&:hover': {
                color: theme.palette.primary.main,
                textDecoration: 'none',
            }
        },
    }));

    const classes = useStyles();
    const date = moment(props.article.date).format('DD/MM');
    // const time = moment(props.article.date).format('HH:mm')

    const session = useSelector(state => state.session);
    const history = useHistory();
    const dispatch = useDispatch();
    const [bookmarked, setBookmarked] = useState(props.article.bookmarked);

    const createBookmark = (article) => {
        return axios.post('/api/v1/favorites', article, {withCredentials: true, cookie: 'sessionToken'})
    }

    const deleteBookmark = (article) => {
        return axios.delete(`/api/v1/favorites`, {data: {target_url: article.target_url}, withCredentials: true, cookie: 'sessionToken'})
    }

    const handleBookmarkClick = () => {
        if (session !== AUTHENTICATED) {
            history.push('/signin')
        } else if (bookmarked) {
            deleteBookmark(props.article)
                .then(response => {
                    setBookmarked(false)
                }).catch(response => {
                    dispatch(showToast('An error occurred while trying to remove this bookmark'))
                })
        } else {
            createBookmark(props.article)
                .then(response => {
                    setBookmarked(true)
                }).catch(response => {
                    dispatch(showToast('An error occurred while trying to bookmark this article.'))
                })
        }
    }

    const handleDeleteClick = () => {
        deleteBookmark(props.article)
            .then(response => {
                dispatch(removeResourceFromList(props.article))
            }).catch(response => {
            dispatch(showToast('An error occurred while trying to remove this bookmark'))
        })
    }


    const bookMarkIcon = bookmarked ? <BookmarkIcon/> : <BookmarkBorderIcon/>
    let action
    if (props.type === ARTICLE) {
        action = (
            <div>
                <IconButton onClick={handleBookmarkClick}>
                    {bookMarkIcon}
                </IconButton>
            </div>
        )
    } else {
        action = (
            <div>
                <IconButton onClick={handleDeleteClick}>
                    <DeleteIcon/>
                </IconButton>
            </div>
        )
    }

    return (
        <Grid item xs={12} classes={{item: classes.item}}>
            <Card elevation={1}>
                <CardHeader
                    avatar={
                        <Avatar
                            classes={{img: classes.avatarImg}}
                            variant="square"
                            alt={`${props.article.category}-logo`}
                            src={`./assets/${props.article.category}_logo.svg`}/>
                    }
                    action={action}
                    title={
                        <Link
                            className={classes.link}
                            href={`${props.article.target_url}`}
                            target="_blank"
                            rel="noreferrer"
                            color="inherit"
                        >
                            {props.article.title}
                        </Link>
                    }
                    subheader={`${props.article.source_name} - ${date}`}
                />
            </Card>
        </Grid>
    );
}

export default Article;
