import React, {useEffect, useState} from "react";
import Badge from "../../../common/components-library/Badge";

type QualityDef = {
    title: string,
    color: string
}
const getQualityConfig = ( quality: number | undefined | null ): QualityDef => {
    switch (quality) {
        case 3:
            return {
                title: "Budget",
                color: '#ffd138',
            };
        case 4:
            return {
                title: "Quality",
                color: '#66c4ff',
            };
        case 5:
            return {
                title: "Premium",
                color: "#56eb34"
            };
        default:
            return {
                title: "Budget",
                color: '#ffd138',
            };
    }
};

type ProductQualityBadgeProps = {
    quality: number | null | undefined
    small?: boolean
}
const ProductQualityBadge: React.FC<ProductQualityBadgeProps> = ( {quality, small} ) => {
    const [config, setConfig] = useState<QualityDef>();

    useEffect( () => {
        setConfig( getQualityConfig( quality ) );

    }, [quality] );

    return <Badge className={small ? "badge-sm" : "badge-lg"} color={config?.color}>
        {config?.title}
    </Badge>;
};

export default ProductQualityBadge;