import React from "react";
import {useTranslation} from "react-i18next";
import {BillingForm} from "./components/BillingForm";

export const BillingPage: React.FC = () => {
    const {t} = useTranslation();
    return <div className="flex flex-col justify-center items-center my-auto">
        <h1 className="my-10">{t('billing.page_title')}</h1>
        <BillingForm/>
    </div>;
};