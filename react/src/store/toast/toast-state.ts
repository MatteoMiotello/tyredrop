import { AlertType } from "../components/shelly-ui/Alert";

type ToastType = AlertType;

export type ToastConfiguration = {
	id: number | string
	type: ToastType
	message: string
}

const createToast = ( message: string, type: ToastType ): ToastConfiguration => {
	return {
		id: ( new Date() ).getTime(),
		type: type,
		message: message
	};
};

export const successToast = ( message: string ): ToastConfiguration =>  {
	return createToast( message, "success" );
};

export const errorToast = ( message: string ): ToastConfiguration =>  {
	return createToast( message, "error" );
};

export const warningToast = ( message: string ): ToastConfiguration =>  {
	return createToast( message, "warning" );
};

export const infoToast = ( message: string ): ToastConfiguration =>  {
	return createToast( message, "info" );
};

export type ToastState = {
    toasts: ToastConfiguration[]
}