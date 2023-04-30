import {useQuery} from "@apollo/client";
import axios from "axios";
import {GraphQLError} from "graphql/error";
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
import {ValidationHandler, isRequired, maxCharacters, minCharacters} from "../../../common/validation/validators";
import {Simulate} from "react-dom/test-utils";
import {useToast} from "../../../hooks/useToast";
import load = Simulate.load;

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
    store: (input: BillingInput) => Promise<any>
}

export const BillingForm: React.FC<BillingFormProps> = (props) => {
    const [form, handleFormError] = useForm();
    const {t} = useTranslation();
    const {loading, error, data} = useQuery<GetLegalEntityTypesQuery>(GET_LEGAL_ENTITY_TYPES);
    const [options, setOptions] = useState<SelectOption[]>([]);
    const {setError} = useToast();

    const [ vatNumber, setVatNumber ] = useState<string|null>( null );
    const [ fiscalCode, setFiscalCode ] = useState<string|null>( null );

    const onSubmit = (billingInput: BillingInput) => {
        props.store(billingInput)
            .catch( ( err: GraphQLError ) => {
                handleFormError( err.toString() );
            } );
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
    }, [data]);

    const validateFiscalCode: ValidationHandler = ( value: string | null ) => {
        
    };
    const validateVatNumber: ValidationHandler = ( value: string | null ) => {

    };


    return <Form onSubmit={onSubmit} form={form} className="lg:w-1/2 relative">
        {loading && <Spinner/>}
        <SelectComponent
            options={options}
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
               validators={[minCharacters(16), maxCharacters(16)]}
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
                      initialOptions={[{value: "IT", title: "Italia"}]}
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
               validators={[maxCharacters(2), isRequired, minCharacters(2)]}
        />
        <Input type="text"
               name="cap"
               placeholder={t('billing.cap_placeholder')}
               className="col-span-4"
               validators={[isRequired]}
        />

        <h4 className="col-span-12 text-center"> Dati di pagamento </h4>
        <Input
            className="col-span-12"
            type="text"
            name="iban"
            placeholder="IBAN"/>
        <Button
            type={"primary"}
            htmlType={"submit"}
            className="col-start-9 col-span-4"
        >
            {t('billing.submit_button')}
        </Button>
    </Form>;
};