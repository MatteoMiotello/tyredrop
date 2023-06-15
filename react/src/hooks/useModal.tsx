import _ from "lodash";
import {ReactNode, useContext, useEffect, useMemo, useState} from "react";
import Modal from "../common/components-library/Modal";
import ModalContext from "../common/contexts/modal-context";

const useModal = (modal: ReactNode | null = null) => {
    const {setModal} = useContext(ModalContext);
    const [modalContent, setModalContent] = useState( modal );
    const [isOpen, setOpen] = useState<boolean>(false);
    const id = useMemo(() => _.toString((new Date()).getTime()), []);

    useEffect(() => {
        setModal(
            {
                content: <Modal
                    open={isOpen}
                    id={id}
                >
                    {modalContent}
                </Modal>,
                id: id
            }
        );

    }, [isOpen, modalContent]);

    const openModal = ( newModal: ReactNode | undefined = undefined) => {
        if (newModal){
            setModalContent(newModal);
        }

        setOpen(true);
    };

    const closeModal = () => {
        setOpen(false);

        setModalContent(null);
    };

    return {openModal, closeModal};
};

export default useModal;