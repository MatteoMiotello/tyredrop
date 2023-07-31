import { createSelector } from "@reduxjs/toolkit";
import {Store} from "../store";

const toastState = ( state: Store ) => state.toast;

export const selectToasts = createSelector( [toastState], ( toastState ) => {
	return toastState.toasts;
} );