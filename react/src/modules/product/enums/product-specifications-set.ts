import {IconDefinition} from "@fortawesome/fontawesome-svg-core";
import {faCloudRain, faGasPump, faVolumeLow} from "@fortawesome/free-solid-svg-icons";

export type ProductSpecificationDefinition = {
    code: string
    value: string
    name?: string
    icon?: IconDefinition | null
}

interface ProductSpecificationIcons {
    [k: string]: IconDefinition
}

export enum ProductCategorySet {
    TYRE='TYRE'
}
enum TyreSpecificationSet {
    NAME = 'NAME',
    REFERENCE = 'REFERENCE',
    WIDTH = 'WIDTH',
    ASPECT_RATIO = 'ASPECT_RATIO',
    CONSTRUCTION = 'CONSTRUCTION',
    RIM = 'RIM',
    LOAD = 'LOAD',
    SPEED = 'SPEED',
    SEASON = 'SEASON',
    EPREL_ID = 'EPREL_ID',
    FUEL_EFFICIENCY = 'FUEL_EFFICIENCY',
    WET_GRIP_CLASS = 'WET_GRIP_CLASS',
    EXTERNAL_ROLLING_NOISE_CLASS = 'EXTERNAL_ROLLING_NOISE_CLASS',
    EXTERNAL_ROLLING_NOISE_LEVEL = 'EXTERNAL_ROLLING_NOISE_LEVEL',
    LOAD_VERSION = 'LOAD_VERSION',
}
export const ProductSpecificationsSet: { [key: string]: { [key: string]: TyreSpecificationSet } } = {
    TYRE: {
        NAME: TyreSpecificationSet.NAME,
        REFERENCE: TyreSpecificationSet.REFERENCE,
        WIDTH: TyreSpecificationSet.WIDTH,
        ASPECT_RATIO: TyreSpecificationSet.ASPECT_RATIO,
        CONSTRUCTION: TyreSpecificationSet.CONSTRUCTION,
        RIM: TyreSpecificationSet.RIM,
        LOAD: TyreSpecificationSet.LOAD,
        SPEED: TyreSpecificationSet.SPEED,
        SEASON: TyreSpecificationSet.SEASON,
        EPREL_ID: TyreSpecificationSet.EPREL_ID,
        FUEL_EFFICIENCY: TyreSpecificationSet.FUEL_EFFICIENCY,
        WET_GRIP_CLASS: TyreSpecificationSet.WET_GRIP_CLASS,
        EXTERNAL_ROLLING_NOISE_CLASS: TyreSpecificationSet.EXTERNAL_ROLLING_NOISE_CLASS,
        EXTERNAL_ROLLING_NOISE_LEVEL: TyreSpecificationSet.EXTERNAL_ROLLING_NOISE_LEVEL,
        LOAD_VERSION: TyreSpecificationSet.LOAD_VERSION
    }
};

const ProductSpecificationIcons: { [k: string]: ProductSpecificationIcons } = {
    TYRE: {
        FUEL_EFFICIENCY: faGasPump,
        WET_GRIP_CLASS: faCloudRain,
        EXTERNAL_ROLLING_NOISE_CLASS: faVolumeLow
    }
};

const getSpecificationIcon = ( categoryCode: ProductCategorySet, code: string ): undefined | IconDefinition => {
    return ProductSpecificationIcons[categoryCode][code];
};

export const ProductSpecifications = {
    getSpecificationIcon: getSpecificationIcon
};