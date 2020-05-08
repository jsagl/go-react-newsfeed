import {SET_CURRENT_USER} from "../constants/constants";

const currentUserReducer = (state, action) => {
    if (state === undefined) {
        return {};
    }

    switch (action.type) {
        case SET_CURRENT_USER:
            return action.payload;
        default:
            return state;
    }
};

export default currentUserReducer;