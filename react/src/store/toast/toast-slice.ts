
import { Middleware, PayloadAction, SliceCaseReducers, createSlice } from "@reduxjs/toolkit";
import { ToastState, errorToast, infoToast, successToast, warningToast } from "./toast-state";

const toastSlice = createSlice<ToastState, SliceCaseReducers<ToastState>>({
	name: 'toast',
	initialState: {
		toasts: []
	},
	reducers: {
		setSuccess: (state, action: PayloadAction<string>) => {
			const toast = successToast( action.payload );

			state.toasts = [
				...state.toasts,
				toast
			];
		},
		setError: (state, action: PayloadAction<string>) => {
			const toast = errorToast( action.payload );
			
			state.toasts = [
				...state.toasts,
				toast
			];
		},
		setInfo: (state, action: PayloadAction<string>) => {
			const toast = infoToast( action.payload );
			state.toasts = [
				...state.toasts,
				toast
			];
		},
		setWarning: (state, action: PayloadAction<string>) => {
			const toast = warningToast( action.payload );
			state.toasts = [
				...state.toasts,
				toast
			];
		},
		removeToast: ( state  ) => {
			const current = [...state.toasts];
			current.splice( 0, 1 );
			state.toasts = [
				...current
			];
		}
	}
});


export const toastMiddleware: Middleware = (store) => (next) => ( action ) => {
	const actions = [ setSuccess.type, setError.type, setInfo.type, setWarning.type ];
	if ( action && actions.includes( action.type ) ) {
		setTimeout( () => {
			store.dispatch( removeToast() );
		}, 3000 );
	}

	if ( action ) {
		return next( action );
	}
};  

export const { setSuccess, setError, setInfo, setWarning, removeToast } = toastSlice.actions;
export default toastSlice.reducer;