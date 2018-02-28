import {RGBColor} from "../core/Color";
import {createAction} from "typesafe-actions";

export const COLOR_CHANGE = 'COLOR_CHANGE';

export const colorActions = {
    change: createAction(COLOR_CHANGE, (c: RGBColor) => ({
        type: COLOR_CHANGE,
        payload: c
    }))
};