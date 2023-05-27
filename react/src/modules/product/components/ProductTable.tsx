import {CellContext, ColumnDef} from "@tanstack/react-table";
import React, {useEffect, useState} from "react";
import {SearchQuery} from "../../../__generated__/graphql";
import Table from "../../../common/components-library/Table";
import ProductTitleCell from "./ProductTitleRow";
import {Img} from "react-image";
import tyrePlaceholder from "../../../assets/placeholder-tyre.jpg";
import Spinner from "../../../common/components/Spinner";
import {Currency} from "../../../common/utilities/currency";
import Button from "../../../common/components-library/Button";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faAngleRight, faShoppingCart} from "@fortawesome/free-solid-svg-icons";
import ProductSpecificationsGroup from "./ProductSpecificationsGroup";
import {
    ProductCategorySet,
    ProductSpecificationDefinition,
    ProductSpecifications
} from "../enums/product-specifications-set";


type ProductTableProps = {
    products: SearchQuery
    handlePaginationChange: (index: number, size: number) => void
    pageCount: number
}

export type ProductRowItemData = {
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
            cell: (props: CellContext<ProductRowItemData, any>) => {
                return <div className="h-full w-32">
                    <Img src={[
                        "http://localhost:8081/resources/tyres/" + props.row.original.code,
                        tyrePlaceholder,
                    ]} loading="lazy" className="h-full" alt={props.row.original.name}/>
                </div>;
            }
        },
        {
            accessorKey: "brand",
            cell: (props: CellContext<ProductRowItemData, any>) => {
                return <span className="my-auto mx-10">
                <Img src={"http://localhost:8081/resources/brands/" + props.row.original.brand.code}
                     loading="lazy"
                     loader={<Spinner/>}
                     unloader={
                         <span
                             className="uppercase text-2xl font-bold text-center w-full">{props.row.original.brand.name}
                    </span>
                     }
                />
                </span>;
            }
        },
        {
            accessorKey: "content",
            cell: (props: CellContext<ProductRowItemData, any>) => <ProductTitleCell data={props.row.original}/>
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
                <Button
                    className="mx-2 aspect-square"
                    type="ghost"
                    outline={true}
                >
                    <FontAwesomeIcon
                        icon={faAngleRight}
                    />
                </Button>
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