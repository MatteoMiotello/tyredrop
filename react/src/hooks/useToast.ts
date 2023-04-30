import { useContext, useEffect} from "react";
import {ToastConfig} from "../common/components/CustomToast";
import ToastContext from "../common/contexts/toast-context";

const generateUnique = (): number => {
    return new Date().getTime();
};

export const useToast = () => {
    const {toasts, setToasts} = useContext(ToastContext);
    const removeLastToast = () => {
        toasts.splice(0, 1);

        setToasts([...toasts]);
    };

    useEffect(() => {
        const i = setInterval(() => {
            removeLastToast();
        }, 3000);

        return () => {
            clearInterval(i);
        };
    });

    const addToast = (newToast: ToastConfig) => {
        setToasts([...toasts, newToast]);
    };

    const setSuccess = (message: string) => {
        addToast({key: generateUnique(), type: 'success', message: message});
    };

    const setError = (message: string) => {
        addToast({key: generateUnique(), type: 'error', message: message});
    };

    const addInfo = (message: string) => {
        addToast({key: generateUnique(), type: 'info', message: message});
    };

    return {setSuccess, setError, addInfo};
};