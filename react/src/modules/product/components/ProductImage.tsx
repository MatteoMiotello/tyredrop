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
const ProductImage: React.FC<ProductImageProps> = ( {product} ) =>  {
    const [referenceElement, setReferenceElement] = useState();
    const [popperElement, setPopperElement] = useState();
    const { styles, attributes } = usePopper(referenceElement, popperElement, {
        placement: 'right',
        modifiers: [
            {
                name: 'shift',
                enabled: true
            }
        ]
    });

    return <Popover>
        <Popover.Button as="button"
                        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                        // @ts-ignore
                        ref={setReferenceElement}
                        className="!outline-none">
            <Img src={[
                (new ProdapiService()).getProductImageUrl(product.code, ProductCategorySet.TYRE),
                tyrePlaceholder,
            ]}
                 loading="lazy"
                 className="mx-auto"
                 alt={product.name as string}/>
        </Popover.Button>
        <Popover.Panel
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
        </Popover.Panel>
    </Popover>;
};

export default ProductImage;