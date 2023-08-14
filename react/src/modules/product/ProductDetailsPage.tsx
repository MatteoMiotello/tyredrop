import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {Img} from "react-image";
import {useLoaderData} from "react-router-dom";
import {Product, ProductItemQuery, ProductSpecificationValue} from "../../__generated__/graphql";
import Button from "../../common/components-library/Button";
import Panel from "../../common/components-library/Panel";
import {Input, Join} from "../../common/components/shelly-ui";
import Spinner from "../../common/components/Spinner";
import {Currency} from "../../common/utilities/currency";
import CompleteProductSpecificationsGroup from "./components/CompleteProductSpecificationsGroup";
import ProductTitle from "./components/ProductTitle";
import ProdapiService from "./services/prodapi/prodapi-service";
import {useDispatch} from "react-redux";
import {addCartItem} from "../cart/store/cart-slice";
import {ThunkDispatch} from "redux-thunk";
import AvailabilityBadge from "./components/AvailabilityBadge";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faMinus, faMoneyBill, faPlus, faTruckFast} from "@fortawesome/free-solid-svg-icons";
import {useToast} from "../../store/toast";
import {Simulate} from "react-dom/test-utils";
import error = Simulate.error;
import ProductImage from "./components/ProductImage";

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
    const {success, error} = useToast();

    useEffect(() => {
        if (res.data) {
            setData(res.data);
        }

        setLoading(res.loading);
    }, [res]);

    if (loading && !data) {
        return loadingPlaceholder;
    }

    return <main className="lg:px-24 p-4">
        <div className="grid md:grid-cols-4 gap-4 w-full">
            <Panel className="flex flex-col justify-center items-center">
                <ProductImage product={data?.productItem?.product as Product}/>
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
                    <span>{data?.productItem?.product.vehicleType.name}</span>
                    <CompleteProductSpecificationsGroup
                        specifications={data?.productItem?.product.productSpecificationValues as ProductSpecificationValue[]}
                    />
                </div>
            </Panel>
            <div>
                <Panel>
                    <div className="rounded-box md:mt-0 mt-4">
                        <div className="flex flex-col gap-4">
                            <div>
                                Totale
                            </div>
                            <span className="text-primary text-5xl font-semibold">
                                {Currency.defaultFormat(data?.productItem?.price[0]?.value as number, data?.productItem?.price[0]?.currency.iso_code as string)}</span>
                            {data?.productItem && <AvailabilityBadge quantity={data?.productItem?.supplierQuantity}/>}
                            <div className="mt-10 flex items-center gap-2">
                                <div className="ml-auto">
                                    <Join>
                                        <Button className="join-item" disabled={quantity == 0} onClick={() => {
                                            if (quantity <= 0) {
                                                return;
                                            }
                                            setQuantity(quantity - 1);
                                        }}><FontAwesomeIcon icon={faMinus}/></Button>
                                        <Input.FormControl className="join-item">
                                            <Input
                                                min={0}
                                                bordered={false}
                                                className="!w-14 !outline-none input-ghost text-center"
                                                value={String( quantity )}
                                                type="number"
                                                name="quantity"
                                                placeholder="1"
                                                onValueChange={ (val) => {
                                                    setQuantity(Number(val));
                                                }}
                                                validators={[(value) => {
                                                    if (Number(value) > (data?.productItem?.supplierQuantity ?? 0)) {
                                                        return 'La quantità selezionata non è disponibile';
                                                    }

                                                    return null;
                                                }]}
                                            />
                                        </Input.FormControl>
                                        <Button className="join-item" onClick={() => {
                                            setQuantity(quantity + 1);
                                        }}><FontAwesomeIcon icon={faPlus}/></Button>
                                    </Join>
                                </div>
                                <Button className="" type="primary" onClick={() => {
                                    if (data?.productItem) {
                                        dispatch(addCartItem({
                                            itemId: data.productItem.id,
                                            quantity: quantity
                                        })).unwrap()
                                            .then(() => success('Elemento aggiunto a carrello'))
                                            .catch( () => error( 'Quantità non disponibile' ) );
                                    }
                                }}>
                                    {t('product_details.order_button')}
                                </Button>
                            </div>
                        </div>
                    </div>
                </Panel>
                <div className="bg-secondary text-base-100 mt-4 rounded-box p-4">
                    <p className="my-2"><FontAwesomeIcon icon={faTruckFast}/> Consenga in 4-5 giorni lavorativi</p>
                    <p className="my-2"><FontAwesomeIcon icon={faMoneyBill}/> Pagamento SEPA o Bonifico Bancario</p>
                </div>
            </div>
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