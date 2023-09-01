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
import MatchcodeField, {MatchcodeResponse} from "./MatchcodeField";
import SpecificationField from "./SpecificationField";

type TypeSpecificSearchFormProps = {
    onSubmit: (req: TyreSearchFormRequest, vehicleCode?: string) => void
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
    speed?: string
    load?: string
    runflat?: string
    matchcode?: MatchcodeResponse
}

export const toSearchDataType = (req: TyreSearchFormRequest, vehicleCode?: string): ProductSearchDataType => {

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
                value: req.matchcode?.width || (req.width ? req.width.toString() : null),
            },
            {
                code: ProductSpecificationsSet.TYRE.RIM,
                value: req.matchcode?.rim || (req.rim ? req.rim.toString() : null),
            },
            {
                code: ProductSpecificationsSet.TYRE.ASPECT_RATIO,
                value: req.matchcode?.aspect_ratio || (req.aspect_ratio ? req.aspect_ratio.toString() : null)
            },
            {
                code: ProductSpecificationsSet.TYRE.SEASON,
                value: req.season ? req.season.toString() : null
            },
            {
                code: ProductSpecificationsSet.TYRE.SPEED,
                value: req.speed ? req.speed.toString() : null
            },
            {
                code: ProductSpecificationsSet.TYRE.LOAD,
                value: req.load ? req.load.toString() : null
            },
            {
                code: ProductSpecificationsSet.TYRE.RUNFLAT,
                value: req.runflat ? req.runflat.toString() : null,
            }
        ]
    };
};

const TyreSearchForm: React.FC<TypeSpecificSearchFormProps> = (props) => {
    const {t} = useTranslation();
    const [tab, setTab] = useState("CAR");
    const [vehicle, setVehicle] = useState<string | undefined>("CAR");
    const {form} = useForm();
    const formRef = useRef<HTMLFormElement>(null);

    useEffect(() => {
        if (tab == 'MATCHCODE') {
            setVehicle(undefined);
            return;
        }

        setVehicle(tab);
    }, [tab]);

    useEffect(() => {
        formRef.current?.requestSubmit();
    }, [vehicle]);

    const isActive = (code: string) => {
        return code == tab;
    };

    return <Form
        onSubmit={(req) => props.onSubmit(req, vehicle)}
        form={form}
        ref={formRef}
    >
        <div className="col-span-12 flex justify-center">
            <div className="tabs tabs-boxed">
                <button type="button" className={`tab ${isActive("MATCHCODE") && "tab-active"}`}
                        onClick={() => setTab("MATCHCODE")}> Matchcode
                </button>
                <button type="button" className={`tab ${isActive("CAR") && "tab-active"}`}
                        onClick={() => setTab("CAR")}>Auto
                </button>
                <button type="button" className={`tab ${isActive("MOTO") && "tab-active"}`}
                        onClick={() => setTab("MOTO")}>Moto
                </button>
                <button type="button" className={`tab ${isActive("TRUCK") && "tab-active"}`}
                        onClick={() => setTab("TRUCK")}>Camion
                </button>
            </div>
        </div>
        {
            isActive("MATCHCODE") &&
            <Field.FormControl className="col-span-6">
                <Field.Label className="text-base-100">
                    Matchcode
                </Field.Label>
                <MatchcodeField/>
            </Field.FormControl>
        }
        {
            !isActive("MATCHCODE") && <>
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
            </>
        }
        <Field.FormControl className={isActive("MATCHCODE") ? 'col-span-6' : 'col-span-4'}>
            <Field.Label className="!text-base-100">
                {t('searchbar.brand_field_label')}
            </Field.Label>
            <BrandField
                name="brand"
            />
        </Field.FormControl>

        <Field.FormControl className="col-span-4">
            <Field.Label className="text-base-100">
                {t('Indice di carico')}
            </Field.Label>
            <SpecificationField
                name="load"
                specificationCode={ProductSpecificationsSet.TYRE.LOAD}
                vehicleCode={vehicle}
                placeholder={t('Indice di carico') as string}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-4">
            <Field.Label className="text-base-100">
                {t('Indice velocità')}
            </Field.Label>
            <SpecificationField
                name="speed"
                specificationCode={ProductSpecificationsSet.TYRE.SPEED}
                vehicleCode={vehicle}
                placeholder={t('Indice velocità') as string}
            />
        </Field.FormControl>
        <Field.FormControl className="col-span-4">
            <Field.Label className="text-base-100">
                {t('Runflat')}
            </Field.Label>
            <SpecificationField
                name="runflat"
                specificationCode={ProductSpecificationsSet.TYRE.RUNFLAT}
                vehicleCode={vehicle}
                placeholder={t('Runflat') as string}
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