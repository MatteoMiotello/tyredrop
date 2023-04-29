import {useQuery} from "@apollo/client";
import axios from "axios";
import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {GetLegalEntityTypesQuery, LegalEntityType} from "../../../__generated__/graphql";
import {GET_LEGAL_ENTITY_TYPES} from "../../../common/backend/graph/query/legal-entities";
import Autocomplete, {AutocompleteQueryHandler} from "../../../common/components-library/Autocomplete";
import Button from "../../../common/components-library/Button";
import Form, {useForm} from "../../../common/components-library/Form";
import Input from "../../../common/components-library/Input";
import {SelectComponent, SelectOption} from "../../../common/components-library/SelectComponent";
import Spinner from "../../../common/components/Spinner";
import {isRequired} from "../../../common/validation/validators";

const controller = new AbortController();
const handleQuery: AutocompleteQueryHandler = async (query: string) => {

    const res = await axios.get(`https://restcountries.com/v3.1/name/${query}`, {
        signal: controller.signal
    });

    return res.data.map((element: any) => {
        return {
            value: element.cca2,
            title: element.translations.ita.common ?? element.name.common
        };
    });
};

export  type BillingInput = {
    entity_type: string
    name: string
    surname: string
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
    store: ( input: BillingInput ) => void
}

export const BillingForm: React.FC<BillingFormProps> = ( props ) => {
    const [form, handleFormError] = useForm();
    const {t} = useTranslation();
    const {loading, error, data} = useQuery<GetLegalEntityTypesQuery>(GET_LEGAL_ENTITY_TYPES);
    const [options, setOptions] = useState<SelectOption[]>([]);

    const onSubmit = (billingInput: BillingInput) => {


        props.store( billingInput ) ;
        return;
    };

    if (error) {
        handleFormError(t('billing.error_loading'));
    }

    useEffect( () => {
        if (data) {
            if (!data.legalEntityTypes) {
                setOptions([]);
            } else {
                setOptions(data.legalEntityTypes?.map((type: LegalEntityType) => {
                        return {
                            value: type.id,
                            title: type.name,
                            disabled: false
                        };
                    }
                ));
            }
        }
    }, [data] );


    return <Form onSubmit={onSubmit} form={form} className="lg:w-1/2 relative">
        {loading && <Spinner/>}
        <SelectComponent
            options={ options }
            className="col-span-12"
            name="entity_type"
        />
        <Input name="name"
        type="text"
        placeholder={t('billing.name_placeholder')}
        className="col-span-6"
        validators={[isRequired]}
        />
        <Input name="surname"
        type="text"
        placeholder={t('billing.surname_placeholder')}
        className="col-span-6"
        validators={[isRequired]}
        />
        <Input type="text"
        name="fiscal_code"
        placeholder={t('billing.fiscal_code_placeholder')}
        className="col-span-12"
        />
        <Input type="text"
        name="vat_number"
        placeholder={t('billing.vat_number_placeholder')}
        className="col-span-12"
        />
        <Input type="text"
        name="address_line_1"
        placeholder={t('billing.address_line_1_placeholder')}
        className="col-span-12"
        validators={[isRequired]}
        />
        <Input type="text"
        name="address_line_2"
        placeholder={t('billing.address_line_2_placeholder')}
        className="col-span-12"
        />
        <Autocomplete name="country"
        className="col-span-12"
        getOptions={handleQuery}
        initialOptions={[ { value: "IT", title: "Italia"} ]}
        validators={[isRequired]}
        placeholder={t('billing.country_placeholder') as string}

        />
        <Input type="text"
        name="city"
        placeholder={t('billing.city_placeholder')}
        className="col-span-4"
        validators={[isRequired]}
        />
        <Input type="text"
        name="province"
        placeholder={t('billing.province_placeholder')}
        className="col-span-4"
        validators={[isRequired]}
        />
        <Input type="text"
        name="postal_code"
        placeholder={t('billing.cap_placeholder')}
        className="col-span-4"
        validators={[isRequired]}
        />

        <h4 className="col-span-12 text-center"> Dati di pagamento </h4>
        <Input
            className="col-span-12"
            type="text"
            name="iban"
            placeholder="IBAN" />
        <Button
        type={"primary"}
        htmlType={"submit"}
        className="col-start-9 col-span-4"
        >
        {t('billing.submit_button')}
    </Button>
</Form>;
};