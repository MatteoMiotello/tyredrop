import {faAngleRight, faShoppingCart} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import {Img} from "react-image";
import {Link} from "react-router-dom";
import {SearchQuery} from "../../../__generated__/graphql";
import tyrePlaceholder from "../../../assets/placeholder-tyre.jpg";
import Button from "../../../common/components-library/Button";
import Table from "../../../common/components-library/Table";
import {Currency} from "../../../common/utilities/currency";
import {
    ProductCategorySet,
    ProductSpecificationDefinition,
    ProductSpecifications
} from "../enums/product-specifications-set";
import ProdapiService from "../services/prodapi/prodapi-service";
import ProductSpecificationsGroup from "./ProductSpecificationsGroup";
import ProductTitle from "./ProductTitle";


type ProductTableProps = {
    products: SearchQuery
    handlePaginationChange: (index: number, size: number) => void
    pageCount: number
}

export type ProductRowItemData = {
    id: string
    brand: {
        name: string,
        code: string
    },
    name: string,
    code: string,
    price: string,
    specifications: (ProductSpecificationDefinition | null)[]
}

const ProductTable: React.FC<ProductTableProps> = (props) => {
    const [data, setData] = useState<(ProductRowItemData | null)[]>([]);
    const colums: ColumnDef<ProductRowItemData>[] = [
        {
            accessorKey: "image",
            enableResizing: true,
            size: 20,
            cell: (props: CellContext<ProductRowItemData, any>) => {
                return <div className="w-24">
                    <Img src={[
                        (new ProdapiService()).getProductImageUrl(props.row.original.code, ProductCategorySet.TYRE),
                        tyrePlaceholder,
                    ]}
                         onErrorCapture={(e) => e.preventDefault()}
                         loading="lazy"
                         className="h-24"
                         alt={props.row.original.name}/>
                </div>;
            }
        },
        {
            accessorKey: "content",
            cell: (props: CellContext<ProductRowItemData, any>) => <ProductTitle showBrand={true} data={props.row.original}/>
        },
        {
            accessorKey: "specifications",
            cell: (props: CellContext<ProductRowItemData, any>) => <ProductSpecificationsGroup
                specifications={props.row.original.specifications}/>
        },
        {
            accessorKey: "price",
            cell: (props) => {
                return <span
                    className="ml-auto font-semibold text-xl"
                >
                    {props.row.original.price}
            </span>;
            }
        },
        {
            accessorKey: "button",
            cell: (props: CellContext<ProductRowItemData, any>) => <>
                <Button
                    className="mx-2 aspect-square"
                    type={"primary"}
                >
                    <FontAwesomeIcon icon={faShoppingCart}/>
                </Button>
                <Link
                    className="mx-2 aspect-square"
                    type="ghost"
                    to={"/products/details/" + props.row.original.id}
                >
                    <FontAwesomeIcon
                        icon={faAngleRight}
                    />
                </Link>
            </>
        }
    ];

    useEffect(() => {
        if (!props.products || !props.products.productItems || !props.products.productItems.productItems) {
            setData([]);
            return;
        }

        const data = props.products.productItems.productItems.map((product): ProductRowItemData | null => {
            if (!product || !product.price[0]) {
                return null;
            }

            return {
                id: product.id,
                brand: {
                    name: product.product.brand.name,
                    code: product.product.brand.code
                },
                name: product.product.name as string,
                code: product.product.code,
                price: Currency.defaultFormat(product.price[0]?.value, product.price[0]?.currency.iso_code),
                specifications: product.product.productSpecificationValues.map((value) => {
                    const icon = ProductSpecifications.getSpecificationIcon(ProductCategorySet.TYRE, value?.specification.code as string);

                    if (!icon) {
                        return null;
                    }

                    return {
                        code: value?.specification.code,
                        value: value?.value,
                        icon: icon
                    } as ProductSpecificationDefinition;
                })
            };
        });

        setData(data ?? []);
    }, [props.products]);

    return <Table
        hideHeader={true}
        data={data}
        columns={colums}
        pageCount={props.pageCount}
        updatePagination={props.handlePaginationChange}
    />;
};

export default ProductTable;