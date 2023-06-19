import React from "react";
import Button from "../../../common/components-library/Button";
import Modal from "../../../common/components-library/Modal";

type ConfirmDeleteModalProps = {
    closeModal:() => void,
    onConfirm: () => void
}
const ConfirmDeleteModal: React.FC<ConfirmDeleteModalProps> = ( props ) => {
    return <Modal.Content>
        <Modal.Header>
            Confermi di voler eliminare l'elemento selezionato?
        </Modal.Header>
        <Modal.Action>
            <Button>

            </Button>
        </Modal.Action>
    </Modal.Content>;
};

export default ConfirmDeleteModal;