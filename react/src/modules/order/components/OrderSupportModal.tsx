import React from "react";
import {Button, Form, Input, Modal, Textarea, useForm} from "../../../common/components/shelly-ui";
import {ModalHandler} from "../../../common/components/shelly-ui/Modal/useModal";
import {Order} from "../../../__generated__/graphql";

type OrderHelpModalProps = {
    modal: ModalHandler
    order?: Order
}

const OrderSupportModal: React.FC<OrderHelpModalProps> = ({modal, order}) => {
    const form = useForm({
        onSuccess: () => modal.close
    });

    return <Modal modal={modal}>
        <Modal.Title>
            Richiesta assistenza per l'ordine: #{order?.id}
        </Modal.Title>
        <Form form={form} saveForm={ ( data ) => {return true;} }>
            <Input.FormControl>
                <Input.Label>
                    Messaggio
                </Input.Label>
                <Textarea {...form.registerInput({name: "message"})}></Textarea>
            </Input.FormControl>
        </Form>
        <Modal.Actions>
            <Button onClick={modal.close}>
                Chiudi
            </Button>
            <Button onClick={ form.submitForm } buttonType="primary">
                Invia richiesta assistenza
            </Button>
        </Modal.Actions>
    </Modal>;
};
export default OrderSupportModal;