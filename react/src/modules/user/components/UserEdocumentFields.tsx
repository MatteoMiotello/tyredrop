import React from "react";
import {FormHandler} from "../../../common/components/shelly-ui/Form/useForm";
import {Input} from "../../../common/components/shelly-ui";
import {useTranslation} from "react-i18next";
import {isEmail, isRequired} from "../../../common/components/shelly-ui/Input";

type UserEdocumentFieldsProps = {
    form: FormHandler
}
const UserEdocumentFields: React.FC<UserEdocumentFieldsProps> = ( {form} ) => {
    const {t} = useTranslation();

    return <>
        <Input.FormControl className="col-span-6">
            <Input.Label>
                {t('billing.sdi_code')}
            </Input.Label>
            <Input
                placeholder={t('billing.sdi_code') as string}
                {...form.registerInput({name: 'sdi_code'})} />
        </Input.FormControl>
        <Input.FormControl className="col-span-6">
            <Input.Label>
                {t('billing.sdi_pec')}
            </Input.Label>
            <Input placeholder={t('billing.sdi_pec') as string}
                   {...form.registerInput({name: 'sdi_pec', validators:[isRequired('La pec è richiesta'), isEmail( 'Il formato della pec non è corretto' )]})}
            />
        </Input.FormControl>
    </>;
};

export default UserEdocumentFields;