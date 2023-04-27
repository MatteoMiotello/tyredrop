import React, {PropsWithChildren} from "react";

type ToastProps = PropsWithChildren
export const Toast: React.FC<ToastProps> = ( props: ToastProps ) => {
    return <div className="toast toast-top toast-end">
            <div>
                {props.children}
            </div>
    </div>;
};