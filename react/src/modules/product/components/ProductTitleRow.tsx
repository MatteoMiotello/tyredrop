import {ProductRowItemData} from "./ProductTable";


type TyreItemRowProps = {
    data: ProductRowItemData
}

const ProductTitleRow: React.FC<TyreItemRowProps | null> = (props: TyreItemRowProps | null) => {
    if (!props?.data) {
        return <></>;
    }

    const data = props.data;

    return <div className="h-36 flex w-full items-center">
        <div>
            <h2 className="font-semibold "> {data.brand.name } </h2>
            <h2 className="break-words"> {data.name} </h2>
        </div>
    </div>;
};

export default ProductTitleRow;