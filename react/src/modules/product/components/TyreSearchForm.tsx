import React from "react";
import Field from "../../../common/components-library/Input";
import Button from "../../../common/components-library/Button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faSearch} from "@fortawesome/free-solid-svg-icons";
import Form, {useForm} from "../../../common/components-library/Form";
import {useTranslation} from "react-i18next";
import BrandField from "../../../common/components/BrandField";
import {ProductSearchDataType} from "../context/product-search-context";
import {ProductSpecificationsSet} from "../enums/product-specifications-set";

type TypeSpecificSearchFormProps = {
    onSubmit: ( req: TyreSearchFormRequest ) => void
}

export type TyreSearchFormRequest = {
    tyre_code?: string
    tyre_name?: string
    width?: number
    aspect_ratio?: number
    rim?: number
    brand?: string
}

export const toSearchDataType = ( req: TyreSearchFormRequest ): ProductSearchDataType => {
    return {
        code: req.tyre_code?.length ? req.tyre_code :  null,
        name: req.tyre_name?.length ? req.tyre_name : null,
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        brand: req["brand[value]"]?.length ? req["brand[value]"] : null,
        specifications: [
            {
                code: ProductSpecificationsSet.TYRE.WIDTH,
                value: req.width ? req.width.toString() : null,
            },
            {
                code: ProductSpecificationsSet.TYRE.RIM,
                value: req.rim ? req.rim.toString() : null,
            },
            {
                code: ProductSpecificationsSet.TYRE.ASPECT_RATIO,
                value: req.aspect_ratio ? req.aspect_ratio.toString() : null
            }
        ]
    };
};

const TyreSearchForm: React.FC<TypeSpecificSearchFormProps> = (props ) => {
    const {t} = useTranslation();
    const [form] = useForm();

    return <Form onSubmit={props.onSubmit} form={form}>
        <Field.FormControl className="col-span-6">
            <Field.Input
                type="text"
                name="tyre_code"
                placeholder={t('searchbar.code_field_label')}
                labelText={t('searchbar.code_field_label')}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-6">
            <Field.Input
                type="text"
                name="tyre_name"
                placeholder={t('searchbar.name_field_label')}
                labelText={t('searchbar.name_field_label')}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-3">
            <Field.Input
                type="number"
                name="width"
                placeholder={t('searchbar.width_field_label')}
                labelText={t('searchbar.width_field_label')}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-3">
            <Field.Input
                type="number"
                name="aspect_ratio"
                placeholder={t('searchbar.height_field_label')}
                labelText={t('searchbar.height_field_label')}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-3">
            <Field.Input
                type="number"
                name="rim"
                placeholder={t('searchbar.rim_field_label')}
                labelText={t('searchbar.rim_field_label')}
            />
        </Field.FormControl>
        <BrandField
            className="col-span-3"
            name="brand"
        />
        <Button className="col-start-6 col-span-2 btn-outline"
                size="sm"
                type="ghost">
            <FontAwesomeIcon icon={faSearch}/>
        </Button>
    </Form>;
};

export default TyreSearchForm;