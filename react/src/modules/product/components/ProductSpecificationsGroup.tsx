import {ProductSpecificationDefinition} from "../enums/product-specifications-set";
import React from "react";
import Badge from "../../../common/components-library/Badge";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";


type ProductSpecificationsGroupProps = {
    specifications: (ProductSpecificationDefinition|null)[]
}
const ProductSpecificationsGroup: React.FC<ProductSpecificationsGroupProps> = (props) => {
    return <div>
        {
            props.specifications.map( ( spec, key ) => {
                if ( spec && spec.icon) {
                    return <Badge outline={true} key={key} type={"secondary"} className="mx-1">
                        <FontAwesomeIcon icon={spec.icon}/>
                        <span className="ml-1">{spec.value}</span>
                    </Badge>;
                }
            } )
        }
    </div>;
};

export default ProductSpecificationsGroup;