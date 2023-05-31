import React from "react";
import {Img} from "react-image";
import {Link} from "react-router-dom";
import Spinner from "../../../common/components/Spinner";
import ProdapiService from "../services/prodapi/prodapi-service";


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
            <span className="text-base-300 font-semibold">
                cod. {data.code}
            </span>
            {
                props.showBrand &&
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
            }
        </div>
    </div>;
};

export default ProductTitle;