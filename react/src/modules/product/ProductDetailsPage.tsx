import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {Img} from "react-image";
import {useLoaderData} from "react-router-dom";
import {ProductItemQuery, ProductSpecificationValue} from "../../__generated__/graphql";
import tyrePlaceholder from "../../assets/placeholder-tyre.jpg";
import Button from "../../common/components-library/Button";
import Field from "../../common/components-library/Input";
import Panel from "../../common/components-library/Panel";
import Spinner from "../../common/components/Spinner";
import {Currency} from "../../common/utilities/currency";
import CompleteProductSpecificationsGroup from "./components/CompleteProductSpecificationsGroup";
import ProductTitle from "./components/ProductTitle";
import {ProductCategorySet} from "./enums/product-specifications-set";
import ProdapiService from "./services/prodapi/prodapi-service";
import {useDispatch} from "react-redux";
import {addCartItem} from "../cart/store/cart-slice";
import {ThunkDispatch} from "redux-thunk";
import {useToast} from "../../hooks/useToast";

const loadingPlaceholder = <main>
    <Spinner/>
</main>;

const ProductDetailsPage: React.FC = () => {
    const [data, setData] = useState<ProductItemQuery | null>(null);
    const [quantity, setQuantity] = useState<number>(4);
    const [loading, setLoading] = useState(true);
    const res = useLoaderData() as { data: ProductItemQuery, loading: boolean };
    const {t} = useTranslation();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const {setSuccess} = useToast();

    useEffect(() => {
        if (res.data) {
            setData(res.data);
        }

        setLoading(res.loading);
    }, [res]);

    if (loading && !data) {
        return loadingPlaceholder;
    }

    console.log(data?.productItem?.product);

    return <main className="lg:px-24 p-4">
        <div className="grid md:grid-cols-4 gap-4 w-full">
            <Panel className="flex flex-col justify-center items-center">
                <Img src={[
                    (new ProdapiService()).getProductImageUrl(data?.productItem?.product.code as string, ProductCategorySet.TYRE),
                    tyrePlaceholder,
                ]}
                     onErrorCapture={(e) => e.preventDefault()}
                     loading="lazy"
                     className="w-fit"
                />
            </Panel>
            <Panel className="col-span-2 row-span-2">
                <div>
                    <div className="flex justify-between mt-3 items-center">
                        <h2 className="text-2xl font-semibold uppercase"> {data?.productItem?.product.brand.name} </h2>
                        <Img
                            src={(new ProdapiService()).getBrandImageUrl(data?.productItem?.product.brand.code as string)}
                            width={100}
                            loading="lazy"
                            className="my-auto"
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
            </Panel>
            <Panel>
                <div className="rounded-box pt-14 md:mt-0 mt-4">
                    <div className="m-5">
                        <div>
                            Totale
                        </div>
                        <span className="text-primary text-5xl font-semibold">
                        {Currency.defaultFormat(data?.productItem?.price[0]?.value as number, data?.productItem?.price[0]?.currency.iso_code as string)}</span>
                        <div className="mt-10 flex flex-col">
                            <div className="">
                                <Field.FormControl className="">
                                    <Field.Input type="number" name="quantity" placeholder={"4"} onValueChange={setQuantity}
                                                 labelText="Quantità"
                                                 value={quantity}/>
                                </Field.FormControl>
                            </div>
                            <Button className="ml-auto mt-4" type="secondary" onClick={() => {
                                if (data?.productItem) {
                                    dispatch(addCartItem({
                                        itemId: data.productItem.id,
                                        quantity: quantity
                                    })).then(() => setSuccess('Elemento aggiunto a carrello'));
                                }
                            }}>
                                {t('product_details.order_button')}
                            </Button>
                        </div>
                    </div>
                </div>
            </Panel>
            {
                data?.productItem?.product?.eprelProductCode &&
                <Panel className="flex justify-center">
                    <Img
                        width={200}
                        className="shadow"
                        src={`https://eprel.ec.europa.eu/api/products/tyres/${data?.productItem?.product?.eprelProductCode}/labels?format=SVG`}
                        alt="eprel label"></Img>
                </Panel>
            }
        </div>
    </main>;
};

export default ProductDetailsPage;