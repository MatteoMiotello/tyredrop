import { createSelector } from "@reduxjs/toolkit";
import { AppState } from "../../store/store";

const toastState = ( state: AppState ) => state.toast;

export const selectToasts = createSelector( [toastState], ( toastState ) => {
	return toastState.toasts;
} );