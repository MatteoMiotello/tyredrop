import _ from "lodash";
import {ReactNode, useContext, useEffect, useMemo, useState} from "react";
import Modal from "../common/components-library/Modal";
import ModalContext from "../common/contexts/modal-context";

const useModal = (modal: ReactNode) => {
    const {setModal} = useContext(ModalContext);
    const [isOpen, setOpen] = useState<boolean>(false);
    const id = useMemo(() => _.toString((new Date()).getTime()), []);

    useEffect(() => {
        setModal(
            {
                content: <Modal
                    open={isOpen}
                    id={id}
                >
                    {modal}
                </Modal>,
                id: id
            }
        );

    }, [isOpen]);

    const openModal = () => {
        setOpen(true);
    };

    const closeModal = () => {
        setOpen(false);
    };
    return {openModal, closeModal};
};

export default useModal;