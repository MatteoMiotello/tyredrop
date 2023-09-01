import React from "react";
import {Button, Form, Input, Modal, Textarea, useForm} from "../../../common/components/shelly-ui";
import {ModalHandler} from "../../../common/components/shelly-ui/Modal/useModal";
import {Order, OrderSupportMutation, OrderSupportMutationVariables} from "../../../__generated__/graphql";
import {useMutation} from "../../../common/backend/graph/hooks";
import {ORDER_SUPPORT} from "../../../common/backend/graph/query/order";
import {isRequired} from "../../../common/components/shelly-ui/Input";

type OrderHelpModalProps = {
    modal: ModalHandler
    order?: Order
}

const OrderSupportModal: React.FC<OrderHelpModalProps> = ({modal, order}) => {
    const [mutate, mutationQuery] = useMutation<OrderSupportMutation, OrderSupportMutationVariables>( ORDER_SUPPORT );

    const form = useForm({
        onSuccess: () => modal.close()
    });

    return <Modal modal={modal}>
        <Modal.Title>
            Richiesta assistenza per l'ordine: #{order?.orderNumber}
        </Modal.Title>
        <Form form={form} saveForm={ ( data ) => {
            if ( !data || !order ) {
                return false;
            }

            return mutate( {
                variables: {
                    orderId: order.id,
                    message: data.message
                }
            } );
        } }>
            <Input.FormControl>
                <Input.Label>
                    Messaggio
                </Input.Label>
                <Textarea {...form.registerInput({name: "message", validators:[isRequired( "Il messaggio è richiesto" )]})}></Textarea>
            </Input.FormControl>
        </Form>
        <Modal.Actions>
            <Button onClick={modal.close}>
                Chiudi
            </Button>
            <Button onClick={ form.submitForm } buttonType="primary" loading={mutationQuery.loading}>
                Invia richiesta assistenza
            </Button>
        </Modal.Actions>
    </Modal>;
};
export default OrderSupportModal;