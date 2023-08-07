import React from "react";
import {Link} from "react-router-dom";


type TyreItemRowProps = {
    data: {
        id: string
        name: string
        code: string
        brand: {
            code: string
            name: string
        }
    }
    showBrand?: boolean
}

const ProductTitle: React.FC<TyreItemRowProps | null> = (props: TyreItemRowProps | null) => {
    if (!props?.data) {
        return <></>;
    }

    const data = props.data;

    return <div className="h-36 flex w-full items-center" >
        <div className="flex flex-col">
            <Link to={"/products/details/" + data.id} className="break-words font-semibold text-xl"> {data.name} </Link>
            <span className="text-secondary font-semibold">
                cod. EAN {data.code}
            </span>
        </div>
    </div>;
};

export default ProductTitle;