import React, {useEffect, useState} from "react";
import {FormHandler} from "../../../common/components/shelly-ui/Form/useForm";
import {isRequired, maxCharacters, minCharacters} from "../../../common/validation/validators";
import CountryField from "../../../common/components/CountryField";
import {useQuery} from "../../../common/backend/graph/hooks";
import {GetLegalEntityTypesQuery, LegalEntityType} from "../../../__generated__/graphql";
import {GET_LEGAL_ENTITY_TYPES} from "../../../common/backend/graph/query/legal-entities";
import {Input, Select} from "../../../common/components/shelly-ui";
import {useTranslation} from "react-i18next";

type UserBillingFieldsProps = {
    form: FormHandler
}
const UserBillingFields: React.FC<UserBillingFieldsProps> = ({form}) => {
    const {t} = useTranslation();
    const query = useQuery<GetLegalEntityTypesQuery>(GET_LEGAL_ENTITY_TYPES);
    const [entityTypes, setEntityTypes] = useState<LegalEntityType[]>([]);
    const [isPerson, setIsPerson] = useState(true);
    useEffect(() => {
        if (!query.data?.legalEntityTypes) {
            return;
        }

        setEntityTypes(query.data.legalEntityTypes as LegalEntityType[]);
    }, [query.data]);

    useEffect(() => {
        const typeId = form.state.formValues.getFormValue("entity_type");

        if ( !typeId ) {
            return;
        }

        setIsPerson(Boolean(entityTypes.find(t => t.id == typeId)?.isPerson));
    }, [form.state.formValues.formValues]);

    return <>
        <Input.FormControl className="col-span-12">
            <Input.Label>
                Persona giuridica
            </Input.Label>
            <Select
                options={entityTypes.map((t) => ({value: t.id, title: t.name}))}
                {...form.registerInput({name: 'entity_type'})}
            />
        </Input.FormControl>
        <Input.FormControl className={isPerson ? "col-span-6" : "col-span-12"}>
            <Input.Label>
                {
                    isPerson ? 'Nome' : 'Ragione Sociale'
                }
            </Input.Label>
            <Input
                placeholder={isPerson ? 'Nome' : 'Ragione Sociale'}
                {...form.registerInput({
                name: 'name',
                validators: [isRequired(t('billing.name_placeholder'))]
            })}/>
        </Input.FormControl>
        {isPerson &&
            <Input.FormControl className="col-span-6">
                <Input.Label>
                    Cognome
                </Input.Label>
                <Input placeholder="Cognome" {...form.registerInput({name: 'surname'})}/>
            </Input.FormControl>
        }
        <Input.FormControl className="col-span-12">
            <Input.Label>
                {t('billing.fiscal_code_placeholder')}
            </Input.Label>
            <Input placeholder={t('billing.fiscal_code_placeholder') as string} {...form.registerInput({
                name: 'fiscal_code',
                validators: [minCharacters(10), maxCharacters(16), isRequired(t('billing.fiscal_code_placeholder'))]
            })}/>
        </Input.FormControl>
        <Input.FormControl className="col-span-12">
            <Input.Label>
                {t('billing.vat_number_placeholder')}
            </Input.Label>
            <Input placeholder={t('billing.vat_number_placeholder') as string}
                   {...form.registerInput({
                       name: 'vat_number',
                       validators: [minCharacters(11), !isPerson ? isRequired('partita iva') : undefined]
                   })}/>
        </Input.FormControl>
        <Input.FormControl className="col-span-12">
            <Input.Label>
                {t('billing.address_line_1_placeholder')}
            </Input.Label>
            <Input {...form.registerInput({
                name: 'address_line_1',
                validators: [isRequired(t('billing.address_line_1_placeholder'))]
            })}/>
        </Input.FormControl>
        <Input.FormControl className="col-span-12">
            <Input.Label>
                {t('billing.address_line_2_placeholder')}
            </Input.Label>
            <Input
                placeholder={t('billing.address_line_2_placeholder') as string}
                {...form.registerInput({
                    name: 'address_line_2'
                })}/>
        </Input.FormControl>
        <Input.FormControl  className="col-span-12">
            <Input.Label>
                {t('billing.country_placeholder')}
            </Input.Label>
            <CountryField name="country"/>
        </Input.FormControl>
        <Input.FormControl className="col-span-4">
            <Input.Label>
                {t('billing.city_placeholder')}
            </Input.Label>
            <Input
                placeholder={t('billing.city_placeholder') as string}
                {...form.registerInput({
                name: 'city',
                validators: [isRequired(t('billing.city_placeholder'))]
            })}/>
        </Input.FormControl>
        <Input.FormControl className="col-span-4">
            <Input.Label>
                {t('billing.province_placeholder')}
            </Input.Label>
            <Input
                placeholder={t('billing.province_placeholder') as string}
                {...form.registerInput({
                name: 'province',
                validators:[maxCharacters(2), isRequired(t('billing.province_placeholder')), minCharacters(2)]
            })}/>
        </Input.FormControl>
        <Input.FormControl className="col-span-4">
            <Input.Label>
                {t('billing.cap_placeholder')}
            </Input.Label>
            <Input
                placeholder={t('billing.cap_placeholder') as string}
                {...form.registerInput({
                name: 'cap',
                validators: [isRequired(t('billing.cap_placeholder')), minCharacters(5), maxCharacters(5)]
            })}/>
        </Input.FormControl>
    </>;
};

export default UserBillingFields;