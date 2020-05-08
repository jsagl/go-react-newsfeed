import {REFRESH_TOKEN, SET_SESSION, SIGN_OUT} from "../constants/constants";

const sessionReducer = (state, action) => {
    if (state === undefined) {
        return {};
    }

    switch (action.type) {
        case REFRESH_TOKEN:
        case SET_SESSION:
        case SIGN_OUT:
            return action.payload;
        default:
            return state;
    }
};

export default sessionReducer;