import React, {useEffect} from "react";
import {Order, UpdateOrderStatusMutation, UpdateOrderStatusMutationVariables} from "../../../../__generated__/graphql";
import {useMutation} from "../../../../common/backend/graph/hooks";
import {UPDATE_ORDER_STATUS} from "../../../../common/backend/graph/mutation/order";
import {Button, Form, Input} from "../../../../common/components/shelly-ui";
import {FormHandler} from "../../../../common/components/shelly-ui/Form/useForm";
import OrderStatusSelect from "../../../order/components/OrderStatusSelect";


type OrderStatusFormProps = {
    form: FormHandler
    order: Order
}
const OrderStatusForm: React.FC<OrderStatusFormProps> = ({form, order}) => {
    const [mutate, mutationQuery] = useMutation<UpdateOrderStatusMutation, UpdateOrderStatusMutationVariables>( UPDATE_ORDER_STATUS, {
        errorMessage: "Non e` possibile impostare lo stato selezionato"
    } );
    useEffect(() => {
        if (order) {
            form.setFormValues({
                status: order.status
            });
        }
    }, [order]);

    return <Form form={form} saveForm={ (data) => mutate({
        variables: {
            orderId: order.id,
            newStatus: data.status
        }
    }) }>
        <Input.FormControl>
            <Input.Label>
                Stato
            </Input.Label>
            <OrderStatusSelect order={order} {...form.registerInput({name: 'status'})}/>
        </Input.FormControl>
        <Form.FormButtons>
            <Button type="submit" buttonType="primary" loading={mutationQuery.loading}>
                Salva
            </Button>
        </Form.FormButtons>
    </Form>;
};

export default OrderStatusForm;