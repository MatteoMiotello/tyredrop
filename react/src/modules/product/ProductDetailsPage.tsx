import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {Img} from "react-image";
import {useLoaderData} from "react-router-dom";
import {ProductItemQuery, ProductSpecificationValue} from "../../__generated__/graphql";
import tyrePlaceholder from "../../assets/placeholder-tyre.jpg";
import Button from "../../common/components-library/Button";
import Field from "../../common/components-library/Input";
import Spinner from "../../common/components/Spinner";
import {Currency} from "../../common/utilities/currency";
import CompleteProductSpecificationsGroup from "./components/CompleteProductSpecificationsGroup";
import ProductTitle from "./components/ProductTitle";
import {ProductCategorySet} from "./enums/product-specifications-set";
import ProdapiService from "./services/prodapi/prodapi-service";

const loadingPlaceholder = <main>
    <Spinner/>
</main>;

const ProductDetailsPage: React.FC = () => {
    const [data, setData] = useState<ProductItemQuery | null>(null);
    const [loading, setLoading] = useState(true);
    const res = useLoaderData() as { data: ProductItemQuery, loading: boolean };
    const {t} = useTranslation();

    useEffect(() => {
        if (res.data) {
            setData(res.data);
        }

        setLoading(res.loading);
    }, [res]);

    if (loading && !data) {
        return loadingPlaceholder;
    }

    return <main className="lg:p-24 p-4">
        <div className="grid md:grid-cols-3 gap-4 w-full items-center">
            <div className="flex justify-center md:justify-normal">
                <Img src={[
                    (new ProdapiService()).getProductImageUrl(data?.productItem?.product.code as string, ProductCategorySet.TYRE),
                    tyrePlaceholder,
                ]}
                     onErrorCapture={(e) => e.preventDefault()}
                     loading="lazy"
                     className="h-96 min-w-fit"
                />
            </div>
            <div>
                <div className="flex justify-between mt-3 items-center">
                    <h2 className="text-2xl font-semibold uppercase"> {data?.productItem?.product.brand.name} </h2>
                    <Img src={(new ProdapiService()).getBrandImageUrl(data?.productItem?.product.brand.code as string)}
                         width={100}
                         loading="lazy"
                         className="my-auto"
                         loader={<Spinner/>}
                         onErrorCapture={(e) => e.preventDefault()}
                    />
                </div>
                <ProductTitle data={
                    {
                        id: data?.productItem?.id as string,
                        name: data?.productItem?.product.name as string,
                        code: data?.productItem?.product.code as string,
                        brand: data?.productItem?.product.brand as { code: string, name: string },
                    }
                }/>
                <CompleteProductSpecificationsGroup
                    specifications={data?.productItem?.product.productSpecificationValues as ProductSpecificationValue[]}
                />
            </div>
            <div className="justify-center bg-base-200 rounded-box p-10 pt-14 md:mt-0 mt-4 h-full">
                <span className="text-primary text-5xl font-semibold">
                    { Currency.defaultFormat( data?.productItem?.price[0]?.value as number, data?.productItem?.price[0]?.currency.iso_code as string ) }
                </span>
                <div className="mt-10 grid grid-cols-3 gap-4">
                    <div className="col-span-2">
                        <Field.Input type="number" name="quantity" placeholder={"4"} labelText="Quantita`"/>
                    </div>
                    <Button type="primary" className="mt-9">
                        { t( 'product_details.order_button' ) }
                    </Button>
                </div>
            </div>
        </div>
    </main>;
};

export default ProductDetailsPage;