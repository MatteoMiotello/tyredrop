import {faAngleRight} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import {Img} from "react-image";
import {Link} from "react-router-dom";
import {Product, SearchQuery} from "../../../__generated__/graphql";
import Table from "../../../common/components-library/Table";
import {Currency} from "../../../common/utilities/currency";
import {
    ProductCategorySet,
    ProductSpecificationDefinition,
    ProductSpecifications
} from "../enums/product-specifications-set";
import ProdapiService from "../services/prodapi/prodapi-service";
import AddItemToCartButton from "./AddItemToCartButton";
import ProductSpecificationsGroup from "./ProductSpecificationsGroup";
import ProductTitle from "./ProductTitle";
import ProductQualityBadge from "./ProductQualityBadge";
import AvailabilityBadge from "./AvailabilityBadge";
import ProductImage from "./ProductImage";


type ProductTableProps = {
    products: SearchQuery
    handlePaginationChange: (index: number, size: number) => void
    pageCount: number
}

export type ProductRowItemData = {
    id: string
    supplier_quantity: number
    brand: {
        name: string,
        code: string,
        quality: number
    },
    name: string,
    code: string,
    price: string,
    specifications: (ProductSpecificationDefinition | null)[]
    product: Product
}

const ProductTable: React.FC<ProductTableProps> = (props) => {
    const [data, setData] = useState<(ProductRowItemData | null)[]>([]);

    const colums: ColumnDef<ProductRowItemData>[] = [
        {
            accessorKey: "image",
            enableResizing: true,
            size: 15,
            cell: (props: CellContext<ProductRowItemData, any>) => {
                return <div className="w-24">
                    <ProductImage product={props.row.original.product}/>
                    <Img src={(new ProdapiService()).getBrandImageUrl(props.row.original.brand.code)}
                         loading="lazy"
                         className="my-auto"
                         unloader={<span
                             className="text-xl uppercase text-neutral font-semibold">{props.row.original.brand.name}</span>}
                         onErrorCapture={(e) => e.preventDefault()}
                    />
                </div>;
            }
        },
        {
            accessorKey: "content",
            size: 250,
            cell: (props: CellContext<ProductRowItemData, any>) => <ProductTitle showBrand={true}
                                                                                 data={props.row.original}/>
        },
        {
            accessorKey: "specifications",
            cell: (props: CellContext<ProductRowItemData, any>) => <div className="flex justify-center items-center">
                <ProductQualityBadge quality={props.row.original.brand.quality}/>
                <ProductSpecificationsGroup
                    specifications={props.row.original.specifications}/>
            </div>
        },
        {
            accessorKey: "supplier_quantity",
            size: 200,
            cell: (props: CellContext<ProductRowItemData, any>) => <AvailabilityBadge quantity={props.getValue()}/>
        },
        {
            accessorKey: "price",
            size: 10,
            cell: (props) => {
                return <div className="w-full flex flex-col justify-center items-center">

                <span className="font-semibold text-xl">
                    {props.row.original.price}
                </span>
                    <span className="text-secondary text-sm font-semibold uppercase">
                        Esentasse
                    </span>
                </div>;
            }
        },
        {
            accessorKey: "button",
            size: 10,
            cell: (props: CellContext<ProductRowItemData, any>) => <div className="flex justify-center items-center">
                <AddItemToCartButton itemId={props.row.original.id} quantity={props.row.original.supplier_quantity}/>
                <Link
                    className="mx-2 aspect-square"
                    type="ghost"
                    to={"/products/details/" + props.row.original.id}
                >
                    <FontAwesomeIcon
                        icon={faAngleRight}
                    />
                </Link>
            </div>
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
                    code: product.product.brand.code,
                    quality: product.product.brand.quality as number
                },
                supplier_quantity: product.supplierQuantity,
                name: product.product.name as string,
                code: product.product.code,
                product: product.product as Product,
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