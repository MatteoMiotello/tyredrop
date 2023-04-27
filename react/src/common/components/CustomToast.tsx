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
    toast: ToastConfig | null
}
export const CustomToast: React.FC<CustomToastProps> = (props) => {
    return <Transition
        show={props.toast ? true : false}
    >
        <Toast>
            {
                props.toast ? <Alert
                    key={props.toast.key}
                    type={props.toast.type}
                >
                    {props.toast.message}
                </Alert> : <></>
            }
        </Toast>
    </Transition>;
};