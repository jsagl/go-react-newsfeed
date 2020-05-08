import {FETCH_ARTICLES, FETCH_BOOKMARKS} from "../actions";
import {EMPTY_ARTICLES, REMOVE_RESOURCE} from "../constants/constants";

const articlesReducer = (state, action) => {
    if (state === undefined) {
        return {};
    }

    const resourceList = [...state];

    switch (action.type) {
        case EMPTY_ARTICLES:
        case FETCH_BOOKMARKS:
        case FETCH_ARTICLES:
            return action.payload;

        case REMOVE_RESOURCE:
            resourceList.find((resource, index, resourceList) => {
                if (resource.target_url === action.payload.target_url) {
                    resourceList.splice(index,1);
                    return true
                }

                return false
            });

            return resourceList;
        default:
            return state;
    }
};

export default articlesReducer;