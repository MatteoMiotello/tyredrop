import React from "react";
import {ProductSpecificationValue} from "../../../__generated__/graphql";

type CompleteProductSpecificationsGroupProps = {
    specifications: ProductSpecificationValue[]
}

const CompleteProductSpecificationsGroup: React.FC<CompleteProductSpecificationsGroupProps> = ( props ) => {
    return <div className="flex flex-wrap w-full">
        {
            props.specifications.map( ( spec, key ) => {
                return ( <div key={key} className="flex flex-col m-1 border-2 p-2 rounded-box">
                    <span className="text-sm text-secondary font-semibold flex-1"> { spec.specification.name } </span>
                    <span className="font-semibold"> { spec.value } </span>
                </div> );
            } )
        }
    </div>;
};

export default CompleteProductSpecificationsGroup;