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
                color: '#ebcc34',
            };
        case 4:
            return {
                title: "Quality",
                color: '#ebcc34',
            };
        case 5:
            return {
                title: "Premium",
                color: "#56eb34"
            };
        default:
            return {
                title: "Budget",
                color: '#ffaa00',
            };
    }
};

type ProductQualityBadgeProps = {
    quality: number | null | undefined
}
const ProductQualityBadge: React.FC<ProductQualityBadgeProps> = ( {quality} ) => {
    const [config, setConfig] = useState<QualityDef>();

    useEffect( () => {
        setConfig( getQualityConfig( quality ) );

    }, [quality] );

    return <Badge outline={true} className="badge-lg" color={config?.color}>
        {config?.title}
    </Badge>;
};

export default ProductQualityBadge;