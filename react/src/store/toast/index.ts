import { ToastState as State} from "./toast-state";
export type ToastState = State;

import toastSlice from './toast-slice';
export { toastMiddleware, removeToast } from "./toast-slice";
export default toastSlice;

import * as actions  from './toast-slice';
export const toastActions = {
	success: actions.setSuccess,
	error: actions.setError,
	info: actions.setInfo,
	warning: actions.setWarning,
};

import {selectToasts} from './toast-selector';
export {useToast} from "./useToast";
export {selectToasts};