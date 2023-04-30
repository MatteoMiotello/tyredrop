import {useMutation} from "@apollo/client";
import React, {useEffect} from "react";
import {useTranslation} from "react-i18next";
import {useNavigate} from "react-router-dom";
import {
    CreateUserBillingMutation,
    CreateUserBillingMutationVariables
} from "../../__generated__/graphql";
import {CREATE_BILLING} from "../../common/backend/graph/mutation/create-billing";
import Spinner from "../../common/components/Spinner";
import {BillingForm, BillingInput} from "./components/BillingForm";
import {useAuth} from "./hooks/useAuth";

const BillingPage: React.FC = () => {
    const {t} = useTranslation();
    const navigate = useNavigate();
    const auth = useAuth();
    const [ saveBilling, { error, data, loading } ] = useMutation<CreateUserBillingMutation, CreateUserBillingMutationVariables>( CREATE_BILLING );

    useEffect( () => {
        if ( !auth.isAuthenticated() ) {
            navigate( '/auth/login' );
        }
    }, [auth] );

    const storeBilling = ( input: BillingInput ) => {
        saveBilling({
            variables: {
                input: {
                    addressLine1: input.address_line_1,
                    addressLine2: input.address_line_2,
                    cap: input.cap,
                    city: input.city,
                    country: input["country[value]"],
                    fiscalCode: input.fiscal_code,
                    iban: input.iban,
                    legalEntityTypeId: input.entity_type,
                    name: input.name,
                    province: input.province,
                    surname: input.surname,
                    vatNumber: input.vat_number as string,
                }
            }
        });
    };

    // useEffect( () => {
    //     if ( data ) {
    //         navigate( '/' );
    //     }
    // }, [data] )

    return <div className="flex flex-col justify-center items-center my-auto">
        { loading ?? <Spinner /> }
        <div className="text-center">
        <h1>{t('billing.page_title')}</h1>
        <p>{t('billing.page_subtitle')}</p>
        </div>
        <BillingForm store={storeBilling}/>
    </div>;
};

export default BillingPage;