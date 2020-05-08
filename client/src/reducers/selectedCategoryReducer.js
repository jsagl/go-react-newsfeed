import {SELECT_CATEGORY} from "../constants/constants";

const selectedCategoryReducer = (state, action) => {
    if (state === undefined) {
        return {};
    }

    switch (action.type) {
        case SELECT_CATEGORY:
            return action.payload;
        default:
            return state;
    }
};

export default selectedCategoryReducer;