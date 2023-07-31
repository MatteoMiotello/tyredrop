import React from "react";
import { selectToasts } from "../../store/toast";
import { useSelector } from "react-redux";
import { Alert } from "./shelly-ui";
import {ToastConfiguration} from "../../store/toast/toast-state";

const ToastContainer: React.FC = () => {
    const toasts = useSelector( selectToasts );

    return <div className="toast toast-top toast-end z-50">
        {
            toasts.map( (toastConfig: ToastConfiguration) => {
                return <Alert type={toastConfig.type} key={toastConfig.id} showCloseButton={true}>
                    {toastConfig.message}
                </Alert>;
            } )
        }
    </div>;
};

export default ToastContainer;