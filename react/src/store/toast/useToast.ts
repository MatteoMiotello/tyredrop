import { useDispatch } from "react-redux";
import { toastActions } from ".";
export const useToast= () => {
	const dispatch = useDispatch();

	const success = (message: string) => {
		dispatch( toastActions.success( message ) );
	};

	const error = (message: string) => {
		dispatch( toastActions.error( message ) );
	};

	const warning = (message: string) => {
		dispatch( toastActions.warning( message ) );
	};

	const info = (message: string) => {
		dispatch( toastActions.info( message ) );
	};

	return { success, error, warning, info };
};