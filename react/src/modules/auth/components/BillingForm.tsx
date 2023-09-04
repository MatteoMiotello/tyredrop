import {ApolloError} from "@apollo/client";
import React, {useEffect} from "react";
import {useTranslation} from "react-i18next";
import {
    GetLegalEntityTypesQuery
} from "../../../__generated__/graphql";
import {GET_LEGAL_ENTITY_TYPES} from "../../../common/backend/graph/query/legal-entities";
import Button from "../../../common/components-library/Button";
import Spinner from "../../../common/components/Spinner";
import {
    isRequired
} from "../../../common/validation/validators";
import {useToast} from "../../../store/toast";
import {useQuery} from "../../../common/backend/graph/hooks";
import UserBillingFields from "../../user/components/UserBillingFields";
import {Form, Input, useForm} from "../../../common/components/shelly-ui";
import UserEdocumentFields from "../../user/components/UserEdocumentFields";
import {maxCharacters, minCharacters} from "../../../common/components/shelly-ui/Input";


export type BillingInput = {
    entity_type: string
    name: string
    surname?: string | undefined
    fiscal_code?: string | null
    vat_number?: string | null
    address_line_1: string
    address_line_2?: string
    country: string
    city: string
    province: string
    postal_code: string
    cap: string
    iban: string
    sdi_code: string
    sdi_pec: string
}


type BillingFormProps = {
    store: (input: BillingInput) => Promise<any>
}

export const BillingForm: React.FC<BillingFormProps> = (props) => {
    const {loading, error, data} = useQuery<GetLegalEntityTypesQuery>(GET_LEGAL_ENTITY_TYPES);

    const form = useForm();
    const {t} = useTranslation();
    const toast = useToast();

    const onSubmit = (billingInput: BillingInput) => {
        return props.store(billingInput)
            .catch((err: ApolloError) => {
                if (!err.graphQLErrors || !err.graphQLErrors.length || !err.graphQLErrors[0].extensions) {
                    form.handleFormError(err.message);
                    return;
                }

                const errorCode = err.graphQLErrors[0].extensions.code;

                if (!errorCode) {
                    form.handleFormError(err.graphQLErrors[0].message);
                    return;
                }

                if (errorCode == "4004") {
                    form.handleFormError(t("billing.user_not_found") as string);
                    return;
                }

                if (errorCode == "5001") {
                    form.handleFormError(t("billing.error_storing") as string);
                    return;
                }

                form.handleFormError(err.graphQLErrors[0].message);
            });
    };

    useEffect(() => {
        if (error) {
            toast.error(t('billing.error_loading'));
        }
    }, [error]);


    return <Form form={form} saveForm={ (data: BillingInput) => onSubmit(data)}>
        <Form.GridLayout>
            <h4 className="col-span-12 divider"> Dati di fatturazione </h4>
            {loading && <Spinner/>}
            <UserBillingFields form={form}/>
            <h4 className="col-span-12 divider"> Dati di fatturazione elettronica </h4>
            <UserEdocumentFields form={form}/>

            <h4 className="col-span-12 divider"> Dati di pagamento </h4>
            <Input.FormControl className="col-span-12">
                <Input.Label>
                    {t('billing.iban_placeholder')}
                </Input.Label>
                <Input
                    placeholder={t('billing.iban_placeholder') as string}
                    {...form.registerInput({
                        name: 'iban',
                        validators: [isRequired(t('billing.iban_placeholder')), minCharacters(27, 'Formato iban non corretto'), maxCharacters( 34, 'Formato iban non corretto' )]
                    })}/>
            </Input.FormControl>
            <Button
                type={"primary"}
                htmlType={"submit"}
                className="col-start-9 col-span-4"
            >
                {t('billing.submit_button')}
            </Button>
        </Form.GridLayout>

    </Form>;
};