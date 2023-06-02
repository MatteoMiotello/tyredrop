import React, {PropsWithChildren} from "react";

type ToastProps = PropsWithChildren
export const Toast: React.FC<ToastProps> = ( props: ToastProps ) => {
    return <div className="toast toast-top toast-end z-50">
            <div>
                {props.children}
            </div>
    </div>;
};