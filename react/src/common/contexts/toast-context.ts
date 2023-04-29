import React from "react";
import {ToastConfig} from "../components/CustomToast";

type ToastContextType = {
    toasts: ToastConfig[]
    setToasts: ( toasts: ToastConfig[] ) => void
}

const ToastContext = React.createContext<ToastContextType>({
    toasts: [],
    setToasts: (toasts: ToastConfig[]): void => { return; }
});

export default ToastContext;