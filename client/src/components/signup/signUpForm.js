import React, {useRef, useState} from "react";
import { Formik } from 'formik';
import * as Yup from 'yup';
import { Link as RouterLink, useHistory } from 'react-router-dom';
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
import {hideToast, showToast} from "../../actions";
import {useDispatch, useSelector} from "react-redux";
import Snackbar from "@material-ui/core/Snackbar";

const SignUpForm = () => {
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
                '& .MuiTypography-body1': {
                    fontSize: '10px',
                    fontStyle: 'italic',
                    color: 'grey',
                }
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
    const [successfulSignUp, setSuccessfulSignUp] = useState(false)
    const history = useHistory()
    const dispatch = useDispatch()
    const usernameInputRef = useRef(null);

    const handleClose = () => dispatch(hideToast());

    return(
        <Container maxWidth="sm" className={classes.formContainer} >
            <Grid item xs={12} style={successfulSignUp ? {display: 'none'} : {display: 'block'}}>
                <Formik
                    initialValues={{ Username: '', Email: '', Password: '', Terms: false }}
                    onSubmit={(values, { setSubmitting }) => {
                        setSubmitting(true);
                        axios.create()({
                            baseURL: '',
                            url: `/signup`,
                            method: 'post',
                            data: values,
                        }).then(response => {
                            setSuccessfulSignUp(true)
                            dispatch(showToast("Welcome on board! Please sign in to enjoy all the website's features."))
                            history.push('/signin')
                        }).catch(response => {
                            // Todo: better handle error types.
                            dispatch(showToast('An unexpected error occurred. The server may be down or a user with similar username/password may already exist. Please try again.'));
                            setSubmitting(false);
                            usernameInputRef.current.focus();
                        })
                    }}
                    validationSchema={
                        Yup.object().shape({
                        Username: Yup.string().required('Required').min(3, 'Username must be at least 3 characters').max(50, 'Username cannot be longer than 50 characters'),
                        Email: Yup.string().email().required('Required'),
                        Password: Yup.string().required('Required').min(8, 'Password must be at least 8 characters').max(45, 'Password cannot be longer than 45 characters'),
                        Terms: Yup.boolean().oneOf([true], 'Please accept terms and conditions'),
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
                                    Create your account
                                </Typography>
                                <div>
                                    <TextField
                                        variant={'outlined'} fullWidth required label="Username" autoFocus
                                        inputRef={usernameInputRef}
                                        name='Username'
                                        value={values.Username}
                                        onChange={handleChange}
                                        // onBlur={handleBlur}
                                        helperText={(errors.Username && touched.Username) && errors.Username}
                                    />
                                </div>
                                <div>
                                    <TextField variant={'outlined'} fullWidth required label="Email"
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
                                        name='Terms'
                                        control={<Checkbox color="primary" />}
                                        value={values.Terms}
                                        onChange={handleChange}
                                        label="I understand that this is a non-commercial side project created for training purpose only,
                                        that the author can close the website at any given time without prior notice and that he cannot
                                        be held responsible for any data breach"
                                    />
                                    <div className={"MuiFormHelperText-contained MuiFormHelperText-root"}>{(errors.Terms && touched.Terms) && errors.Terms}</div>
                                </div>
                                <Button type="submit" className={classes.submitButton}
                                        variant="contained" color="secondary" size="large"
                                        disabled={isSubmitting}
                                >
                                    Sign up
                                </Button>
                                <Grid container justify="flex-end">
                                    <Grid item>
                                        <Typography variant={'body1'}>
                                            Already have an account? <Link component={RouterLink} to={'/signin'} >Sign In</Link>
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

export default SignUpForm;