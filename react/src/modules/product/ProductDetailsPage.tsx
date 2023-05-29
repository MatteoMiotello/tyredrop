import {useQuery} from "@apollo/client";
import React, {useEffect} from "react";
import {Img} from "react-image";
import {useParams} from "react-router-dom";
import tyrePlaceholder from "../../assets/placeholder-tyre.jpg";
import {PRODUCT_ITEM} from "../../common/backend/graph/query/products";
import Button from "../../common/components-library/Button";
import Field from "../../common/components-library/Input";
import Spinner from "../../common/components/Spinner";
import {ProductCategorySet} from "./enums/product-specifications-set";
import ProdapiService from "./services/prodapi/prodapi-service";

const loadingPlaceholder = <main>
    <Spinner/>
</main>;

const ProductDetailsPage: React.FC = () => {
    const {id} = useParams<{ id: string }>();

    if (!id) {
        return loadingPlaceholder;
    }

    const {data, loading, error, refetch} = useQuery(PRODUCT_ITEM, {
        variables: {
            id: id as string
        }
    });

    useEffect(() => {
        if (id) {
            refetch({
                id: id
            });
        }
    }, []);

    if (error) {
        return <div>
            {error.message}
        </div>;
    }

    if (loading && !data) {
        return loadingPlaceholder;
    }

    return <main className="lg:p-24 p-4">
        <div className="grid grid-cols-3 w-full">
            <div>
                <Img src={[
                    (new ProdapiService()).getProductImageUrl(data?.productItem?.product.code as string, ProductCategorySet.TYRE),
                    tyrePlaceholder,
                ]}
                     onErrorCapture={(e) => e.preventDefault()}
                     loading="lazy"
                     className="h-72"
                />
            </div>
            <div>
                ciao
            </div>
            <div className="bg-base-200 rounded-box p-10 pt-14">
                <span className="text-primary text-5xl font-semibold">
                    78,90€
                </span>
                <div className="flex mt-10 grid grid-cols-3 gap-4">
                    <div className="col-span-2">
                        <Field.Input type="number" name="quantity" placeholder={"4"} labelText="Quantita`"/>
                    </div>
                    <Button type="primary" className="mt-9">
                        Ordina
                    </Button>
                </div>
            </div>
        </div>
    </main>;
};

export default ProductDetailsPage;