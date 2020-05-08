import React, {useRef} from "react";
import { Formik } from 'formik';
import * as Yup from 'yup';
import { Link as RouterLink, useHistory } from 'react-router-dom';
import {useDispatch, useSelector} from "react-redux";
import axios from 'axios';

import {makeStyles} from "@material-ui/core/styles";
import Grid from "@material-ui/core/Grid";
import TextField from "@material-ui/core/TextField";
import Container from "@material-ui/core/Container";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Checkbox from "@material-ui/core/Checkbox";
import {Link} from "@material-ui/core";
import {
    emptyArticles,
    fetchArticles,
    hideToast, selectCategory,
    setCurrentUser,
    setSession,
    showToast
} from "../../actions";
import Snackbar from "@material-ui/core/Snackbar";
import {AUTHENTICATED} from "../../constants/constants";

axios.defaults.withCredentials = true

const SignInForm = () => {
    const useStyles = makeStyles((theme) => ({
        root: {
            padding: '40px',
            paddingTop: '20px',
            '& h5': {
                margin: theme.spacing(1),
            },
            '& .MuiTextField-root': {
                margin: theme.spacing(1),
            },
            '& .MuiFormControlLabel-root': {
                margin: theme.spacing(1),
            },
        },
        formContainer: {
            height: 'calc(100vh - 110px)',
            [theme.breakpoints.down('md')]: {
                marginTop: "70px",
            },
            [theme.breakpoints.up('md')]: {
                marginTop: "80px",
            },
        },
        submitButton: {
            width: '100%',
            margin: theme.spacing(1),
        },
        success: {
            padding: '40px',
            paddingTop: '20px',
            '& h5': {
                margin: theme.spacing(1),
            },
            '& button': {
                margin: theme.spacing(1),
            }
        }
    }));
    const classes = useStyles();

    const toast = useSelector(state => state.toast)

    const emailFieldRef = useRef(null);

    const dispatch = useDispatch();
    let history = useHistory();

    const handleClose = () => {
        dispatch(hideToast());
    };

    const handleSuccessfulSignIn = (response) => {
        const user = response.data['username']
        history.push('/');
        dispatch(selectCategory('all'));
        dispatch(emptyArticles());
        dispatch(fetchArticles());
        dispatch(setCurrentUser(user));
        dispatch(setSession(AUTHENTICATED));
        dispatch(showToast(`Welcome back ${user}!`));
    }

    const handleFailedSignIn = () => {
        dispatch(showToast("Invalid username or password."));
        emailFieldRef.current.focus();
    }

    return(
        <Container maxWidth="sm" className={classes.formContainer} >
            <Grid item xs={12}>
                <Formik
                    initialValues={{ Email: '', Password: '', RememberMe: false }}
                    onSubmit={(values, { setSubmitting, resetForm }) => {
                        setSubmitting(true);

                        axios.post('/api/v1/signin', values)
                        .then(response => handleSuccessfulSignIn(response) )
                        .catch(_ => {
                            resetForm();
                            handleFailedSignIn();
                        })
                    }}
                    validationSchema={
                        Yup.object().shape({
                            Email: Yup.string().email().required('Required'),
                            Password: Yup.string().required('Required'),
                        })}
                >
                    {({
                          values,
                          touched,
                          errors,
                          isSubmitting,
                          handleChange,
                          // handleBlur,
                          handleSubmit,
                      }) => (
                        <form className={classes.root} onSubmit={handleSubmit}>
                            <Typography variant="h5">
                                Sign in
                            </Typography>
                            <div>
                                <TextField variant={'outlined'} fullWidth autoFocus required label="Email"
                                           inputRef={emailFieldRef}
                                           name='Email'
                                           value={values.Email}
                                           onChange={handleChange}
                                            // onBlur={handleBlur}
                                           helperText={(errors.Email && touched.Email) && errors.Email}
                                />
                            </div>
                            <div>
                                <TextField
                                    fullWidth variant={'outlined'} required label="Password" type="password"
                                    name='Password'
                                    value={values.Password}
                                    onChange={handleChange}
                                    // onBlur={handleBlur}
                                    helperText={(errors.Password && touched.Password) && errors.Password}
                                />
                            </div>
                            <div>
                                <FormControlLabel
                                    name='RememberMe'
                                    control={<Checkbox color="primary" />}
                                    value={values.RememberMe}
                                    onChange={handleChange}
                                    label="Remember me"
                                />
                            </div>
                            <Button type="submit" className={classes.submitButton}
                                    variant="contained" color="secondary" size="large"
                                    disabled={isSubmitting}
                            >
                                Sign in
                            </Button>
                            <Grid container justify="flex-end">
                                <Grid item>
                                    <Typography variant={'body1'}>
                                        You don't have an account yet? <Link component={RouterLink} to={'/signup'} >Sign up</Link>
                                    </Typography>
                                </Grid>
                            </Grid>
                        </form>
                    )}
                </Formik>
            </Grid>
            <Snackbar open={toast !== ''} autoHideDuration={5000} onClose={handleClose} message={toast}/>
        </Container>
    )
}

export default SignInForm;