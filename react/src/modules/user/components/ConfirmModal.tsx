import React from "react";
import Button from "../../../common/components-library/Button";
import {Modal} from "../../../common/components/shelly-ui";
import {ModalHandler} from "../../../common/components/shelly-ui/Modal/useModal";

type ConfirmDeleteModalProps = {
    modal: ModalHandler
    onConfirm: () => void
    modalTitle?: string
}
const ConfirmModal: React.FC<ConfirmDeleteModalProps> = (props ) => {
    return <Modal modal={props.modal}>
        <Modal.Title>
            {
                props.modalTitle || "Confermi di voler eseguire l'azione richiesta?"
            }
        </Modal.Title>
        <Modal.Actions>
            <Button onClick={props.modal.close}>
                Annulla
            </Button>
            <Button type="primary" onClick={() => {
                props.onConfirm();
                props.modal.close();
            }}>
                Conferma
            </Button>
        </Modal.Actions>
    </Modal>;
};

export default ConfirmModal;