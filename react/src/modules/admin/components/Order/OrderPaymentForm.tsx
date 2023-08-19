import React from "react";
import {Order, PayOrderMutation, PayOrderMutationVariables} from "../../../../__generated__/graphql";
import {useMutation, useQuery} from "../../../../common/backend/graph/hooks";
import {PAY_ORDER} from "../../../../common/backend/graph/mutation/order";
import {ALL_PAYMENT_METHODS} from "../../../../common/backend/graph/query/payments";
import {Button, Form, Input, Select} from "../../../../common/components/shelly-ui";
import {FormHandler} from "../../../../common/components/shelly-ui/Form/useForm";

type OrderPaymentFormProps  = {
    form: FormHandler
    order: Order
}

const OrderPaymentForm: React.FC<OrderPaymentFormProps>= ({form, order}) => {
    const query = useQuery( ALL_PAYMENT_METHODS );
    const [mutate, mutationQuery] = useMutation<PayOrderMutation, PayOrderMutationVariables>( PAY_ORDER );

    return <Form form={form} saveForm={(data) => mutate( {
        variables: {
            orderId: order.id,
            methodCode: data.method,
        }
    } )}>
        <Input.FormControl>
            <Input.Label>
                Seleziona il metodo di pagamento
            </Input.Label>
            <Select
                {...form.registerInput({name: 'method'})}
                options={ query.data?.paymentMethods ? query.data.paymentMethods.map( m => ({value: m?.code as string, title: m?.name as string}) ) : [] }
            />
        </Input.FormControl>
        <Form.FormButtons>
            <Button type="submit" buttonType="primary" loading={mutationQuery.loading}>
                Salva
            </Button>
        </Form.FormButtons>
    </Form>;
};

export default OrderPaymentForm;