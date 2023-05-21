import {faAngleRight, faShoppingCart} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {Img} from "react-image";
import tyrePlaceholder from "../../../assets/placeholder-tyre.jpg";
import Button from "../../../common/components-library/Button";
import Spinner from "../../../common/components/Spinner";

export type ProductRowItemData = {
    brand: {
        name: string,
        code: string
    },
    name: string,
    code: string,
    price: {
        value: number,
        symbol: string
    }
}

type TyreItemRowProps = {
    data: ProductRowItemData
}

const TyreItemRow: React.FC<TyreItemRowProps | null> = (props: TyreItemRowProps | null) => {
    if (!props?.data) {
        return <></>;
    }

    const data = props.data;

    return <div className="h-36 flex w-full items-center">
        <Img src={[
            "http://localhost:8081/resources/tyres/" + data.code,
            tyrePlaceholder,
        ]} loading="lazy" className="h-full" alt={data.name}/>
        <span className="w-32 my-auto mx-10">
                <Img src={"http://localhost:8081/resources/brands/" + props.data.brand.code}
                     loading="lazy"
                     loader={<Spinner/>}
                     unloader={<span className="uppercase text-2xl font-bold text-center w-full">{data.brand.name}
                </span>}/>
        </span>
        <div>
            <h2 className="font-semibold"> {data.brand.name } </h2>
            <h2> {data.name} </h2>
        </div>
        <span className="ml-auto"> {data.price.value + " " + data.price.symbol} </span>
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
    </div>;
};

export default TyreItemRow;