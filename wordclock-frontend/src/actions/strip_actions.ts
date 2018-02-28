import {createAction} from "typesafe-actions";

export const SET_BRIGHTNESS = 'SET_BRIGHTNESS';

export const stripActions = {
    setBrightness: createAction(SET_BRIGHTNESS, (b: number) => ({
        type: SET_BRIGHTNESS,
        payload: b
    }))
};