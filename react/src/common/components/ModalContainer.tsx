import React, {ReactNode} from "react";

export type ModalData = {
    content: ReactNode
    id: string
}

type ModalContainerProps = {
    modal: ModalData | null
}

const ModalContainer: React.FC<ModalContainerProps> = ( props ) => {
    return <>
        { props.modal && props.modal.content}
    </>;
};

export default ModalContainer;