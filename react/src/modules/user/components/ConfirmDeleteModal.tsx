import React from "react";
import Button from "../../../common/components-library/Button";
import {Modal} from "../../../common/components/shelly-ui";
import {ModalHandler} from "../../../common/components/shelly-ui/Modal/useModal";

type ConfirmDeleteModalProps = {
    modal: ModalHandler
    onConfirm: () => void
}
const ConfirmDeleteModal: React.FC<ConfirmDeleteModalProps> = ( props ) => {
    return <Modal modal={props.modal}>
        <Modal.Title>
            Confermi di voler eliminare l'elemento selezionato?
        </Modal.Title>
        <Modal.Actions>
            <Button onClick={props.modal.close}>
                Annulla
            </Button>
            <Button type="primary" onClick={props.onConfirm}>
                Conferma
            </Button>
        </Modal.Actions>
    </Modal>;
};

export default ConfirmDeleteModal;