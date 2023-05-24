
export type ProductRowItemData = {
    brand: {
        name: string,
        code: string
    },
    name: string,
    code: string,
    price: string
}

type TyreItemRowProps = {
    data: ProductRowItemData
}

const ProductItemTitleRow: React.FC<TyreItemRowProps | null> = (props: TyreItemRowProps | null) => {
    if (!props?.data) {
        return <></>;
    }

    const data = props.data;

    return <div className="h-36 flex w-full justify-end items-center">
        <div>
            <h2 className="font-semibold"> {data.brand.name } </h2>
            <h2> {data.name} </h2>
        </div>
        <span
            className="ml-auto font-semibold text-xl"
        >
            {data.price}
        </span>
    </div>;
};

export default ProductItemTitleRow;