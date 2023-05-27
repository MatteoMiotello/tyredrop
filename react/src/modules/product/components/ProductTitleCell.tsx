import React from "react";
import {Img} from "react-image";
import {Link} from "react-router-dom";
import Spinner from "../../../common/components/Spinner";
import ProdapiService from "../services/prodapi/prodapi-service";
import {ProductRowItemData} from "./ProductTable";


type TyreItemRowProps = {
    data: ProductRowItemData
}

const ProductTitleCell: React.FC<TyreItemRowProps | null> = (props: TyreItemRowProps | null) => {
    if (!props?.data) {
        return <></>;
    }

    const data = props.data;

    return <Link className="h-36 flex w-full items-center" to={"/products/details/" + data.id}>
        <div>
            <div className="break-words font-semibold text-xl"> {data.name} </div>
            <div className="mt-2 flex items-center">
                <div className="border-2 rounded-full aspect-square max-h-12 flex justify-center">
                    <Img src={(new ProdapiService()).getBrandImageUrl(data.brand.code)}
                         loading="lazy"
                         className="my-auto"
                         loader={<Spinner/>}
                         onErrorCapture={(e) => e.preventDefault()}
                    />
                </div>
                <span className="ml-2 font-semibold text-slate-400">
                 {data.brand.name}
                </span>
            </div>
        </div>
    </Link>;
};

export default ProductTitleCell;