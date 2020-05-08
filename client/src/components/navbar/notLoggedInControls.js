import React from 'react';
import { Link } from 'react-router-dom';

import {withStyles} from '@material-ui/core/styles';
import Button from "@material-ui/core/Button";

const NotLoggedInControls = () => {
    const LoginButton = withStyles((theme) => ({
        root: {
            color: 'white',
            marginLeft: theme.spacing(2),
        },
    }))(Button);

    return (
        <div>
            <Button size="small" variant="contained" color="secondary" component={Link} to={'/signup'}>
                Sign Up
            </Button>
            <LoginButton size="small" color="secondary" component={Link} to={'/signin'}>Sign in</LoginButton>
        </div>
    );
}

export default  NotLoggedInControls