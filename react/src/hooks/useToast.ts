import {useContext} from "react";
import {ToastContext} from "../Root";

export const useToast = () => {
    const { toasts, setToasts: setToast } = useContext<any>( ToastContext );

    const setSuccess = (message: string ) => {
        setToast( { type: 'success', message: message } );

        removeItem();
    };

    const setError = (message: string ) => {
        setToast( { type: 'error', message: message } );

        removeItem();
    };

    const setInfo = (message: string ) => {
        setToast( { type: 'info', message: message } );

        removeItem();
    };

    const removeItem = ( ) => {
        setTimeout( (  ) => {
            setToast( null );
        }, 3000 );
    };

    return { addSuccess: setSuccess, addError: setError };
};