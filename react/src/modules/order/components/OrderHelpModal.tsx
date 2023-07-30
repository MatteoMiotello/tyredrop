import React from "react";
import {Button, Modal} from "../../../common/components/shelly-ui";
import {ModalHandler} from "../../../common/components/shelly-ui/Modal/useModal";

type OrderHelpModalProps = {
    modal: ModalHandler
}

const OrderHelpModal: React.FC<OrderHelpModalProps> = ({modal, order}) => {
    return <Modal modal={modal}>
        <Modal.Title>
            Richiesta assistenza per l'ordine
        </Modal.Title>
        test
        <Modal.Actions>
            <Button onClick={modal.close}>
                Chiudi
            </Button>
        </Modal.Actions>
    </Modal>;
};

export default OrderHelpModal;