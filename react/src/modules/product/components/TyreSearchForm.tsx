import React, {useEffect, useRef, useState} from "react";
import Field from "../../../common/components-library/Input";
import Button from "../../../common/components-library/Button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faSearch} from "@fortawesome/free-solid-svg-icons";
import Form, {useForm} from "../../../common/components-library/Form";
import {useTranslation} from "react-i18next";
import BrandField from "../../../common/components/BrandField";
import {ProductSearchDataType} from "../context/product-search-context";
import {ProductSpecificationsSet} from "../enums/product-specifications-set";
import SpecificationField from "./SpecificationField";

type TypeSpecificSearchFormProps = {
    onSubmit: (req: TyreSearchFormRequest, vehicleCode: string) => void
}

export type TyreSearchFormRequest = {
    tyre_code?: string
    tyre_name?: string
    vehicle_code?: string
    width?: number
    aspect_ratio?: number
    rim?: number
    season?: string
    brand?: string
}

export const toSearchDataType = (req: TyreSearchFormRequest, vehicleCode: string): ProductSearchDataType => {
    return {
        code: req.tyre_code?.length ? req.tyre_code : null,
        name: req.tyre_name?.length ? req.tyre_name : null,
        vehicleCode: vehicleCode,
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        brand: req.brand?.length ? req.brand : null,
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
            },
            {
                code: ProductSpecificationsSet.TYRE.SEASON,
                value: req.season ? req.season.toString() : null
            }
        ]
    };
};

const TyreSearchForm: React.FC<TypeSpecificSearchFormProps> = (props) => {
    const {t} = useTranslation();
    const [vehicle, setVehicle] = useState("CAR");
    const {form} = useForm();
    const formRef = useRef<HTMLFormElement>(null);

    useEffect(() => {
        formRef.current?.requestSubmit();
    }, [vehicle]);

    const isActive = (code: string) => {
        return code == vehicle;
    };

    return <Form
        onSubmit={(req) => props.onSubmit(req, vehicle)}
        form={form}
        ref={formRef}
    >
        <div className="col-span-12 flex justify-center">
            <div className="tabs tabs-boxed">
                <button type="button" className={`tab ${isActive("CAR") && "tab-active"}`}
                        onClick={() => setVehicle("CAR")}>Auto
                </button>
                <button type="button" className={`tab ${isActive("MOTO") && "tab-active"}`}
                        onClick={() => setVehicle("MOTO")}>Moto
                </button>
                <button type="button" className={`tab ${isActive("TRUCK") && "tab-active"}`}
                        onClick={() => setVehicle("TRUCK")}>Camion
                </button>
            </div>
        </div>
        <Field.FormControl className="col-span-3">
            <Field.Label className="text-base-100">
                {t('searchbar.width_field_label')}
            </Field.Label>
            <SpecificationField
                name="width"
                specificationCode={ProductSpecificationsSet.TYRE.WIDTH}
                vehicleCode={vehicle}
                placeholder={t('searchbar.width_field_label') as string}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-3">
            <Field.Label className="text-base-100">
                {t('searchbar.height_field_label')}
            </Field.Label>
            <SpecificationField
                name="aspect_ratio"
                specificationCode={ProductSpecificationsSet.TYRE.ASPECT_RATIO}
                vehicleCode={vehicle}
                placeholder={t('searchbar.height_field_label') as string}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-3">
            <Field.Label className="!text-base-100">
                {t('searchbar.rim_field_label')}
            </Field.Label>
            <SpecificationField
                name="rim"
                specificationCode={ProductSpecificationsSet.TYRE.RIM}
                vehicleCode={vehicle}
                placeholder={t('searchbar.rim_field_label') as string}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-3">
            <Field.Label className="text-base-100">
                {t('searchbar.season_field_label')}
            </Field.Label>
            <SpecificationField
                name="season"
                specificationCode={ProductSpecificationsSet.TYRE.SEASON}
                vehicleCode={vehicle}
                placeholder={t('searchbar.season_field_label') as string}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-4">
            <Field.Label className="!text-base-100">
                {t('searchbar.code_field_label')}
            </Field.Label>
            <Field.Input
                type="text"
                name="tyre_code"
                placeholder={t('searchbar.code_field_label')}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-4">
            <Field.Label className="!text-base-100">
                {t('searchbar.name_field_label')}
            </Field.Label>
            <Field.Input
                type="text"
                name="tyre_name"
                placeholder={t('searchbar.name_field_label')}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-4">
            <Field.Label className="!text-base-100">
                {t('searchbar.brand_field_label')}
            </Field.Label>
            <BrandField
                name="brand"
            />
        </Field.FormControl>
        <Button className="col-start-6 col-span-2 btn-outline "
                size="sm"
                type="ghost">
            <FontAwesomeIcon icon={faSearch}/>
        </Button>
    </Form>;
};

export default TyreSearchForm;