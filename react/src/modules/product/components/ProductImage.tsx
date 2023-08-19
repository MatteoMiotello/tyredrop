import {faMagnifyingGlass} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useState} from "react";
import {Popover} from "@headlessui/react";
import {Img} from "react-image";
import ProdapiService from "../services/prodapi/prodapi-service";
import {ProductCategorySet} from "../enums/product-specifications-set";
import tyrePlaceholder from "../../../assets/placeholder-tyre.jpg";
import {Product} from "../../../__generated__/graphql";
import {usePopper} from "react-popper";


type ProductImageProps = {
    product: Product
}
const ProductImage: React.FC<ProductImageProps> = ({product}) => {
    const [referenceElement, setReferenceElement] = useState();
    const [popperElement, setPopperElement] = useState();
    const [display, setDisplay] = useState<boolean>(false);
    const [arrowElement, setArrowElement] = useState(null);
    const {styles, attributes} = usePopper(referenceElement, popperElement, {
        placement: 'right',
        modifiers: [
            {
                name: 'shift',
                enabled: true
            },
            {
                name: 'arrow',
                options: {
                    element: arrowElement
                },
                enabled: true
            },
            {
                name: 'offset',
                options: {
                    offset: [0, 8],
                },
            },
        ]
    });

    // @ts-ignore
    return <Popover>
        <div className="relative" onMouseEnter={() => setDisplay(true)} onMouseLeave={() => setDisplay(false)}>
            <Popover.Button as="button"
                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                // @ts-ignore
                            ref={setReferenceElement}
                            className="!outline-none">
                {
                    display &&
                    <div className="absolute z-30 w-full h-full bg-transparent/20 font-bold text-base-100 text-2xl">
                        <div className="flex justify-center items-center h-full">
                            <FontAwesomeIcon icon={faMagnifyingGlass}/>
                        </div>
                    </div>
                }
                <Img src={[
                    (new ProdapiService()).getProductImageUrl(product.code, ProductCategorySet.TYRE),
                    tyrePlaceholder,
                ]}
                     loading="lazy"
                     className="mx-auto"
                     alt={product.name as string}/>
            </Popover.Button>
        </div>
        <Popover.Panel
            id="tooltip"
            className="shadow border rounded-box z-10"
            as="div"
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            ref={setPopperElement}
            style={styles.popper}
            {...attributes.popper}
        >
            <Img src={[
                (new ProdapiService()).getProductImageUrl(product.code, ProductCategorySet.TYRE),
                tyrePlaceholder,
            ]}
                 loading="lazy"
                 className="max-w-lg"
                 alt={product.name as string}/>
            <div id="arrow"
                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                // @ts-ignore
                 ref={setArrowElement}
                 style={styles.arrow}
                 className="!bg-base-100 before:border-l before:border-b"
            />
        </Popover.Panel>
    </Popover>;
};

export default ProductImage;