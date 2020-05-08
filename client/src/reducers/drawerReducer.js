import {SET_DRAWER_WIDTH} from "../constants/constants";

const drawerReducer = (state, action) => {
    if (state === undefined) {
        return {};
    }

    switch (action.type) {
        case SET_DRAWER_WIDTH:
            return action.payload;
        default:
            return state;
    }
};

export default drawerReducer;