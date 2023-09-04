import React, {useEffect} from "react";
import {Form} from "../../../common/components/shelly-ui";
import {FormHandler} from "../../../common/components/shelly-ui/Form/useForm";
import UserEdocumentFields from "./UserEdocumentFields";
import {
    UpdateUserBillingMutation,
    UpdateUserBillingMutationVariables,
    UserBilling
} from "../../../__generated__/graphql";
import {useMutation} from "../../../common/backend/graph/hooks";
import {UPDATE_USER_BILLING} from "../../../common/backend/graph/mutation/users";

type UserEdocumentFormProps = {
    form: FormHandler
    userBilling?: UserBilling
}
const UserEdocumentForm: React.FC<UserEdocumentFormProps> = ( {form, userBilling} ) => {
    const [mutate] = useMutation<UpdateUserBillingMutation, UpdateUserBillingMutationVariables>( UPDATE_USER_BILLING );

    useEffect( () => {
        if ( !userBilling ) {
            return;
        }

        form.setFormValues({
            sdi_code: userBilling.sdiCode,
            sdi_pec: userBilling.sdiPec,
        });
    }, [userBilling] );

    return <Form form={form} saveForm={(data) => {
        if (!userBilling) {
            return false;
        }

        return mutate({
            variables: {
                billingID: userBilling?.id,
                edocumentInput: {
                    sdiCode: data.sdi_code,
                    sdiPec: data.sdi_pec
                }
            }
        });
    }}>
        <UserEdocumentFields form={form}/>
    </Form>;
};

export default UserEdocumentForm;