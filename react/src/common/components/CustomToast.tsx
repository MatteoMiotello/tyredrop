import React from "react";
import Alert, {AlertType} from "../components-library/Alert";
import {Toast} from "../components-library/Toast";
import {Transition} from "@headlessui/react";

export type ToastConfig = {
    key: string | number
    type: AlertType
    message: string
}

type CustomToastProps = {
    toasts: ToastConfig[]
}
export const CustomToast: React.FC<CustomToastProps> = (props) => {
    return <Toast>
        {
            props.toasts.map((toast: ToastConfig, key: number) => {
                return (
                    <Transition
                        key={key}
                        show={!!toast}
                    >
                        <Alert
                            key={toast.key}
                            type={toast.type}
                        >
                            {toast.message}
                        </Alert>
                    </Transition>
                );
            })
        }
    </Toast>;
};