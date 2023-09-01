import {ReactNode} from "react";
import {ProductSpecificationValue} from "../../../__generated__/graphql";
import {ProductSpecificationsSet} from "../enums/product-specifications-set";
import sunny from "../../../assets/icons/sunny.png";
import snowy from "../../../assets/icons/snowy.png";
import snow from "../../../assets/icons/snow.png";

export type ProviderHandler = (value: ProductSpecificationValue) => {
    title: string,
    content: ReactNode
}

export const specificationProvider = (value: ProductSpecificationValue) => {
    switch (value.specification.code) {
        case ProductSpecificationsSet.TYRE.RUNFLAT:
            return runflatProvider(value);
        case ProductSpecificationsSet.TYRE.SEASON:
            return seasonProvider(value);
        default:
            return {
                title: value.value,
                content: value.value
            };
    }
};

export const runflatProvider: ProviderHandler = (value: ProductSpecificationValue) => {
    if (value.value == 'true') {
        return {
            title: 'Sì',
            content: 'Sì',
        };
    }

    if (value.value == 'false') {
        return {
            title: 'No',
            content: 'No',
        };
    }

    return {
        title: value.value,
        content: value.value,
    };
};

const seasonProvider: ProviderHandler = (value: ProductSpecificationValue) => {
    switch (value.value) {
        case 'SUMMER':
            return {
                title: 'Estate',
                content: <div className="flex items-center"><img src={sunny} className="h-6 mr-2"/> Estate </div>
            };
        case 'WINTER':
            return {
                title: 'Inverno',
                content: <div className="flex items-center"><img src={snow} className="h-6 mr-2"/> Inverno </div>
            };
        case 'ALL_SEASON':
            return {
                title: 'All season',
                content: <div className="flex items-center"><img src={snowy} className="h-6 mr-2"/> All season </div>
            };
        default:
            return {
                title: value.value,
                content: value.value,
            };
    }
};