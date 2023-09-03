import {useMutation} from "@apollo/client";
import React, {useEffect} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch} from "react-redux";
import {CreateUserBillingMutation, CreateUserBillingMutationVariables} from "../../__generated__/graphql";
import {CREATE_BILLING} from "../../common/backend/graph/mutation/create-billing";
import Spinner from "../../common/components/Spinner";
import {BillingForm, BillingInput} from "../auth/components/BillingForm";
import {useAuth} from "../auth/hooks/useAuth";
import { setUserCompleted} from "../auth/store/auth-slice";
import {useToast} from "../../store/toast";

const BillingPage: React.FC = () => {
    const {t} = useTranslation();
    const auth = useAuth();
    const dispatch = useDispatch();
    const {success} = useToast();
    const [saveBilling, {
        error,
        data,
        loading
    }] = useMutation<CreateUserBillingMutation, CreateUserBillingMutationVariables>(CREATE_BILLING);


    useEffect(() => {
        if (data && !error) {
            dispatch(setUserCompleted());
            success(t("billing.store_success"));
        }
    }, [data, error]);

    useEffect(() => {
        if (auth.isUserCompleted()) {
            auth.tryRefreshToken();
        }
    }, [auth]);

    const storeBilling = (input: BillingInput) => {
        return saveBilling({
            variables: {
                input: {
                    addressLine1: input.address_line_1,
                    addressLine2: input.address_line_2,
                    cap: input.cap,
                    city: input.city,
                    country: input.country,
                    fiscalCode: input.fiscal_code,
                    iban: input.iban,
                    legalEntityTypeId: input.entity_type,
                    name: input.name,
                    province: input.province,
                    surname: input.surname as string,
                    vatNumber: input.vat_number as string,
                    sdiCode: input.sdi_code,
                    sdiPec: input.sdi_pec
                }
            }
        });
    };

    return <div className="flex flex-col justify-center items-center my-auto lg:p-24 p-4">
        {loading ?? <Spinner/>}
        <div className="text-center">
            <h1>{t('billing.page_title')}</h1>
            <p>{t('billing.page_subtitle')}</p>
        </div>
        <BillingForm store={storeBilling}/>
    </div>;
};

export default BillingPage;