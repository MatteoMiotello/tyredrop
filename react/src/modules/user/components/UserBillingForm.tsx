import React, {useEffect} from "react";
import {Form} from "../../../common/components/shelly-ui";
import UserBillingFields from "./UserBillingFields";
import {FormHandler} from "../../../common/components/shelly-ui/Form/useForm";
import {
    UpdateUserBillingMutation,
    UpdateUserBillingMutationVariables,
    UserBilling
} from "../../../__generated__/graphql";
import {useMutation} from "../../../common/backend/graph/hooks";
import {UPDATE_USER_BILLING} from "../../../common/backend/graph/mutation/users";

type UserBillingFormProps = {
    form: FormHandler
    userBilling?: UserBilling
}
const UserBillingForm: React.FC<UserBillingFormProps> = ({form, userBilling}) => {
    const [mutation] = useMutation<UpdateUserBillingMutation, UpdateUserBillingMutationVariables>(UPDATE_USER_BILLING);

    useEffect(() => {
        form.setFormValues({
            entity_type: userBilling?.legalEntityType.id,
            name: userBilling?.name,
            surname: userBilling?.surname,
            fiscal_code: userBilling?.fiscalCode,
            vat_number: userBilling?.vatNumber,
            address_line_1: userBilling?.addressLine1,
            address_line_2: userBilling?.addressLine2,
            country: userBilling?.country,
            city: userBilling?.city,
            province: userBilling?.province,
            cap: userBilling?.cap,
        });
    }, [userBilling]);

    return <Form form={form} saveForm={(data) => mutation({
        variables: {
            billingID: userBilling?.id as string,
            input: {
                legalEntityTypeID: data.entity_type,
                name: data.name,
                surname: data.surname,
                fiscalCode: data.fiscal_code,
                vatNumber: data.vat_number,
                addressLine1: data.address_line_1,
                addressLine2: data.address_line_2,
                city: data.city,
                province: data.province,
                cap: data.cap,
                country: data.country,
            }
        }
    })}>
        <Form.GridLayout>
            <UserBillingFields form={form}/>
        </Form.GridLayout>
    </Form>;
};

export default UserBillingForm;