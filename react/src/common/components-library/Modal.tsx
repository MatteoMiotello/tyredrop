import clsx from "clsx";
import React, {PropsWithChildren} from "react";
import {twMerge} from "tailwind-merge";

type ModalHeaderProps = PropsWithChildren;
const ModalHeader: React.FC<ModalHeaderProps> = (props) => {
    return <h3 className="font-bold text-lg mb-4">{props.children}</h3>;
};

type ModalContentProps = PropsWithChildren;
const ModalContent: React.FC<ModalContentProps> = (props) => {
    return <div className="modal-box">
        {props.children}
    </div>;
};

type ModalActionProps = PropsWithChildren;

const ModalAction: React.FC<ModalActionProps> = (props) => {
    return <div className="modal-action">
        {props.children}
    </div>;
};

type ModalProps = {
    id: string
    open: boolean
} & PropsWithChildren

const Modal: React.FC<ModalProps> = (props) => {
    const classNames = twMerge(
        'modal',
        clsx({
            'modal-open': props.open
        })
    );

    return <dialog
        className={classNames}
        aria-label="Modal"
        aria-modal={props.open}
        aria-hidden={!props.open}
        id={props.id}
    >
        {props.children}
    </dialog>;
};

export default Object.assign(Modal, {
    Header: ModalHeader,
    Content: ModalContent,
    Action: ModalAction
});