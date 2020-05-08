import {HIDE_TOAST, SHOW_TOAST} from "../constants/constants";

const toastReducer = (state, action) => {
    if (state === undefined) {
        return {};
    }

    switch (action.type) {
        case SHOW_TOAST:
            return action.payload;
        case HIDE_TOAST:
            return action.payload;
        default:
            return state;
    }
};

export default toastReducer;