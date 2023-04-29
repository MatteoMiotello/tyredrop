import {useContext, useEffect} from "react";
import {ToastConfig} from "../common/components/CustomToast";
import ToastContext from "../common/contexts/toast-context";

const generateUnique = (): number => {
    return new Date().getTime();
};

export const useToast = () => {
    const {toasts, setToasts: setToast} = useContext(ToastContext);

    useEffect(() => {
        const i = setInterval(() => {

        }, 3000);

        return () => {
            clearInterval(i);
        };
    }, []);

    const addToast = (newToast: ToastConfig) => {
        setToast([...toasts, newToast]);
    };

    const setSuccess = (message: string) => {
        addToast({key: generateUnique(), type: 'success', message: message});

        removeItem();
    };

    const setError = (message: string) => {
        addToast({key: generateUnique(), type: 'error', message: message});

        removeItem();
    };

    const setInfo = (message: string) => {
        addToast({key: generateUnique(), type: 'info', message: message});

        removeItem();
    };

    const removeItem = () => {
        setTimeout(() => {
            setToast(null);
        }, 3000);
    };

    return {setSuccess, setError};
};