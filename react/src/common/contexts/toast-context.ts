import React from "react";
import {ToastConfig} from "../components/ToastContainer";

type ToastContextData = {
    toasts: ToastConfig[]
    setToasts: ( toasts: ToastConfig[] ) => void
}

const ToastContext = React.createContext<ToastContextData>({
    toasts: [],
    setToasts: (toasts: ToastConfig[]): void => { return; }
});

export default ToastContext;