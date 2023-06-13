import {ApolloError, useQuery} from "@apollo/client";
import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {
    GetLegalEntityTypesQuery,
    LegalEntityType
} from "../../../__generated__/graphql";
import {GET_LEGAL_ENTITY_TYPES} from "../../../common/backend/graph/query/legal-entities";
import Button from "../../../common/components-library/Button";
import Form, {useForm} from "../../../common/components-library/Form";
import Field from "../../../common/components-library/Input";
import {SelectComponent, SelectOption} from "../../../common/components-library/SelectComponent";
import CountryField from "../../../common/components/CountryField";
import Spinner from "../../../common/components/Spinner";
import {ValidationHandler, isRequired, maxCharacters, minCharacters} from "../../../common/validation/validators";
import {useToast} from "../../../hooks/useToast";


export type BillingInput = {
    entity_type: string
    name: string
    surname?: string | undefined
    fiscal_code?: string | null
    vat_number?: string | null
    address_line_1: string
    address_line_2?: string
    "country[title]": string
    "country[value]": string
    city: string
    province: string
    postal_code: string
    cap: string
    iban: string
}


type BillingFormProps = {
    store: (input: BillingInput) => Promise<any>
}

export const BillingForm: React.FC<BillingFormProps> = (props) => {
    const [isPerson, setIsPerson] = useState<boolean>(true);
    const [options, setOptions] = useState<(SelectOption | null)[]>([]);

    const {loading, error, data} = useQuery<GetLegalEntityTypesQuery>(GET_LEGAL_ENTITY_TYPES);

    const {form, handleFormError} = useForm();
    const {t} = useTranslation();
    const {setError} = useToast();

    const onSubmit = (billingInput: BillingInput) => {
        props.store(billingInput)
            .catch((err: ApolloError) => {
                if (!err.graphQLErrors || !err.graphQLErrors.length || !err.graphQLErrors[0].extensions) {
                    handleFormError(err.message);
                    return;
                }

                const errorCode = err.graphQLErrors[0].extensions.code;

                if (!errorCode) {
                    handleFormError(err.graphQLErrors[0].message);
                    return;
                }

                if (errorCode == "4004") {
                    handleFormError(t("billing.user_not_found"));
                    return;
                }

                if (errorCode == "5001") {
                    handleFormError(t("billing.error_storing"));
                    return;
                }

                handleFormError(err.graphQLErrors[0].message);
            });

        return;
    };

    useEffect(() => {
        if (error) {
            setError(t('billing.error_loading'));
        }
    }, [error]);

    useEffect(() => {
        if (data && !loading) {
            if (!data.legalEntityTypes) {
                setOptions([]);
            } else {
                setOptions(data.legalEntityTypes?.map((type: LegalEntityType | null) => {
                        if (!type) {
                            return null;
                        }

                        return {
                            value: type.id,
                            title: type.name,
                            disabled: false
                        };
                    }
                ));
            }
        }
    }, [data]);


    const validateEntityType: ValidationHandler<SelectOption | null> = (option): string | null => {
        if (!option) {
            return null;
        }

        if (!data || !data?.legalEntityTypes) {
            return null;
        }

        const currentOption = data?.legalEntityTypes.find((opt: LegalEntityType | null) => {
            return opt?.id === option.value;
        });

        if (currentOption) {
            setIsPerson(currentOption.isPerson);
        }

        return null;
    };


    return <Form onSubmit={onSubmit} form={form} className="lg:w-1/2 relative">
        {loading && <Spinner/>}
        <SelectComponent
            options={options}
            className="col-span-12"
            name="entity_type"
            validators={[validateEntityType]}
        />
        <Field.FormInput name="name"
                         type="text"
                         placeholder={t('billing.name_placeholder')}
                         className={isPerson ? "col-span-6" : "col-span-12"}
                         validators={[isRequired(t('billing.name_placeholder'))]}
        />
        {isPerson &&
            <Field.FormInput name="surname"
                             type="text"
                             placeholder={t('billing.surname_placeholder')}
                             className="col-span-6"
            />
        }
        <Field.FormInput type="text"
                         name="fiscal_code"
                         placeholder={t('billing.fiscal_code_placeholder')}
                         className="col-span-12"
                         validators={[minCharacters(16), maxCharacters(16), isRequired(t('billing.fiscal_code_placeholder'))]}
        />
        <Field.FormInput type="text"
                         name="vat_number"
                         placeholder={t('billing.vat_number_placeholder')}
                         className="col-span-12"
        />
        <Field.FormInput type="text"
                         name="address_line_1"
                         placeholder={t('billing.address_line_1_placeholder')}
                         className="col-span-12"
                         validators={[isRequired(t('billing.address_line_1_placeholder'))]}
        />
        <Field.FormInput type="text"
                         name="address_line_2"
                         placeholder={t('billing.address_line_2_placeholder')}
                         className="col-span-12"
        />
        <CountryField name="country"
                      className="col-span-12"
        />
        <Field.FormInput type="text"
                         name="city"
                         placeholder={t('billing.city_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('billing.city_placeholder'))]}
        />
        <Field.FormInput type="text"
                         name="province"
                         placeholder={t('billing.province_placeholder')}
                         className="col-span-4"
                         validators={[maxCharacters(2), isRequired(t('billing.province_placeholder')), minCharacters(2)]}
        />
        <Field.FormInput type="text"
                         name="cap"
                         placeholder={t('billing.cap_placeholder')}
                         className="col-span-4"
                         validators={[isRequired(t('billing.cap_placeholder'))]}
        />

        <h4 className="col-span-12 text-center"> Dati di pagamento </h4>
        <Field.FormInput
            className="col-span-12"
            type="text"
            name="iban"
            placeholder={t('billing.iban_placeholder')}
            validators={[isRequired(t('billing.iban_placeholder'))]}
        />
        <Button
            type={"primary"}
            htmlType={"submit"}
            className="col-start-9 col-span-4"
        >
            {t('billing.submit_button')}
        </Button>
    </Form>;
};